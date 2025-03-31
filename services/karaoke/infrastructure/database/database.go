package database

import (
	"database/sql"

	"github.com/ryo-kagawa/Vercel/environment"
	"github.com/ryo-kagawa/Vercel/infrastructure"
	"github.com/ryo-kagawa/Vercel/services/karaoke/domain/model"
	"github.com/ryo-kagawa/Vercel/services/karaoke/infrastructure/database/table"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(environment environment.EnvironmentDatabase) (Database, error) {
	db, err := infrastructure.NewDatabase(environment, "karaoke")
	if err != nil {
		return Database{}, err
	}
	return Database{
		DB: db,
	}, err
}

func (d Database) Dam() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.damId AS damId
FROM
	(
		SELECT
			artistId,
			name,
			left(lyrics, 50) AS lyrics,
			damId
		FROM
			song
		WHERE
			damId IS NOT NULL
		ORDER BY
			random()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	random()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		res := &model.KaraokeSong{}
		err = rows.Scan(
			&res.ArtistName,
			&res.SongName,
			&res.Lyrics,
			&res.DamId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(karaokeSongs, *res)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (d Database) Joysound() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.joysoundId AS joysoundId
FROM
	(
		SELECT
			artistId,
			name,
			left(lyrics, 50) AS lyrics,
			joysoundId
		FROM
			song
		WHERE
			joysoundId IS NOT NULL
		ORDER BY
			random()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	random()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		res := &model.KaraokeSong{}
		err = rows.Scan(
			&res.ArtistName,
			&res.SongName,
			&res.Lyrics,
			&res.JoysoundId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(karaokeSongs, *res)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (d Database) Ramdom() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.damId AS damId,
	song.joysoundId AS joysoundId
FROM
	(
		SELECT
			artistId,
			name,
			left(lyrics, 50) AS lyrics,
			damId,
			joysoundId
		FROM
			song
		ORDER BY
			random()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	random()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		type KaraokeSong struct {
			Artist table.Artist
			Song   table.Song
		}

		res := &KaraokeSong{}
		err = rows.Scan(
			&res.Artist.Name,
			&res.Song.Name,
			&res.Song.Lyrics,
			&res.Song.DamId,
			&res.Song.JoysoundId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(
			karaokeSongs,
			model.KaraokeSong{
				ArtistName: res.Artist.Name,
				SongName:   res.Song.Name,
				Lyrics:     res.Song.Lyrics,
				DamId:      res.Song.DamId.V,
				JoysoundId: res.Song.JoysoundId.V,
			},
		)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}
