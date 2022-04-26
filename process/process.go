package process

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Jarod-CG/Prototipo_Libros/dbaccess"
	"github.com/Jarod-CG/Prototipo_Libros/structs"
)

func CreateAuthor(req *structs.CreateAuthorReq, resp *structs.CreateAuthorResp) {

	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.CreateAuthor(req.Name, req.DateOfBirth)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}

}

func CreateBook(req *structs.CreateBookReq, resp *structs.CreateBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.CreateBook(req.AuthorName, req.BookName, req.Topic, req.Date, req.Price, req.Quantity)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}

}

func CreateReader(req *structs.CreateReaderReq, resp *structs.CreateReaderResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.CreateReader(req.ReaderName)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func OrderBook(req *structs.OrderBookReq, resp *structs.OrderBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.OrderBook(req.ReaderName, req.BookName)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func BuyBook(req *structs.BuyBookReq, resp *structs.BuyBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.BuyBook(req.ReaderName, req.BookName)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func WatchAuthor(req *structs.WatchAuthorReq, resp *structs.WatchAuthorResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.WatchAuthor(req.ReaderName, req.AuthorName)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func WatchBook(req *structs.WatchBookReq, resp *structs.WatchBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.WatchBook(req.ReaderName, req.BookName)
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func CleanDB(req *structs.CleanDBReq, resp *structs.CleanDBResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		dbaccess.CleanDB()
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func ValidateRequest(req interface{}) (string, bool) {
	t := reflect.TypeOf(req).Name()

	fmt.Println(t)
	v := reflect.ValueOf(req)

	num := v.NumField()
	for i := 0; i < num; i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int:
			value := v.Field(i).Int()
			if value <= 0 {
				typeField := reflect.TypeOf(field).Name()
				msg := fmt.Sprintf("Error : field type '%s', '%s' value is '%d'", field.Kind(), typeField, value)
				return msg, false
			}
		case reflect.String:
			value := v.Field(i).String()
			if len(strings.TrimSpace(value)) == 0 {
				typeField := reflect.TypeOf(field).Name()
				msg := fmt.Sprintf("Error : field type '%s', '%s' value is '%s'", field.Kind(), typeField, value)
				return msg, false
			}
		case reflect.Float64:
			value := v.Field(i).Float()
			if value <= 0 {
				typeField := reflect.TypeOf(field).Name()
				msg := fmt.Sprintf("Error : field type '%s', '%s' value is '%f'", field.Kind(), typeField, value)
				return msg, false
			}

		default:
			typeField := reflect.TypeOf(field).Name()
			msg := fmt.Sprintf("Unsuported type '%s', '%s' value is '%v'", field.Kind(), typeField, field)
			return msg, true
		}
	}
	return "", true
}
