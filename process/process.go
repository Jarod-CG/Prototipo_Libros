package process

import (
	"github.com/Jarod-CG/Prototipo_Libros/dbaccess"
	"github.com/Jarod-CG/Prototipo_Libros/structs"
)

func CreateAuthor(req *structs.CreateAuthorReq, resp *structs.CreateAuthorResp) {

	dbaccess.CreateAuthor(req.Name, req.DateOfBirth)
	resp.Success = true
	resp.Msg = "all good"
}
