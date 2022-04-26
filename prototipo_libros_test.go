package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/Jarod-CG/Prototipo_Libros/dbaccess"
	"github.com/Jarod-CG/Prototipo_Libros/process"
	"github.com/Jarod-CG/Prototipo_Libros/structs"
	"github.com/valyala/fasthttp"
)

const (
	server = "http://localhost:8080"
)

func TestNeo4j(t *testing.T) {
	dbaccess.TestNeo4j()
}

func TestCreateAuthor(t *testing.T) {

	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.CreateAuthorReq{
		{Name: "test", DateOfBirth: "1900-01-01"},
	}

	for _, req := range requests {

		r, err := doPost("/CreateAuthor", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.CreateAuthorResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestCreateBook(t *testing.T) {
	dbaccess.CreateBook("tolkien", "lord of the rings", "fantasy", "1900-01-01", 10.99, 20)
}

func TestCreateReader(t *testing.T) {
	dbaccess.CreateReader("jarod")
}

func TestOrderBook(t *testing.T) {
	dbaccess.OrderBook("jarod", "lord of the rings")
}

func TestBuyBook(t *testing.T) {
	dbaccess.BuyBook("jarod", "lord of the rings")
}

func TestWatchAuthor(t *testing.T) {
	dbaccess.WatchAuthor("jarod", "tolkien")
}

func TestWatchBook(t *testing.T) {
	dbaccess.WatchBook("jarod", "lord of the rings")
}

func doPost(uri string, body interface{}) ([]byte, error) {

	bytes, _ := json.Marshal(body)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(server + uri)
	req.Header.SetMethod("POST")
	req.SetBody(bytes)

	resp := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func TestValidateRequest(t *testing.T) {

	strct := []struct {
		Name  string
		Age   int
		Heigh float64
		Alive bool
	}{
		{Name: "jarod", Age: 22, Heigh: 179.00000, Alive: false},
		{Name: "jarod", Age: -22, Heigh: 179.00000, Alive: false},
		{Name: "jarod", Age: 22, Heigh: -179.00000, Alive: false},
		{Name: "", Age: 22, Heigh: 179.00000, Alive: false},
	}

	for _, st := range strct {

		bytes, _ := json.Marshal(st)

		msg, _ := process.ValidateRequest(st)

		fmt.Printf("%s response %s\n", bytes, msg)
	}

}
