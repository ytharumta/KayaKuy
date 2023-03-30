package controllers

import (
	"KayaKuy/models"
	"KayaKuy/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type journalHandler struct {
	journalService services.JournalService
}

func NewJournalHandler(journalService services.JournalService) *journalHandler {
	return &journalHandler{journalService}
}

func (b *journalHandler) GetAllJournal(c *gin.Context) {
	var (
		result gin.H
	)
	UserID := int64(c.MustGet("jwt_user_id").(float64))
	account, err := b.journalService.GetAllJournal(UserID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to get Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	} else {
		result = gin.H{
			"result": account,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (b *journalHandler) InsertJournal(c *gin.Context) {
	var journal models.Journal_entry

	err := c.ShouldBindJSON(&journal)
	journal.UserID = int64(c.MustGet("jwt_user_id").(float64))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to insert Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	err = b.journalService.InsertJournal(journal)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to insert Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Journal",
	})
}

func (a *journalHandler) UpdateJournal(c *gin.Context) {
	var journal models.Journal_entry
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to update Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	}
	journal.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := a.journalService.UpdateJournal(journal, int64(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to update Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success update journal",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to update Journal Entries",
		})

		c.Abort()
		return
	}

}

func (b *journalHandler) DeleteJournal(c *gin.Context) {
	var journal models.Journal_entry
	id, _ := strconv.Atoi(c.Param("id"))

	journal.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := b.journalService.DeleteJournal(journal, int64(id))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to delete Journal Entries",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success Delete Journal",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to delete Journal Entries",
		})

		c.Abort()
		return
	}
}
