package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type DefaultResponse struct {
	Status bool        `json:"status"`
	Error  string      `json:"error"'`
	Body   interface{} `json:"body"`
}

func ReturnResponseWriter(err error, w http.ResponseWriter, i interface{}, logMsg string) {
	res := DefaultResponse{}
	if err != nil {
		res.Status = false
		res.Body = i
		res.Error = err.Error()
		log.Println(logMsg + " " + err.Error())
	} else {
		res.Status = true
		res.Body = i
		log.Println(logMsg)
	}
	json.NewEncoder(w).Encode(res)
}
