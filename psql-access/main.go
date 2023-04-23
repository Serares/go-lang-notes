package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"example.com/psql-access/queries"
	_ "github.com/lib/pq"
)

func main() {

	var databaseSource = fmt.Sprintf("postgres://%s:%s@localhost:5432/go-tutorial?sslmode=disable", os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"))
	db, err := sql.Open("postgres", databaseSource)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected success!")

	// query albums by artists
	albums, err := queries.AlbumsByArtists(db, "John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	album, err := queries.AlbumById(db, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", album)

	insertAlbum := queries.Album{
		Title:  "Go fast",
		Artist: "The focker",
		Price:  54.22,
	}

	id, err := queries.AddAlbum(db, &insertAlbum)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album inserted with ID %v\n", id)

}
