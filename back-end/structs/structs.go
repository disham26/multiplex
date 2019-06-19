package structs

type Seat struct {
	SeatNo string `json:"seat_no"`
	Taken  bool   `json:"taken"`
	Active bool   `json:active`
}
type Row struct {
	Seats []Seat `json:"seat"`
}
type Show struct {
	Rows []Row `json:"row"`
}

type Movies struct {
	Name        string `json:"movie_name"`
	Description string `json:"description"`
	ImageUrl    string `json:"url"`
}
type Shows struct {
	ShowID string     `json:"showid"`
	Rows   []SeatType `json:"type"`
}
type SeatType struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Seats []Seat `json:"seat"`
}

func (c Shows) ID() (jsonField string, value interface{}) {
	value = c.ShowID
	jsonField = "showid"
	return
}

//Page Sample struct
type Page struct {
	Title string
	Text  string
	Config string
	Shows  string
}

//BookResponse struct is used as an XHR response
type BookResponse struct {
	Success bool  `json:"success"`
	Shows   Shows `json:"shows"`
}
