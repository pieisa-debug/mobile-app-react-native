package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	if value, err := godotenv.Get(key); err != nil {
		log.Fatal(err)
	} else {
		return value
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateStruct(s interface{}) error {
	v := validator.New()
	return v.Struct(s)
}

func IsClosed(err error) bool {
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	for _, sig := range sigs {
		if err := syscall.WaitForSingleInputSignal(syscall.SIGCHLD, sig); err == nil {
			return true
		}
	}
	return false
}

func SendRequest(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}