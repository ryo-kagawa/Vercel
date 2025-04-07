package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ryo-kagawa/Vercel/infrastructure"
	"github.com/ryo-kagawa/Vercel/services/karaoke/infrastructure/database/table"
)

const karaokeDataFile = "karaoke-data.json"

type KaraokeSongList []struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	DamId      *string `json:"damId"`
	JoysoundId *string `json:"joysoundId"`
	Songs      []struct {
		Id         int64   `json:"id"`
		Name       string  `json:"name"`
		Lyrics     string  `json:"lyrics"`
		DamId      *string `json:"damId"`
		JoysoundId *string `json:"joysoundId"`
	} `json:"songs"`
}

func createSqlNull[T any](value *T) sql.Null[T] {
	valid := value != nil
	v := *new(T)
	if valid {
		v = *value
	}
	return sql.Null[T]{
		V:     v,
		Valid: valid,
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("initialize finish")
}

func run() error {
	environment, err := GetEnvironment()
	if err != nil {
		return err
	}

	{
		if environment.DATABASE_INITIALIZE_DATABASE {
			db, err := infrastructure.NewDatabase(environment.DATABASE_URL_INITIALZE+"?sslmode=disable", "")
			if err != nil {
				return err
			}
			_, err = db.Exec("DROP DATABASE IF EXISTS " + "vercel")
			if err != nil {
				return err
			}
			_, err = db.Exec("CREATE DATABASE " + "vercel")
			if err != nil {
				return err
			}
		}
		db, err := infrastructure.NewDatabase(environment.DATABASE_URL, "")
		if err != nil {
			return err
		}
		_, err = db.Exec("DROP SCHEMA IF EXISTS " + "karaoke")
		if err != nil {
			return err
		}
		_, err = db.Exec("CREATE SCHEMA " + "karaoke")
		if err != nil {
			return err
		}
	}

	db, err := infrastructure.NewDatabase(environment.DATABASE_URL, "karaoke")
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`
CREATE TABLE artist(
id         bigserial             NOT NULL,
name       character varying(50) NOT NULL,
damId      character varying(6)      NULL,
joysoundId character varying(6)      NULL,

PRIMARY KEY(id)
)
`,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`
CREATE TABLE song(
id         bigserial               NOT NULL,
artistId   bigint                  NOT NULL,
name       character varying(50)   NOT NULL,
lyrics     character varying(1000) NOT NULL,
damId      character varying(7)    NULL UNIQUE,
joysoundId character varying(6)    NULL UNIQUE,

PRIMARY KEY(id),
FOREIGN KEY (artistId) REFERENCES artist(id)
)
`,
	)
	if err != nil {
		return err
	}

	file, err := os.Open(karaokeDataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	binary, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	jsonData := KaraokeSongList{}
	if err := json.Unmarshal(binary, &jsonData); err != nil {
		return err
	}
	for _, jsonArtist := range jsonData {
		artist := table.Artist{
			Name:       jsonArtist.Name,
			DamId:      createSqlNull(jsonArtist.DamId),
			JoysoundId: createSqlNull(jsonArtist.JoysoundId),
		}
		_, err := db.Exec(
			`
INSERT INTO artist
(
	name,
	damId,
	joysoundId
)VALUES(
	$1,
	$2,
	$3
)
			`,
			artist.Name,
			artist.DamId,
			artist.JoysoundId,
		)
		if err != nil {
			return err
		}
		artistId := int64(0)
		if err := db.QueryRow(`SELECT id FROM artist WHERE name=$1`, artist.Name).Scan(&artistId); err != nil {
			return err
		}
		for _, jsonSong := range jsonArtist.Songs {
			song := table.Song{
				ArtistId:   artistId,
				Name:       jsonSong.Name,
				Lyrics:     jsonSong.Lyrics,
				DamId:      createSqlNull(jsonSong.DamId),
				JoysoundId: createSqlNull(jsonSong.JoysoundId),
			}
			_, err := db.Exec(
				`
INSERT INTO song
(
	artistId,
	name,
	lyrics,
	damId,
	joysoundId
)VALUES(
	$1,
	$2,
	$3,
	$4,
	$5
)
				`,
				song.ArtistId,
				song.Name,
				song.Lyrics,
				song.DamId,
				song.JoysoundId,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
