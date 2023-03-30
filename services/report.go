package services

import (
	"KayaKuy/models"
	"KayaKuy/repository"
)

type ReportService interface {
	GetAllReport(UserId int64) (models.Report, error)
}

type reportService struct {
	reportRepository repository.ReportRepository
}

func NewReportService(reportRepository repository.ReportRepository) *reportService {
	return &reportService{reportRepository}
}

func (a *reportService) GetAllReport(UserId int64) (models.Report, error) {
	var report models.Report

	err, history := a.reportRepository.GetReportHistory(UserId)
	if err != nil {
		return report, err
	}

	err, account := a.reportRepository.GetReportAccountBalance(UserId)
	if err != nil {
		return report, err
	}

	err, total := a.reportRepository.GetReportTotal(UserId)
	if err != nil {
		return report, err
	}

	report.Total = total
	report.AccountBalance = account
	report.History = history

	return report, nil
}
