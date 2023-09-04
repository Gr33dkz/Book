package data

type Book struct {
	Id          string  `json:"id"`
	Author      string  `json:"author"`
	Quantity    int32   `json:"quantity"`
	Price       float64 `json:"price"`
	ReleaseDate string  `json:"releaseDate"` // TODO MAKE FIELD time.Time
	Description string  `json:"description"`
}
