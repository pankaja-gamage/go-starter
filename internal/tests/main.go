package tests

import (
	"go-starter/internal/common/data"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Test struct {
}

func New() *Test {
	e := godotenv.Load()
	if e != nil {
		log.Panic(e)
	}

	if len(os.Getenv("DB_USR_PW")) == 0 || len(os.Getenv("DB_USR")) == 0 {
		log.Panic("DB user and password should be available in env. " +
			"Ex: env DB_USR_PW, env DB_USR ")
	}

	data.InitDb()

	return &Test{}
}
