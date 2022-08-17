package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"movie/database"
	"movie/models"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Showlist struct {
	ShowName         string `json:"showName"`
	NumberOfSeats    int    `json:"numberOfSeats"`
	GoldenclassSeats string `json:"goldenclassSeats"`
	BalnconySeats    string `json:"balnconySeats"`
	FirstclassSeats  string `json:"firstclassSeats"`
	ShowTime         string `json:"showTime"`
	GoldClassPrice   int    `json:"goldClassPrice"`
	BalconyPrice     int    `json:"balconyPrice"`
	FirstClassPrice  int    `json:"firstClassPrice"`
}

func getSeatsLIstByRow(show models.Show) ([]int, []int, []int, map[int]string, map[int]string, map[int]string) {
	var a_rows = make(map[int]string)
	var b_rows = make(map[int]string)
	var c_rows = make(map[int]string)
	_ = json.Unmarshal(show.GoldClassSeats, &a_rows)
	_ = json.Unmarshal(show.BalconySeats, &b_rows)
	_ = json.Unmarshal(show.FirstClassSeats, &c_rows)
	var sorted_a_rows []int
	var sorted_b_rows []int
	var sorted_c_rows []int
	for key := range a_rows {
		sorted_a_rows = append(sorted_a_rows, key)
	}
	for key := range b_rows {
		sorted_b_rows = append(sorted_b_rows, key)
	}
	for key := range c_rows {
		sorted_c_rows = append(sorted_c_rows, key)
	}
	sort.Ints(sorted_a_rows)
	sort.Ints(sorted_b_rows)
	sort.Ints(sorted_c_rows)
	return sorted_a_rows, sorted_b_rows, sorted_c_rows, a_rows, b_rows, c_rows
}

func GetShows(context *gin.Context) {
	var show []models.Show
	var showlist []Showlist
	record := database.Instance.Find(&show)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	for i, j := range show {
		var temp Showlist
		temp.GoldClassPrice = 250
		temp.BalconyPrice = 200
		temp.FirstClassPrice = 150
		temp.ShowName = j.ShowName
		temp.NumberOfSeats = j.NumberOfSeats
		temp.ShowTime = j.ShowTime
		sorted_a_rows, sorted_b_rows, sorted_c_rows, a_rows, b_rows, c_rows := getSeatsLIstByRow(show[i])
		for i := 0; i < len(sorted_a_rows); i++ {
			if temp.GoldenclassSeats == "" {
				temp.GoldenclassSeats = a_rows[sorted_a_rows[i]]
			} else {
				temp.GoldenclassSeats += "," + a_rows[sorted_a_rows[i]]

			}
		}
		for i := 0; i < len(sorted_b_rows); i++ {
			if temp.BalnconySeats == "" {
				temp.BalnconySeats = b_rows[sorted_b_rows[i]]
			} else {
				temp.BalnconySeats += "," + b_rows[sorted_b_rows[i]]

			}
		}
		for i := 0; i < len(sorted_c_rows); i++ {
			if temp.FirstclassSeats == "" {
				temp.FirstclassSeats = c_rows[sorted_c_rows[i]]
			} else {
				temp.FirstclassSeats += "," + c_rows[sorted_c_rows[i]]

			}
		}
		showlist = append(showlist, temp)
	}
	context.JSON(http.StatusCreated, showlist)
}

