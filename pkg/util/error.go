package util

import (
	"fmt"

	"github.com/jonathanmoreiraa/planejja/pkg/log"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() error {
	return fmt.Errorf("%d: %s", e.Code, e.Message)
}

func ErrorWithMessage(err error, errMessage string) error {
	if err != nil {
		log.NewLogger().Error(err)
	}
	return fmt.Errorf("%s", errMessage)
}
