package member

import (
	"go-api-server/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) entity.MemberRepository {
	return &Repository{db: db}
}

func (r *Repository) Insert(member *entity.Member) error {
	result := r.db.Create(&member)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindByNickname(nickname string) (*entity.Member, error) {
	var member entity.Member
	result := r.db.Find(&member).Where("nickname", nickname)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &member, nil
}
