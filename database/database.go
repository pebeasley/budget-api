package database

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/pebeasley/leaves-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig) error {
	var err error
	dsn := "user=leaves" +
		" password=ocp*AW#1h95dVGrFK@9CT*%I" +
		" dbname=leaves" +
		" host=127.0.0.1" +
		" port=" +
		strconv.Itoa(5432) +
		" TimeZone=UTC"

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}

	// go seedDatabase()

	return err
}

func seedDatabase() {
	fmt.Println("SEEDING DATABASE...")
	dropTables()

	autoMigrateModels()

	fmt.Println("CREATING RECORDS...")
	bytes, err := bcrypt.GenerateFromPassword([]byte("randompassword"), 10)
	if err != nil {
		panic("Password creation failed")
	}
	DB.Create(&models.User{
		FirstName: "Evan",
		LastName:  "Beasley",
		Email:     "test@example.com",
		IsActive:  true,
		Password:  string(bytes),
	})
	var CreatedUser models.User
	DB.First(&CreatedUser)

	DB.Create(&models.Budget{Name: "Test", UserID: CreatedUser.ID})
	var CreatedBudget models.Budget
	DB.First(&CreatedBudget)

	DB.Create(&models.Account{Name: "test", Type: "Checking", BudgetId: CreatedBudget.ID})
	var CreatedAccount models.Account
	DB.First(&CreatedAccount)

	DB.Create(&models.Transaction{
		Amount:    64000.00,
		AccountID: CreatedAccount.ID,
		Memo:      "Starting Balance",
		IsIncome:  true,
	})

	for i := 0; i < 10; i++ {
		r := rand.Float64()
		r = r * 300
		r = math.Round(r*100) / 100
		DB.Create(&models.Transaction{
			Amount:    float32(r),
			AccountID: CreatedAccount.ID,
			Memo:      "Test",
			IsIncome:  false,
		})
	}

	fmt.Println("DATABASE FINISHED")
}

func dropTables() {
	fmt.Println("DROPPING TABLES...")
	DB.Migrator().DropTable("users")
	DB.Migrator().DropTable("accounts")
	DB.Migrator().DropTable("transactions")
	DB.Migrator().DropTable("budgets")
	DB.Migrator().DropTable("payees")
	DB.Migrator().DropTable("budget_configs")
	DB.Migrator().DropTable("categories")
}

func autoMigrateModels() {
	fmt.Println("CREATING TABLES...")
	DB.AutoMigrate(&models.User{},
		models.Transaction{},
		models.Account{},
		models.Payee{},
		models.BudgetConfig{},
		models.Category{})
}
