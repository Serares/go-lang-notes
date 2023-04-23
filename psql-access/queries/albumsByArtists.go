package queries

import (
	"database/sql"
	"fmt"
)

func AlbumsByArtists(db *sql.DB, name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album
	// sending the parameters separated oposed to string concatenation
	// can prevent sqlI attacks
	// TODO how to query this in psql using the quotes?
	rows, err := db.Query("SELECT * FROM album WHERE artist = 'John Coltrane'")
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		// using rows.scan you will loop through each column and
		// assign it to the Album struct
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
