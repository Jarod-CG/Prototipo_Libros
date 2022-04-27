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
		{Name: "Virginia Woolf", DateOfBirth: "1882-01-25"},
		{Name: "Yolanda Oreamuno", DateOfBirth: "1916-04-08"},
		{Name: "John Tolkien", DateOfBirth: "1892-01-03"},
		{Name: "Jane Austen", DateOfBirth: "1775-12-16"},
		{Name: "Isabel Allende", DateOfBirth: "1942-08-02"},
		{Name: "Albert Camus", DateOfBirth: "1913-11-07"},
		{Name: "Isaac Asimov", DateOfBirth: "1920-01-02"},
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

	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.CreateBookReq{
		{AuthorName: "Virginia Woolf", BookName: "La Señora Dalloway", Topic: "Novela", Date: "1920-01-02", Price: 99.99, Quantity: 100},
		{AuthorName: "Virginia Woolf", BookName: "Una Habitación Propia", Topic: "Ensayo", Date: "1920-01-02", Price: 199.99, Quantity: 100},
		{AuthorName: "Yolanda Oreamuno", BookName: "La Ruta de su Evasión", Topic: "Latinoamericano", Date: "1920-01-02", Price: 10.99, Quantity: 100},
		{AuthorName: "Yolanda Oreamuno", BookName: "A lo Largo del Corto Camino", Topic: "Latinoamericano", Date: "1920-01-02", Price: 99.99, Quantity: 100},
		{AuthorName: "John Tolkien", BookName: "El Señor de los Anillos", Topic: "Fantasía", Date: "1920-01-02", Price: 199.99, Quantity: 100},
		{AuthorName: "John Tolkien", BookName: "El Silmarillion", Topic: "Fantasía", Date: "1920-01-02", Price: 10, Quantity: 100},
		{AuthorName: "Jane Austen", BookName: "Orgullo y Prejuicio", Topic: "Novela", Date: "1920-01-02", Price: 199.99, Quantity: 100},
		{AuthorName: "Jane Austen", BookName: "Emma", Topic: "Novela", Date: "1920-01-02", Price: 99.99, Quantity: 100},
		{AuthorName: "Isabel Allende", BookName: "La Casa de los Espíritus", Topic: "Latinoamericano", Date: "1920-01-02", Price: 1, Quantity: 100},
		{AuthorName: "Isabel Allende", BookName: "Violeta", Topic: "Latinoamericano", Date: "1920-01-02", Price: 199.99, Quantity: 100},
		{AuthorName: "Albert Camus", BookName: "El Extranjero", Topic: "Novela filósofica", Date: "1942-01-02", Price: 12.99, Quantity: 100},
		{AuthorName: "Albert Camus", BookName: "La Peste", Topic: "Novela filósofica", Date: "1947-01-02", Price: 99.99, Quantity: 100},
		{AuthorName: "Isaac Asimov", BookName: "Fundacion", Topic: "Scifi", Date: "1942-01-02", Price: 100, Quantity: 100},
		{AuthorName: "Isaac Asimov", BookName: "Yo, Robot", Topic: "Scifi", Date: "1920-01-02", Price: 99.99, Quantity: 100},
	}

	for _, req := range requests {

		r, err := doPost("/CreateBook", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.CreateBookResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestCreateReader(t *testing.T) {

	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.CreateReaderReq{
		{ReaderName: "jarod"},
		{ReaderName: "laura"},
		{ReaderName: "mohammed"},
		{ReaderName: "li"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza"},
	}

	for _, req := range requests {

		r, err := doPost("/CreateReader", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.CreateReaderResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestOrderBook(t *testing.T) {

	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.OrderBookReq{
		{ReaderName: "jarod", BookName: "La Señora Dalloway"},
		{ReaderName: "laura", BookName: "Una Habitación Propia"},
		{ReaderName: "mohammed", BookName: "La Ruta de su Evasión"},
		{ReaderName: "li", BookName: "A lo Largo del Corto Camino"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "El Señor de los Anillos"},
		{ReaderName: "jarod", BookName: "El Silmarillion"},
		{ReaderName: "laura", BookName: "Orgullo y Prejuicio"},
		{ReaderName: "mohammed", BookName: "Emma"},
		{ReaderName: "li", BookName: "La Casa de los Espíritus"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "Violeta"},
		{ReaderName: "jarod", BookName: "El Extranjero"},
		{ReaderName: "laura", BookName: "La Peste"},
		{ReaderName: "mohammed", BookName: "Fundacion"},
		{ReaderName: "li", BookName: "Yo, Robot"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "La Señora Dalloway"},
	}

	for _, req := range requests {

		r, err := doPost("/OrderBook", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.OrderBookResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestBuyBook(t *testing.T) {
	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.BuyBookReq{
		{ReaderName: "jarod", BookName: "Una Habitación Propia"},
		{ReaderName: "laura", BookName: "La Ruta de su Evasión"},
		{ReaderName: "mohammed", BookName: "A lo Largo del Corto Camino"},
		{ReaderName: "li", BookName: "El Señor de los Anillos"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "El Silmarillion"},
		{ReaderName: "jarod", BookName: "Orgullo y Prejuicio"},
		{ReaderName: "laura", BookName: "Emma"},
		{ReaderName: "mohammed", BookName: "La Casa de los Espíritus"},
		{ReaderName: "li", BookName: "Violeta"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "El Extranjero"},
		{ReaderName: "jarod", BookName: "La Peste"},
		{ReaderName: "laura", BookName: "Fundacion"},
		{ReaderName: "mohammed", BookName: "Yo, Robot"},
		{ReaderName: "li", BookName: "La Señora Dalloway"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "La Ruta de su Evasión"},
		{ReaderName: "jarod", BookName: "A lo Largo del Corto Camino"},
		{ReaderName: "laura", BookName: "El Señor de los Anillos"},
		{ReaderName: "mohammed", BookName: "El Silmarillion"},
		{ReaderName: "li", BookName: "Orgullo y Prejuicio"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "Emma"},
	}

	for _, req := range requests {

		r, err := doPost("/BuyBook", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.BuyBookResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestWatchAuthor(t *testing.T) {

	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.WatchAuthorReq{
		{ReaderName: "jarod", AuthorName: "Jane Austen"},
		{ReaderName: "laura", AuthorName: "Virginia Woolf"},
		{ReaderName: "mohammed", AuthorName: "Isabel Allende"},
		{ReaderName: "li", AuthorName: "Albert Camus"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", AuthorName: "John Tolkien"},
		{ReaderName: "jarod", AuthorName: "Yolanda Oreamuno"},
		{ReaderName: "laura", AuthorName: "Isaac Asimov"},
		{ReaderName: "mohammed", AuthorName: "Jane Austen"},
		{ReaderName: "li", AuthorName: "Virginia Woolf"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", AuthorName: "Isabel Allende"},
		{ReaderName: "jarod", AuthorName: "Albert Camus"},
		{ReaderName: "laura", AuthorName: "John Tolkien"},
		{ReaderName: "mohammed", AuthorName: "Yolanda Oreamuno"},
		{ReaderName: "li", AuthorName: "Isaac Asimov"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", AuthorName: "Jane Austen"},
	}

	for _, req := range requests {

		r, err := doPost("/WatchAuthor", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.WatchAuthorResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestWatchBook(t *testing.T) {
	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.WatchBookReq{
		{ReaderName: "jarod", BookName: "La Casa de los Espíritus"},
		{ReaderName: "laura", BookName: "Violeta"},
		{ReaderName: "mohammed", BookName: "El Extranjero"},
		{ReaderName: "li", BookName: "La Peste"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "Fundacion"},
		{ReaderName: "jarod", BookName: "Yo, Robot"},
		{ReaderName: "laura", BookName: "La Señora Dalloway"},
		{ReaderName: "mohammed", BookName: "La Ruta de su Evasión"},
		{ReaderName: "li", BookName: "A lo Largo del Corto Camino"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza", BookName: "El Señor de los Anillos"},
	}

	for _, req := range requests {

		r, err := doPost("/WatchBook", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.WatchBookResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestWatchedAuthorByReader(t *testing.T) {
	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.WatchedAuthorByReaderReq{
		{ReaderName: "jarod"},
		{ReaderName: "laura"},
		{ReaderName: "mohammed"},
		{ReaderName: "li"},
		{ReaderName: "kkwazzawazzakkwaquikkwalaquaza"},
	}

	for _, req := range requests {

		r, err := doPost("/WatchedAuthorByReader", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.WatchedAuthorByReaderResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestAuthorByTopic(t *testing.T) {
	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.AuthorByTopicReq{
		{Topic: "Novela"},
		{Topic: "Ensayo"},
		{Topic: "Latinoamericano"},
		{Topic: "Fantasía"},
		{Topic: "Novela filósofica"},
		{Topic: "Scifi"},
	}

	for _, req := range requests {

		r, err := doPost("/AuthorByTopic", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.AuthorByTopicResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
}

func TestReadersByAuthor(t *testing.T) {
	go main()
	time.Sleep(10 * time.Second)

	requests := []structs.ReadersByAuthorReq{
		{AuthorName: "Virginia Woolf"},
		{AuthorName: "Yolanda Oreamuno"},
		{AuthorName: "John Tolkien"},
		{AuthorName: "Jane Austen"},
		{AuthorName: "Isabel Allende"},
		{AuthorName: "Albert Camus"},
		{AuthorName: "Isaac Asimov"},
	}

	for _, req := range requests {

		r, err := doPost("/ReadersByAuthor", req)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := &structs.ReadersByAuthorResp{}
		err = json.Unmarshal(r, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response '%s\n'", string(r))
	}
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
