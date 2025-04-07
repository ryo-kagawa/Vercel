package main

import (
	"github.com/ryo-kagawa/go-utils/commandline"
)

type EnvironmentInitialize struct {
}

type Environment struct {
	DATABASE_URL                 string `key:"DATABASE_URL"`
	DATABASE_URL_INITIALZE       string `key:"DATABASE_URL_INITIALZE"`
	DATABASE_INITIALIZE_DATABASE bool   `key:"DATABASE_INITIALIZE_DATABASE"`
}

func GetEnvironment() (Environment, error) {
	return commandline.EnvironmentVariableParse[Environment]()
}
