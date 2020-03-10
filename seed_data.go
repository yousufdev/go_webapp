package main

import "github.com/yousufdev/go_webapp/model"

func getUsers() []model.User {
	users := []model.User{
		model.User{
			Name:     "yousuf",
			Email:    "abcd@gmail.com",
			Password: "12345",
			Role:     model.RoleAdmin,
		},
		model.User{
			Name:     "ahmed",
			Email:    "zxcv@gmail.com",
			Password: "password",
			Role:     model.RoleAnalyst,
		},
	}

	return users
}
