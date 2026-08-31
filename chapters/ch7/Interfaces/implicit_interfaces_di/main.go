package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

//factory function to create an instance of a SimpleDataStore

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{ //initializing the map
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// interfaces to avoid concrete classes to depend on, but instead interfaces to make it based on abstraction rather than concrete classes
type DataStore interface {
	UserNameForId(userId string) (string, bool)
}
type Logger interface {
	Log(message string)
}

// to make LogOutput function meet or fit on the interface of Logger, you define a function type with a method on it:
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// let's define the business logic
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in say hello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("in say goodbye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

//whe you want a SimpleLogic instance, you call a factory function, passing interfaces as parameters and returning concrete classes  or (structs)

func NewSimpleLogic(l Logger, ds SimpleDataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// controller interface
type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Say Hello Controller")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return
	}
}

// factory function for the controller
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// wire up all  your components in your main function
func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
