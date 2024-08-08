package main

import (
	"errors"
	"fmt"
)

type bizError struct {
	code int32
	msg  string
}

func (e *bizError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.msg)
}

type baseResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func err2Resp(err error) *baseResp {
	if err == nil {
		return &baseResp{
			Code: 0,
			Msg:  "success",
		}
	}
	var be *bizError
	if errors.As(err, &be) {
		return &baseResp{
			Code: be.code,
			Msg:  be.msg,
		}
	}

	return &baseResp{
		Code: -1,
		Msg:  "server internal error",
	}

}
