package service

import (
	"fmt"
	"githup.com/tally/internal/modal"
	"gorm.io/gorm"
)

type userServant struct {
	db *gorm.DB
}

func (d *userServant) CreateUser() {
	fmt.Println("method:CreateUser")
}

func (d *userServant) GetUser(serId int64) (*modal.TallyUsers, error) {
	fmt.Println("method:GetUser")
	return (&modal.TallyUsers{}), nil
}
