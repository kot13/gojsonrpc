package v1

import (
	"encoding/json"
	"gojsonrpc/jsonrpc"
	"gojsonrpc/jsonrpc/v1/test"
	"io/ioutil"
	"net/http"
	"errors"
)

var services = jsonrpc.Services{
	"test.test": test.Test,
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Разбираем запрос
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		jsonrpc.RenderResponse(jsonrpc.NewError(err, jsonrpc.InvalidRequestErrorCode, 1), w)
		return
	}

	var req jsonrpc.Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		jsonrpc.RenderResponse(jsonrpc.NewError(err, jsonrpc.ParseErrorCode, 1), w)
		return
	}

	// Валидируем запрос
	err = req.Validate()
	if err != nil {
		jsonrpc.RenderResponse(jsonrpc.NewError(err, jsonrpc.InvalidRequestErrorCode, req.Id), w)
		return
	}

	// Ищем процедуру
	procedure, ok := services[req.Method]
	if !ok {
		err = errors.New("Method not found")
		jsonrpc.RenderResponse(jsonrpc.NewError(err, jsonrpc.MethodNotFoundErrorCode, req.Id), w)
		return
	}

	// Запускаем процедуру с параметрами
	result, err := procedure(req.Params)
	if err != nil {
		jsonrpc.RenderResponse(jsonrpc.NewError(err, jsonrpc.InternalErrorCode, req.Id), w)
		return
	}

	jsonrpc.RenderResponse(jsonrpc.SuccessResponse{
		Version: jsonrpc.Version,
		Result:  result,
		Id:      req.Id,
	}, w)
}