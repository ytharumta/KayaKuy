package services

import (
	"KayaKuy/helper"
	"KayaKuy/models"
	"KayaKuy/repository"
)

type JournalService interface {
	GetAllJournal(UserId int64) ([]models.Journal_entry_select, error)
	InsertJournal(journal models.Journal_entry) error
	UpdateJournal(inputJournal models.Journal_entry, id int64) (int64, error)
	DeleteJournal(inputJournal models.Journal_entry, id int64) (int64, error)
}

type journalService struct {
	journalRepository repository.JournalRepository
}

func NewJournalService(journalRepository repository.JournalRepository) *journalService {
	return &journalService{journalRepository}
}

func (a *journalService) GetAllJournal(UserId int64) ([]models.Journal_entry_select, error) {
	err, journal := a.journalRepository.GetAllJournal(UserId)
	if err != nil {
		return journal, err
	}

	return journal, nil
}

func (a *journalService) InsertJournal(inputJournal models.Journal_entry) error {
	var journal models.Journal_entry

	journal.UserID = inputJournal.UserID
	journal.Code = helper.GenerateCode(inputJournal.UserID)
	journal.CustomerId = inputJournal.CustomerId
	journal.AccountId = inputJournal.AccountId
	journal.Value = inputJournal.Value
	journal.Note = inputJournal.Note
	journal.TransactionType = inputJournal.TransactionType

	err := a.journalRepository.InsertJournal(journal)
	if err != nil {
		return err
	}

	return nil
}

func (a *journalService) UpdateJournal(inputJournal models.Journal_entry, id int64) (int64, error) {
	var journal models.Journal_entry

	journal.UserID = inputJournal.UserID
	journal.ID = id
	journal.CustomerId = inputJournal.CustomerId
	journal.AccountId = inputJournal.AccountId
	journal.Value = inputJournal.Value
	journal.Note = inputJournal.Note
	journal.TransactionType = inputJournal.TransactionType

	newJournal, err := a.journalRepository.UpdateJournal(journal)
	if err != nil {
		return newJournal, err
	}
	return newJournal, nil
}

func (a *journalService) DeleteJournal(inputJournal models.Journal_entry, id int64) (int64, error) {
	var journal models.Journal_entry

	journal.ID = id
	journal.UserID = inputJournal.UserID

	newJournal, err := a.journalRepository.DeleteJournal(journal)
	if err != nil {
		return newJournal, err
	}
	return newJournal, nil
}
