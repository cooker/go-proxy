package testing

import (
	"cooker/go-proxy/utils"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRandomPort(t *testing.T) {

}

func TestResponseJson(t *testing.T) {
	response := http.Response{StatusCode: http.StatusOK}
	jsonStr, err := json.Marshal(response)
	if err != nil {

	}
	println(utils.ToString(&jsonStr))
}
