package karaoke

import "github.com/ryo-kagawa/Vercel/environment"

type Environment struct {
	Database environment.EnvironmentDatabase
	Line     environment.EnvironmentLine
}

func (k Karaoke) GetEnvironment() (Environment, error) {
	database, err := environment.GetEnvironmentDatabase()
	if err != nil {
		return Environment{}, err
	}
	line, err := environment.GetEnvironmentLine()
	if err != nil {
		return Environment{}, err
	}
	return Environment{
		Database: database,
		Line:     line,
	}, nil
}

func (e Environment) Validate() error {
	if err := e.Database.Validate(); err != nil {
		return err
	}
	if err := e.Line.Validate(); err != nil {
		return err
	}
	return nil
}
