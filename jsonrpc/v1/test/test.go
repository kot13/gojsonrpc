package test

import (
	"encoding/json"
	"errors"
)

type TestParams struct {
	Param    string `json:"param" valid:"string,required"`
}

func Test(p json.RawMessage) (result interface{}, err error) {
	var params TestParams

	err = json.Unmarshal(p, &params)
	if err != nil {
		return
	}

	if params.Param == "" {
		err = errors.New("Not valid param")
		return
	}

	result = map[string]string{
		"param": params.Param,
	}

	return
}
