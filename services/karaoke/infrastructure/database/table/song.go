package table

import "database/sql"

type Song struct {
	Id         int64
	ArtistId   int64
	Name       string
	Lyrics     string
	DamId      sql.Null[string]
	JoysoundId sql.Null[string]
}

func CreateTableSong(db *sql.DB) error {
	_, err := db.Exec(
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

	return err
}

func (s Song) Insert(db *sql.DB) error {
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
		s.ArtistId,
		s.Name,
		s.Lyrics,
		s.DamId,
		s.JoysoundId,
	)
	if err != nil {
		return err
	}

	return nil
}
