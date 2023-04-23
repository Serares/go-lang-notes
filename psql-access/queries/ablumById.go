package queries

import (
	"database/sql"
	"fmt"
)

func AlbumById(db *sql.DB, id int) (Album, error) {

	var alb Album

	row := db.QueryRow("SELECT * FROM album where id = 2")
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}
