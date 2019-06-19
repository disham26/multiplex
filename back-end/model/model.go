package model

import (
	"TechVeritasMultiplex/structs"
	"strconv"
)

var Movies []structs.Movies
var Multiplex []structs.Shows

func DBInit() {

	Movies = []structs.Movies{
		structs.Movies{
			Name:        "Thackery",
			Description: "Starring Nawazuddin Siddique",
			ImageUrl:    "Thackery.jpg",
		},
		structs.Movies{
			Name:        "Manikarnika",
			Description: "Starring Kangana",
			ImageUrl:    "Manikarnika.jpg",
		},
		structs.Movies{
			Name:        "URI-The Surgical Strike",
			Description: "Starring Vicky Kaushal",
			ImageUrl:    "URI-poster.jpg",
		},
	}
	customer := structs.Shows{
		ShowID: "1",
		Rows: []structs.SeatType{
			structs.SeatType{
				Name:  "A",
				Price: 320,
				Seats: []structs.Seat{structs.Seat{
					SeatNo: "A1",
					Taken:  false, Active: false,
				},
					structs.Seat{
						SeatNo: "A2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A3",
						Taken:  false, Active: false,
					},
					structs.Seat{
						SeatNo: "A4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A5",
						Taken:  false, Active: false,
					},
					structs.Seat{
						SeatNo: "A6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A7",
						Taken:  false, Active: false,
					},
					structs.Seat{
						SeatNo: "A8",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A9",
						Taken:  false, Active: false,
					},
				},
			},
			structs.SeatType{
				Name:  "B",
				Price: 280,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "B1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B6",
						Taken:  false, Active: true,
					},
				},
			},
			structs.SeatType{
				Name:  "C",
				Price: 240,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "C1",
						Taken:  false,
						Active: false,
					},
					structs.Seat{
						SeatNo: "C2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C7",
						Taken:  false, Active: true,
					},
				},
			},
		},
	}
	Multiplex = append(Multiplex, customer)
	//InsertObject(driver, customer, 1)
	customer = structs.Shows{
		ShowID: "2",
		Rows: []structs.SeatType{
			structs.SeatType{
				Name:  "A",
				Price: 320,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "A1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A7",
						Taken:  false, Active: true,
					},
				},
			},
			structs.SeatType{
				Name:  "B",
				Price: 280,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "B1",
						Taken:  false, Active: false,
					},
					structs.Seat{
						SeatNo: "B2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B6",
						Taken:  false, Active: true,
					},
				},
			},
			structs.SeatType{
				Name:  "C",
				Price: 240,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "C1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C7",
						Taken:  false, Active: true,
					},
				},
			},
		},
	}
	Multiplex = append(Multiplex, customer)
	//InsertObject(driver, customer, 2)
	customer = structs.Shows{
		ShowID: "3",
		Rows: []structs.SeatType{
			structs.SeatType{
				Name:  "A",
				Price: 320,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "A1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "A5",
						Taken:  false, Active: true,
					},
				},
			},
			structs.SeatType{
				Name:  "B",
				Price: 280,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "B1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B7",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "B8",
						Taken:  false, Active: true,
					},
				},
			},
			structs.SeatType{
				Name:  "C",
				Price: 240,
				Seats: []structs.Seat{
					structs.Seat{
						SeatNo: "C1",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C2",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C3",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C4",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C5",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C6",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C7",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C8",
						Taken:  false, Active: true,
					},
					structs.Seat{
						SeatNo: "C9",
						Taken:  false, Active: true,
					},
				},
			},
		},
	}
	Multiplex = append(Multiplex, customer)

}

//UpdateVariable function to update the seats in the variable
func UpdateVariable(showID int, seats []string) structs.Shows {
	rowSelected := 0
	for _, element := range seats {
		row := element[0:1]
		seat := element[1:]
		seatNumber, _ := strconv.Atoi(seat)
		seatNumber--

		switch row {
		case "A":
			rowSelected = 0

			break
		case "B":
			rowSelected = 1

			break
		case "C":
			rowSelected = 2

			break

		}

		Multiplex[showID].Rows[rowSelected].Seats[seatNumber].Taken = true

	}
	return Multiplex[showID]
}
