package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request, resp *Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	bts, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(fmt.Errorf("marshalling response to json failed, err: %v", err))
		return
	}
	_, err = w.Write(bts)
	if err != nil {
		fmt.Println(fmt.Errorf("writing response to response-writer failed, err: %v", err))
	}
}
