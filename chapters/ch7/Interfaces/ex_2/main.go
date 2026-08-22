package main

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	//get data somewhere
	//c.L.Process(data)
}

func main() {
	// Go codes proves  and interface but only the caller (Client) knows about it
	//Nothing id declared on LogicProvider to indicate that it meets the interface
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
