package service

import (
	"fmt"
	"githup.com/tally/internal/modal"
	"gorm.io/gorm"
)

type UserServant struct {
	Db *gorm.DB
}

func (d *UserServant) CreateUser() {
	fmt.Println("method:CreateUser")
}

func (d *UserServant) GetUser(serId int64) (*modal.TallyUsers, error) {
	fmt.Println("method:GetUser")
	//return (&modal.TallyUsers{}), nil
	user := modal.TallyUsers{}
	result := d.Db.First(&user)
	fmt.Println(result.RowsAffected)
	return &user, nil
}
