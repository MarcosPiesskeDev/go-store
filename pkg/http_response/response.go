package http_response

import (
	"encoding/json"
	"net/http"
)

func ErrResponse(wr http.ResponseWriter, statusCode int, message string) {
	JsonResponse(wr, statusCode, map[string]string{"json_error": message})
}

func JsonResponse(wr http.ResponseWriter, statusCode int, pl interface{}) {
	resp, _ := json.Marshal(pl)
	wr.Header().Set("Content-Type", "application-json")
	wr.WriteHeader(statusCode)
	wr.Write(resp)
}

//'pl' means Payload, that means the part of transmitted data that is the actual intended message.
//Headers and metadata are sent only to enable payload delivery.
