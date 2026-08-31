-- migrations/0000_init.up.sql
CREATE TABLE projects (
                          id          UUID PRIMARY KEY,
                          name        TEXT NOT NULL CHECK (length(name) BETWEEN 1 AND 200),
                          status      TEXT NOT NULL CHECK (status IN ('active','archived')),
                          created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
                          version     INT  NOT NULL DEFAULT 1
);

CREATE TABLE tasks (
                       id          UUID PRIMARY KEY,
                       project_id  UUID NOT NULL REFERENCES projects(id) ON DELETE RESTRICT,
                       title       TEXT NOT NULL CHECK (length(title) BETWEEN 1 AND 300),
                       description TEXT NOT NULL DEFAULT '',
                       status      TEXT NOT NULL CHECK (status IN ('todo','in_progress','done','cancelled')),
                       priority    SMALLINT NOT NULL CHECK (priority BETWEEN 1 AND 4),
                       assignee_id UUID NULL,
                       due_at      TIMESTAMPTZ NULL,
                       created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
                       version     INT NOT NULL DEFAULT 1
);

-- the index that actually matters: the list query's WHERE + ORDER BY
CREATE INDEX idx_tasks_project_status ON tasks (project_id, status, created_at DESC);
CREATE INDEX idx_tasks_assignee ON tasks (assignee_id) WHERE assignee_id IS NOT NULL;

-- idempotency for task creation
CREATE TABLE task_idempotency (
                                  key        TEXT PRIMARY KEY,
                                  task_id    UUID NOT NULL REFERENCES tasks(id),
                                  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- transactional outbox
CREATE TABLE outbox (
                        id           BIGSERIAL PRIMARY KEY,
                        aggregate_id UUID NOT NULL,
                        event_type   TEXT NOT NULL,
                        payload      JSONB NOT NULL,
                        created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
                        published_at TIMESTAMPTZ NULL
);
CREATE INDEX idx_outbox_unpublished ON outbox (id) WHERE published_at IS NULL;
