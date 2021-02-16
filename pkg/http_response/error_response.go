package http_response

import (
	"errors"
	"net/http"
)

func ErrorsReturnEntity(rw http.ResponseWriter, er error, exists bool, returnResp interface{}) { //Improve this
	switch {
	case er != nil:
		ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	case exists == false:
		ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
		return
	default:
		JsonResponse(rw, http.StatusOK, returnResp)
		return
	}
}
