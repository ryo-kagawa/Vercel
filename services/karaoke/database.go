package karaoke

import (
	karaokedatabase "github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/Vercel/domain"
	"github.com/ryo-kagawa/Vercel/environment"
)

func (k Karaoke) NewDatabase(environment environment.EnvironmentDatabase) (karaokedatabase.Database, error) {
	db, err := domain.NewDatabase(environment, "karaoke")
	if err != nil {
		return karaokedatabase.Database{}, err
	}
	return karaokedatabase.Database{
		DB: db,
	}, err
}
