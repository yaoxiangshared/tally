package service

import "githup.com/tally/internal/modal"

type UserService interface {
	CreateUser()
	GetUser(userId int64) (*modal.TallyUsers, error)
}
