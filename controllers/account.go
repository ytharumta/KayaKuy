package controllers

import (
	"KayaKuy/models"
	"KayaKuy/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type accountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *accountHandler {
	return &accountHandler{accountService}
}

func (b *accountHandler) GetAllAccount(c *gin.Context) {
	var (
		result gin.H
	)
	UserID := int64(c.MustGet("jwt_user_id").(float64))
	account, err := b.accountService.GetAllAccount(UserID)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": account,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (b *accountHandler) InsertAccount(c *gin.Context) {
	var account models.Account

	err := c.ShouldBindJSON(&account)
	account.UserID = int64(c.MustGet("jwt_user_id").(float64))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to Insert Account",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	err = b.accountService.InsertAccount(account)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to Insert Account",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Account",
	})
}

func (a *accountHandler) UpdateAccount(c *gin.Context) {
	var account models.Account
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&account)
	if err != nil {
		panic(err)
	}
	account.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := a.accountService.UpdateAccount(account, int64(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to Update Account",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success update Account",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to Update Account",
		})

		c.Abort()
		return
	}

}

func (b *accountHandler) DeleteAccount(c *gin.Context) {
	var account models.Account
	id, _ := strconv.Atoi(c.Param("id"))

	account.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := b.accountService.DeleteAccount(account, int64(id))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to delete Account",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success Delete Account",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to delete Account",
		})

		c.Abort()
		return
	}
}
