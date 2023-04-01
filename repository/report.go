package repository

import (
	"KayaKuy/models"
	"database/sql"
)

type ReportRepository interface {
	GetReportHistory(UserId int64) (error, []models.History)
	GetReportAccountBalance(UserId int64) (error, []models.AccountBalance)
	GetReportTotal(UserId int64) (error, []models.Total)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepo(db *sql.DB) *reportRepository {
	return &reportRepository{db}
}

func (a *reportRepository) GetReportHistory(UserID int64) (err error, results []models.History) {
	sql := `SELECT je.code, je.note, c.name to_from, a.name account, je.transaction_type, coalesce(je.value,0) jvalue, je.created_at FROM journal_entries je
         join customers c on c.id = je.customer_id
         join accounts a on a.id = je.account_id
         join users u on u.id = je.user_id
         where je.user_id = $1
         order by je.created_at asc`

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.History{}
	}

	for rows.Next() {
		var history = models.History{}
		err = rows.Scan(&history.Code, &history.Note, &history.ToFrom, &history.Account, &history.TransactionType, &history.Value, &history.CreatedAt)
		if err != nil {
			return err, []models.History{}
		}
		results = append(results, history)
	}
	return
}

func (a *reportRepository) GetReportAccountBalance(UserID int64) (err error, results []models.AccountBalance) {
	sql := `select a."name" account_name, coalesce(je1.total,0) total_debit, coalesce(je2.total,0) total_kredit, coalesce((je1.total - je2.total),0) total from accounts a 
    left join (
    	select sum(je.value) total, je.user_id, je.account_id, je.transaction_type from journal_entries je group by je.account_id, je.user_id, je.transaction_type
    ) je1 on je1.account_id = a.id and je1.transaction_type = 'Debit'
    left join (
    	select sum(je.value) total, je.user_id, je.account_id, je.transaction_type from journal_entries je group by je.account_id, je.user_id, je.transaction_type
    ) je2 on je2.account_id = a.id and je2.transaction_type = 'Credit'
    where a.user_id = $1`

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.AccountBalance{}
	}

	for rows.Next() {
		var accountBalance = models.AccountBalance{}
		err = rows.Scan(&accountBalance.Account, &accountBalance.TotalDebit, &accountBalance.TotalKredit, &accountBalance.Total)
		if err != nil {
			return err, []models.AccountBalance{}
		}
		results = append(results, accountBalance)
	}
	return
}

func (a *reportRepository) GetReportTotal(UserID int64) (err error, results []models.Total) {
	sql := `select coalesce(sum(total_debit),0) total_debit, coalesce(sum(total_kredit),0) total_kredit, coalesce((sum(total_kredit)/sum(total_debit)*100),0) percentage,
   case 
   	when (sum(total_kredit)/sum(total_debit)*100) <= 10 then 'HEMAT! PERTAHANKAN'
   	when (sum(total_kredit)/sum(total_debit)*100) < 50 and (sum(total_kredit)/sum(total_debit)*100) >= 20 then 'JANGAN TERLALU BOROS!'
   	when (sum(total_kredit)/sum(total_debit)*100) >= 50 then 'BOROS!'
   end note
   from (
   select a."name" account_name, coalesce(je1.total,0) total_debit, coalesce(je2.total,0) total_kredit, (je1.total - je2.total) total, (je2.total/je1.total*100) from accounts a 
    left join (
    	select sum(je.value) total, je.user_id, je.account_id, je.transaction_type from journal_entries je group by je.account_id, je.user_id, je.transaction_type
    ) je1 on je1.account_id = a.id and je1.transaction_type = 'Debit'
    left join (
    	select sum(je.value) total, je.user_id, je.account_id, je.transaction_type from journal_entries je group by je.account_id, je.user_id, je.transaction_type
    ) je2 on je2.account_id = a.id and je2.transaction_type = 'Credit'
    where a.user_id = $1) a`

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.Total{}
	}

	for rows.Next() {
		var total = models.Total{}
		err = rows.Scan(&total.TotalDebit, &total.TotalKredit, &total.Precentage, &total.Note)
		if err != nil {
			return err, []models.Total{}
		}
		results = append(results, total)
	}
	return
}
