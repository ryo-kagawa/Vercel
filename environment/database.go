package environment

import (
	"errors"

	"github.com/ryo-kagawa/go-utils/commandline"
)

type EnvironmentDatabase struct {
	DATABASE_URL string `key:"DATABASE_URL"`
}

func GetEnvironmentDatabase() (EnvironmentDatabase, error) {
	return commandline.EnvironmentVariableParse[EnvironmentDatabase]()
}

func (e EnvironmentDatabase) Validate() error {
	if e.DATABASE_URL == "" {
		return errors.New("環境変数DATABASE_URLが未設定です")
	}

	return nil
}
