package controllers

import (
	"KayaKuy/models"
	"KayaKuy/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type customerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *customerHandler {
	return &customerHandler{customerService}
}

func (b *customerHandler) GetAllCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	UserID := int64(c.MustGet("jwt_user_id").(float64))
	account, err := b.customerService.GetAllCustomer(UserID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to get Customer",
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

func (b *customerHandler) InsertCustomer(c *gin.Context) {
	var customer models.Customer

	err := c.ShouldBindJSON(&customer)
	customer.UserID = int64(c.MustGet("jwt_user_id").(float64))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to insert Customer",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	err = b.customerService.InsertCustomer(customer)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to insert Customer",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Customer",
	})
}

func (a *customerHandler) UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to update Customer",
			"message": err.Error(),
		})

		c.Abort()
		return
	}
	customer.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := a.customerService.UpdateCustomer(customer, int64(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to update Customer",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success update customer",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "Data is forbidden",
		})
	}

}

func (b *customerHandler) DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	id, _ := strconv.Atoi(c.Param("id"))

	customer.UserID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := b.customerService.DeleteCustomer(customer, int64(id))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to delete Customer",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if ct > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success Delete customer",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to delete Customer",
		})

		c.Abort()
		return
	}
}
