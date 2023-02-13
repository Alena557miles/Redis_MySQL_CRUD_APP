package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseCreate(obj string, str string, rw http.ResponseWriter) {
	resp := make(map[string]string)
	resp["message"] = obj + ` with name: ` + str + ` was created successfully`
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func ResponseAction(obj1 string, obj1N string, obj2 string, obj2N string, act string, rw http.ResponseWriter) {
	resp := make(map[string]string)
	switch {
	case act == "update":
		resp["message"] = obj1 + ` : ` + obj1N + ` was renamed to ` + obj2 + ` : ` + obj2N
	case act == "deleted":
		resp["message"] = obj1 + ` : ` + obj1N + ` was ` + act + ` successfully ` + obj2 + obj2N
	default:
		resp["message"] = obj1 + ` : ` + obj1N + `is` + act + ` on ` + obj2 + ` : ` + obj2N
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)

}
