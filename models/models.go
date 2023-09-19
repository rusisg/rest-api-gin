package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func NewAlbums() []Album {
	albums := []Album{
		{ID: "1", Title: "Jitak", Artist: "Rasul Nurlybekov", Price: 20.21},
		{ID: "2", Title: "Luffy be my hoe", Artist: "Nami onepiece", Price: 14.12},
		{ID: "3", Title: "Season 10", Artist: "Fortnite battlepass", Price: 23.20},
		{ID: "4", Title: "Shart", Artist: "Rusisg", Price: 42},
		{ID: "5", Title: "SlaTt", Artist: "Playboi Carti", Price: 12.12},
	}
	return albums
}
