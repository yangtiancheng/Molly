package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type DsnInfo struct {
	Type     string
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

func TestMain(dsn DsnInfo) {
	var db *gorm.DB
	var err error

	if dsn.Type == "postgres" {
		db, err = getPgDb(dsn)
		if err != nil {
			fmt.Printf("Failed to connect to PG database: %v\n", err)
			return
		}
	} else if dsn.Type == "mysql" {
		db, err = getMySQLDb(dsn)
		if err != nil {
			fmt.Printf("Failed to connect to MySQL database: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Unsupported database type: %s\n", dsn.Type)
		return
	}

	if CheckTableExist(db, "users") {
		fmt.Println("Table users already exists")
	} else {
		fmt.Println("Table users does not exist")
		db.AutoMigrate(&User{})
		fmt.Println("Table users table created")
	}

	// TestInsert(db)
	// TestSelect(db)
	// TestUpdate(db)
	TestDelete(db)
}

func getPgDb(dsn DsnInfo) (db *gorm.DB, err error) {
	dsnStr := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	db, err = gorm.Open(postgres.Open(fmt.Sprintf(dsnStr, dsn.Host, dsn.User, dsn.Password, dsn.Dbname, dsn.Port)), &gorm.Config{})
	return db, err
}
func getMySQLDb(dsn DsnInfo) (db *gorm.DB, err error) {
	dsnStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(fmt.Sprintf(dsnStr, dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.Dbname)), &gorm.Config{})
	return db, err
}

func TestInsert(db *gorm.DB) {
	db.Create(&User{Name: "John", Email: "john@example.com"})
	db.Create(&User{Name: "Jane", Email: "jane@example.com"})
	db.Create(&User{Name: "Jim", Email: "jim@example.com"})
	db.Create(&User{Name: "Jill", Email: "jill@example.com"})
}

func TestSelect(db *gorm.DB) {
	var users []User
	db.Find(&users)
	fmt.Println(users)
}

func TestDelete(db *gorm.DB) {
	db.Delete(&User{Name: "John"}, "name = ?", "John")
	db.Delete(&User{Name: "Jane"}, "name = ?", "Jane")
}
func TestUpdate(db *gorm.DB) {
	db.Model(&User{}).Where("name = ?", "John").Update("email", "john1@example.com")
	db.Model(&User{}).Where("name = ?", "Jane").Update("email", "jane1@example.com")
}

func CheckTableExist(db *gorm.DB, tableName string) bool {
	return db.Migrator().HasTable(tableName)
}

func main() {
	dsn := DsnInfo{
		Type:     "postgres",
		Host:     "localhost",
		User:     "postgres",
		Password: "123456",
		Dbname:   "postgres",
		Port:     "5432",
	}
	TestMain(dsn)
}
