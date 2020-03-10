package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yousufdev/go_webapp/app"
	"github.com/yousufdev/go_webapp/model"
	"github.com/yousufdev/go_webapp/store"
)

// gorm naming convention
// https://gorm.io/docs/conventions.html

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "postgres"
	password = "docker"
	sslMode  = "disable"
)

type relation struct {
	model    interface{}
	field    string
	dest     string
	onDelete string
	onUpdate string
}

func initializeApp() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslMode)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// db.LogMode(true) // remove this line only for debugging

	if err := db.Transaction(migrate); err != nil {
		return db, err
	}

	return db, nil
}

func main() {
	db, err := initializeApp()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	userStore := store.NewUserStore(db)
	users := getUsers()

	for _, u := range users {
		if err := userStore.Create(&u); err != nil {
			panic(err)
		}
	}

	user, _ := userStore.FindOne("id = ?", 1)
	// fmt.Printf("%+v", user)
	// fmt.Println(user.Role.Name)

	app := &app.App{
		UserStore: userStore,
		User:      user,
	}

	if app.User.Role.Name == model.RoleAdmin.Name {
		users, _ := userStore.FindAll()
		for _, u := range users {
			fmt.Println(u.Email)
		}
	}

}

func migrate(tx *gorm.DB) error {
	models := []interface{}{
		&model.User{},
		&model.ClientTarget{},
		&model.Target{},
		&model.Client{},
		&model.RolePermission{},
		&model.Permission{},
		&model.Role{},
	}

	relations := []relation{
		{&model.User{}, "client_id", "clients(id)", "RESTRICT", "CASCADE"},
		{&model.User{}, "role_id", "roles(id)", "RESTRICT", "CASCADE"},
		{&model.RolePermission{}, "role_id", "roles(id)", "RESTRICT", "CASCADE"},
		{&model.RolePermission{}, "permission_id", "permissions(id)", "RESTRICT", "CASCADE"},
		{&model.ClientTarget{}, "client_id", "clients(id)", "RESTRICT", "CASCADE"},
		{&model.ClientTarget{}, "target_id", "targets(id)", "RESTRICT", "CASCADE"},
	}

	if err := tx.DropTableIfExists(models...).Error; err != nil {
		return fmt.Errorf("drop table: %w", err)
	}
	if err := tx.CreateTable(models...).Error; err != nil {
		return fmt.Errorf("create table: %w", err)
	}

	for _, r := range relations {
		if err := tx.Model(r.model).AddForeignKey(r.field, r.dest, r.onDelete, r.onUpdate).Error; err != nil {
			return fmt.Errorf("%T constraint (%s): %w", r.model, r.field, err)
		}
	}

	return nil
}
