package queries

// use this struct to hold
// row data returned from queries
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}
