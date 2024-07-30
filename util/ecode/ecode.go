package ecode

import (
	"fmt"
	"github.com/pkg/errors"
)

// New Error
func New(id, msg string) Ecode {
	return Ecode{
		ID:  id,
		Msg: msg,
	}
}

type Ecode struct {
	ID   string
	Msg  string
	Data interface{}
}

func (e Ecode) Error() string {
	return e.Msg
}

func (e *Ecode) AddData(data interface{}) {
	e.Data = data
}

func (e *Ecode) AddMsg(msg string) {
	e.Msg = fmt.Sprintf("%s - 【%s】", e.Msg, msg)
}

func GenInternalError(internalErr Ecode, err error) Ecode {
	internalErr.AddMsg(err.Error())
	return internalErr
}

// Code return code
func (e *Ecode) Code() string {
	return e.ID
}

func DecodeErr(err error) (id, msg string, data interface{}) {
	if err == nil {
		return OK.ID, OK.Msg, OK.Data
	}

	switch e := errors.Cause(err).(type) {
	case *Ecode:
		return e.ID, e.Msg, e.Data
	case Ecode:
		return e.ID, e.Msg, e.Data
	}

	return ServerErr.ID, err.Error(), nil
}
