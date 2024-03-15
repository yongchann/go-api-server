package entity

import "time"

type Member struct {
	Id        int64 `gorm:"primaryKey, autoIncrement"`
	Nickname  string
	Password  string
	CreatedAt time.Time
}

type MemberRepository interface {
	Insert(member *Member) error
	FindByNickname(nickname string) (*Member, error)
}

type MemberUseCase interface {
	Join(member *Member) error
	FindByNickname(nickname string) (*Member, error)
	Login(nickname string, password string) (string, error)
}
