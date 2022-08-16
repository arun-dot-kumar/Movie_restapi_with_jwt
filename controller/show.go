package controllers

import (
	"github.com/gin-gonic/gin"
	"movie/database"
	"movie/models"
	"net/http"
)

func AddMovie(context *gin.Context) {
	var show = models.Show{NumberOfSeats: 30,
		GoldClassSeats:  []byte(`{"1":"A1", "2":"A2", "3":"A3","4":"A4", "5":"A5","6":"A6","7":"A7","8": "A8", "9": "A9", "10":"A10"}`),
		BalconySeats:    []byte(`{"1":"B1", "2":"B2", "3":"B3","4":"B4", "5":"B5","6":"B6","7":"B7","8": "B8", "9": "B9", "10":"B10"}`),
		FirstClassSeats: []byte(`{"1":"C1", "2":"C2", "3":"C3","4":"C4", "5":"C5","6":"C6","7":"C7","8": "C8", "9": "C9", "10":"C10"}`)}
	var oldShow models.Show
	if err := context.ShouldBindJSON(&show); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Where("show_name = ?", show.ShowName).First(&oldShow)
	if record.Error == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	if oldShow.ShowName != "" {
		context.JSON(http.StatusOK, gin.H{"error": "show is already added"})
		context.Abort()
		return
	}
	var showTime = []string{"10:00 AM", "2:00 PM", "5:00 PM", "9:00 PM"}
	for _, value := range showTime {
		var newshow models.Show
		newshow.ShowName = show.ShowName
		newshow.GoldClassSeats = show.GoldClassSeats
		newshow.BalconySeats = show.BalconySeats
		newshow.FirstClassSeats = show.FirstClassSeats
		newshow.ShowTime = value
		newshow.NumberOfSeats = show.NumberOfSeats
		bookingStatus := database.Instance.Create(&newshow)
		if bookingStatus.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": bookingStatus.Error.Error()})
			context.Abort()
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"ShowId": show.ID, "ShowName": show.ShowName, "ShowTime": showTime})

}

func Getbookings(context *gin.Context) {
	var booklist []models.BookedList
	record := database.Instance.Find(&booklist)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, booklist)
}
