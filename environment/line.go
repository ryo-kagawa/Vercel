package environment

import (
	"errors"

	"github.com/ryo-kagawa/go-utils/commandline"
)

type EnvironmentLine struct {
	LINE_CHANNEL_SECRET string `key:"LINE_CHANNEL_SECRET"`
	LINE_CHANNEL_TOKEN  string `key:"LINE_CHANNEL_TOKEN"`
}

func GetEnvironmentLine() (EnvironmentLine, error) {
	return commandline.EnvironmentVariableParse[EnvironmentLine]()
}

func (env EnvironmentLine) Validate() error {
	if env.LINE_CHANNEL_SECRET == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}
	if env.LINE_CHANNEL_TOKEN == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}

	return nil
}
