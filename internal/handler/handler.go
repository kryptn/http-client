package handler

import (
	"github.com/kryptn/http-client/language"
)

type Result interface {
	Value() []byte
	HasValue() bool
	Error() error
}

type HandlerFunc func() Result

type BlockHandler func(language.Block) Result
type FunctionHandler func(Result, language.FunctionInvocation) Result