func BookMovie(context *gin.Context) {
	var show models.Show
	var user models.User
	var bookedlist models.BookedList
	if err := context.ShouldBindJSON(&bookedlist); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//check if email exists
	record := database.Instance.Where("email = ?", bookedlist.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	if err := database.Instance.Where("show_name = ? and show_time = ?", bookedlist.ShowName, bookedlist.ShowTime).First(&show).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "show is not available!"})
		return
	}
	if show.NumberOfSeats-bookedlist.NumOfSeats < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "All tickets are sold out for show " + bookedlist.ShowName})
		return
	}
	booking_seats := strings.Split(bookedlist.Seats, ",")
	var a_rows = make(map[int]string)
	var b_rows = make(map[int]string)
	var c_rows = make(map[int]string)
	_ = json.Unmarshal(show.GoldClassSeats, &a_rows)
	_ = json.Unmarshal(show.BalconySeats, &b_rows)
	_ = json.Unmarshal(show.FirstClassSeats, &c_rows)
	var a_rows_avaiable, b_rows_avaiable, c_rows_avaiable []string
	var a_key_avaiable, b_key_avaiable, c_key_avaiable []int
	for k, v := range a_rows {
		a_rows_avaiable = append(a_rows_avaiable, v)
		a_key_avaiable = append(a_key_avaiable, k)
	}
	for k, v := range b_rows {
		b_rows_avaiable = append(b_rows_avaiable, v)
		b_key_avaiable = append(b_key_avaiable, k)
	}
	for k, v := range c_rows {
		c_rows_avaiable = append(c_rows_avaiable, v)
		c_key_avaiable = append(c_key_avaiable, k)
	}
	var unavaiable_seats string
	var a_row_bookedseats []int
	var b_row_bookedseats []int
	var c_row_bookedseats []int

	for i := 0; i < len(booking_seats); i++ {
		if strings.HasPrefix(booking_seats[i], "A") {
			if !slices.Contains(a_rows_avaiable, booking_seats[i]) {
				if unavaiable_seats == "" {
					unavaiable_seats += booking_seats[i]
				} else {
					unavaiable_seats += "," + booking_seats[i]
				}
			} else {
				index, _ := strconv.Atoi(strings.Split(booking_seats[i], "")[1])
				a_row_bookedseats = append(a_row_bookedseats, index)
			}
		} else if strings.HasPrefix(booking_seats[i], "B") {
			if !slices.Contains(b_rows_avaiable, booking_seats[i]) {
				if unavaiable_seats == "" {
					unavaiable_seats += booking_seats[i]
				} else {
					unavaiable_seats += "," + booking_seats[i]
				}
			} else {
				index, _ := strconv.Atoi(strings.Split(booking_seats[i], "")[1])
				b_row_bookedseats = append(b_row_bookedseats, index)
			}
		} else if strings.HasPrefix(booking_seats[i], "C") {
			if !slices.Contains(c_rows_avaiable, booking_seats[i]) {
				if unavaiable_seats == "" {
					unavaiable_seats += booking_seats[i]
				} else {
					unavaiable_seats += "," + booking_seats[i]
				}
			} else {
				index, _ := strconv.Atoi(strings.Split(booking_seats[i], "")[1])
				c_row_bookedseats = append(c_row_bookedseats, index)
			}
		}
	}
	if len(unavaiable_seats) == 0 {
		for _, val := range booking_seats {
			if strings.HasPrefix(val, "A") {
				bookedlist.Price += 250
			}
			if strings.HasPrefix(val, "B") {
				bookedlist.Price += 200
			} else {
				bookedlist.Price += 150
			}
		}
		bookingStatus := database.Instance.Create(&bookedlist)
		if bookingStatus.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": bookingStatus.Error.Error()})
			context.Abort()
			return
		}
		// Validate input
		show.NumberOfSeats = show.NumberOfSeats - bookedlist.NumOfSeats
		for i := range a_row_bookedseats {
			delete(a_rows, a_row_bookedseats[i])
		}
		for i := range b_row_bookedseats {
			delete(b_rows, b_row_bookedseats[i])
		}
		for i := range c_row_bookedseats {
			delete(c_rows, c_row_bookedseats[i])
		}
		show.GoldClassSeats, _ = json.Marshal(a_rows)
		show.BalconySeats, _ = json.Marshal(b_rows)
		show.FirstClassSeats, _ = json.Marshal(c_rows)

		database.Instance.Model(&show).Updates(show)

		context.JSON(http.StatusCreated, gin.H{"status": strconv.Itoa(bookedlist.NumOfSeats) + " tickets " + bookedlist.Seats + " for the movie " + bookedlist.ShowName + " at " + show.ShowTime + " has successfully booked, Total cost of Tickets are " + fmt.Sprintf("%v", bookedlist.Price) + " and booking id is " + strconv.Itoa(int(bookedlist.ID))})

	} else {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Tickets " + unavaiable_seats + " for the show " + show.ShowName + " are unavaiable"})

	}

}
