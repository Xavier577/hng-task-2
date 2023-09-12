package env

import (
	"github.com/Xavier577/hng-task-2/pkg/types"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

type Value interface {
	string | types.Number
}

func Get(key string) string {
	return os.Getenv(key)
}
