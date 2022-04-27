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
		msg = dbaccess.CreateAuthor(req.Name, req.DateOfBirth)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}

}

func CreateBook(req *structs.CreateBookReq, resp *structs.CreateBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.CreateBook(req.AuthorName, req.BookName, req.Topic, req.Date, req.Price, req.Quantity)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}

}

func CreateReader(req *structs.CreateReaderReq, resp *structs.CreateReaderResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.CreateReader(req.ReaderName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func OrderBook(req *structs.OrderBookReq, resp *structs.OrderBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.OrderBook(req.ReaderName, req.BookName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func BuyBook(req *structs.BuyBookReq, resp *structs.BuyBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.BuyBook(req.ReaderName, req.BookName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func WatchAuthor(req *structs.WatchAuthorReq, resp *structs.WatchAuthorResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.WatchAuthor(req.ReaderName, req.AuthorName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func WatchBook(req *structs.WatchBookReq, resp *structs.WatchBookResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.WatchBook(req.ReaderName, req.BookName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func CleanDB(req *structs.CleanDBReq, resp *structs.CleanDBResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		msg = dbaccess.CleanDB()
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
		} else {
			resp.Success = false
			resp.Msg = msg
		}
		resp.Success = true
		resp.Msg = "all good"
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func WatchedAuthorByReader(req *structs.WatchedAuthorByReaderReq, resp *structs.WatchedAuthorByReaderResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		list, msg := dbaccess.WatchedAuthorByReader(req.ReaderName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
			resp.AuthorNames = list
		} else {
			resp.Success = false
			resp.Msg = msg
			resp.AuthorNames = nil
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func AuthorByTopic(req *structs.AuthorByTopicReq, resp *structs.AuthorByTopicResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		list, msg := dbaccess.AuthorByTopic(req.Topic)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
			resp.AuthorNames = list
		} else {
			resp.Success = false
			resp.Msg = msg
			resp.AuthorNames = nil
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func ReadersByAuthor(req *structs.ReadersByAuthorReq, resp *structs.ReadersByAuthorResp) {
	msg, isValid := ValidateRequest(*req)
	if isValid {
		list, msg := dbaccess.ReadersByAuthor(req.AuthorName)
		if len(strings.TrimSpace(msg)) == 0 {
			resp.Success = true
			resp.Msg = msg
			resp.AuthorNames = list
		} else {
			resp.Success = false
			resp.Msg = msg
			resp.AuthorNames = nil
		}
	} else {
		resp.Success = isValid
		resp.Msg = msg
	}
}

func ValidateRequest(req interface{}) (string, bool) {
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
