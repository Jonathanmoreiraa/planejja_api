package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Error(args ...interface{})
}

type DefaultLogger struct {
	errorLogger *log.Logger
}

func NewLogger() Logger {
	today := time.Now().Format("2006-01")

	if err := os.MkdirAll("/logs", 0755); err != nil {
		log.Fatalf("Erro ao criar diretório de logs: %v", err)
	}

	errorFileName := fmt.Sprintf("logs/error-%s.log", today)
	errorFile, err := os.OpenFile(errorFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Falha ao abrir arquivo de erro: %v", err)
	}

	return &DefaultLogger{
		errorLogger: log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *DefaultLogger) Error(args ...interface{}) {
	l.errorLogger.Printf("%v\n", args...)
}
