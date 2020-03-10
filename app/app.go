package app

import (
	"github.com/yousufdev/go_webapp/model"
	"github.com/yousufdev/go_webapp/store"
)

// App ...
type App struct {
	UserStore store.UserStore
	User      *model.User
}
