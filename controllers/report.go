package controllers

import (
	"KayaKuy/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reportHandler struct {
	reportService services.ReportService
}

func NewReportHandler(reportService services.ReportService) *reportHandler {
	return &reportHandler{reportService}
}

func (b *reportHandler) GetReport(c *gin.Context) {
	var (
		result gin.H
	)
	UserID := int64(c.MustGet("jwt_user_id").(float64))
	report, err := b.reportService.GetAllReport(UserID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to get Report",
			"message": err.Error(),
		})

		c.Abort()
		return
	} else {
		result = gin.H{
			"result": report,
		}
	}

	c.JSON(http.StatusOK, result)
}
