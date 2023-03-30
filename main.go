package main

import (
	"KayaKuy/controllers"
	"KayaKuy/database"
	"KayaKuy/middleware"
	"KayaKuy/repository"
	"KayaKuy/services"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//Router Gin
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		userRepository := repository.NewUserRepo(DB)
		userService := services.NewUserService(userRepository)
		userHandler := controllers.NewUserHandler(userService)
		v1.GET("/auth/:provider", userHandler.RedirectHandler)
		v1.GET("/auth/:provider/callback", userHandler.CallbackHandler)
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)
		v1.PUT("/user/update", middleware.IsAuth(), userHandler.UpdateUser)

		accountRepository := repository.NewAccountRepo(DB)
		accountService := services.NewAccountService(accountRepository)
		accountHandler := controllers.NewAccountHandler(accountService)
		account := v1.Group("account")
		{
			account.GET("/", middleware.IsAuth(), accountHandler.GetAllAccount)
			account.POST("/", middleware.IsAuth(), accountHandler.InsertAccount)
			account.PUT("/:id", middleware.IsAuth(), accountHandler.UpdateAccount)
			account.DELETE("/:id", middleware.IsAuth(), accountHandler.DeleteAccount)
		}

		customerRepository := repository.NewCustomerRepo(DB)
		customerService := services.NewCustomerService(customerRepository)
		customerHandler := controllers.NewCustomerHandler(customerService)
		customer := v1.Group("customer")
		{
			customer.GET("/", middleware.IsAuth(), customerHandler.GetAllCustomer)
			customer.POST("/", middleware.IsAuth(), customerHandler.InsertCustomer)
			customer.PUT("/:id", middleware.IsAuth(), customerHandler.UpdateCustomer)
			customer.DELETE("/:id", middleware.IsAuth(), customerHandler.DeleteCustomer)
		}

		journalRepository := repository.NewJournalRepo(DB)
		journalService := services.NewJournalService(journalRepository)
		journalHandler := controllers.NewJournalHandler(journalService)
		journal := v1.Group("journal")
		{
			journal.GET("/", middleware.IsAuth(), journalHandler.GetAllJournal)
			journal.POST("/", middleware.IsAuth(), journalHandler.InsertJournal)
			journal.PUT("/:id", middleware.IsAuth(), journalHandler.UpdateJournal)
			journal.DELETE("/:id", middleware.IsAuth(), journalHandler.DeleteJournal)
		}

		reportRepository := repository.NewReportRepo(DB)
		reportService := services.NewReportService(reportRepository)
		reportHandler := controllers.NewReportHandler(reportService)
		report := v1.Group("report")
		{
			report.GET("/", middleware.IsAuth(), reportHandler.GetReport)
		}
	}

	router.Run("localhost:8080")
}
