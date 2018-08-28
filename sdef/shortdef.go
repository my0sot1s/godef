package sdef

import (
	"strings"
	"time"
)

// M is map any
type M = map[string]interface{}

// MS is map string
type MS = map[string]string

// TD is time duration
type TD = time.Duration

// CB is define callback
type CB = func(any ...interface{})

// StrRep is define to shooter
var StrRep = strings.Replace

// ErrMsg is alias ErrorMessage
type ErrMsg = ErrorMessage

// ErrorMessage is define Message Error
type ErrorMessage struct {
	code        int
	err         error
	description string
}

// OKMsg is alias SuccessMessage
type OKMsg = SuccessMessage

// SuccessMessage is define Message Error
type SuccessMessage struct {
	code int
	data interface{}
}
