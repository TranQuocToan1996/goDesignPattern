package structuralPatterns

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator it must return the text %sthe expected text, not '%s'", pizzaResult,
			expectedText)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the onion decorator without"+
			"an IngredientAdd on its Ingredient field must return an error, not a string with '%s'", onionResult)
	}
	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it"+
			"must return a text with the word 'onion', not '%s'", onionResult)
	}
}
func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the meat decorator without"+
			"an IngredientAdd in its Ingredient field must return an error,"+"not a string with '%s'", meatResult)
	}
	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the add ingredient of the meat decorator it"+
			"must return a text with the word 'meat', not '%s'", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned "+
			"string must contain the text '%s' but '%s' didn't have it",
			expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n",
		r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n",
		r.Method)
	fmt.Fprintf(s.LogWriter, "--------------------------------\n")
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

func TestA(t *testing.T) {
	http.Handle("/", &LoggerServer{
		LogWriter: os.Stdout,
		Handler:   &MyServer{},
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
