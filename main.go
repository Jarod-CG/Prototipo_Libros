package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Jarod-CG/Prototipo_Libros/process"
	"github.com/Jarod-CG/Prototipo_Libros/structs"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func init() {
	//load json
	if structs.Config.LoadConfig() {
		fmt.Println("Json config loaded succesfully!!")
	}
}

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.POST("/CreateAuthor", CreateAuthor)
	router.POST("/CreateBook", CreateBook)
	router.POST("/CreateReader", CreateReader)
	router.POST("/OrderBook", OrderBook)
	router.POST("/BuyBook", BuyBook)
	router.POST("/WatchAuthor", WatchAuthor)
	router.POST("/WatchBook", WatchBook)
	router.POST("/CleanDB", CleanDB)

	fmt.Printf("server starsted at port '%d'\n", structs.Config.Port)

	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(structs.Config.Port)), router.Handler))
}

func CreateAuthor(ctx *fasthttp.RequestCtx) {

	var (
		request  structs.CreateAuthorReq
		response structs.CreateAuthorResp
	)

	response = structs.CreateAuthorResp{}

	err := json.Unmarshal(ctx.Request.Body(), &request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateAuthor(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}

}

func CreateBook(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.CreateBookReq
		response structs.CreateBookResp
	)

	response = structs.CreateBookResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateBook(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func CreateReader(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.CreateReaderReq
		response structs.CreateReaderResp
	)

	response = structs.CreateReaderResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateReader(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func OrderBook(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.OrderBookReq
		response structs.OrderBookResp
	)

	response = structs.OrderBookResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.OrderBook(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func BuyBook(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.CreateAuthorReq
		response structs.CreateAuthorResp
	)

	response = structs.CreateAuthorResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateAuthor(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func WatchAuthor(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.WatchAuthorReq
		response structs.WatchAuthorResp
	)

	response = structs.WatchAuthorResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.WatchAuthor(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func WatchBook(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.CreateAuthorReq
		response structs.CreateAuthorResp
	)

	response = structs.CreateAuthorResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateAuthor(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}

func CleanDB(ctx *fasthttp.RequestCtx) {
	var (
		request  structs.CleanDBReq
		response structs.CleanDBResp
	)

	response = structs.CleanDBResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CleanDB(&request, &response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}
}
