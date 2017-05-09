package jsonrpc

import (
	"encoding/json"
	"net/http"
	"errors"
)

const (
	Version = "2.0"
)

type Request struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Id      int             `json:"id"`
}

type SuccessResponse struct {
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Id      int         `json:"id"`
}

type Procedure func(params json.RawMessage) (result interface{}, err error)

type Services map[string]Procedure

type Requests map[string]interface{}

func (r Request) Validate() error {
	if r.Version != Version {
		return errors.New("Not valid version")
	}

	if r.Method == "" {
		return errors.New("Method is required")
	}

	return nil
}

func RenderResponse(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}