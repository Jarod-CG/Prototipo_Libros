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

const (
	port = 8080
)

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

	fmt.Printf("server starsted at port '%d'\n", port)

	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), router.Handler))
}

func CreateAuthor(ctx *fasthttp.RequestCtx) {

	var (
		request  *structs.CreateAuthorReq
		response *structs.CreateAuthorResp
	)

	response = &structs.CreateAuthorResp{}

	err := json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		process.CreateAuthor(request, response)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ctx.Response.SetBody(bytes)
	}

}

func CreateBook(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func CreateReader(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func OrderBook(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func BuyBook(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func WatchAuthor(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func WatchBook(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}
