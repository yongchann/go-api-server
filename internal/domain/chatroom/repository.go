package chatroom

import (
	"go-api-server/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) entity.ChatroomRepository {
	return &Repository{db: db}
}

func (r *Repository) Insert(chatroom *entity.Chatroom) error {
	result := r.db.Create(chatroom)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindById(id int64) (*entity.Chatroom, error) {
	var chatroom entity.Chatroom
	result := r.db.Limit(1).Find(&chatroom, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &chatroom, nil
}

func (r *Repository) FindAll() ([]*entity.Chatroom, error) {
	var chatrooms []*entity.Chatroom
	result := r.db.Model(&entity.Chatroom{}).Scan(&chatrooms)
	if err := result.Error; err != nil {
		return nil, err
	}
	return chatrooms, nil
}

func (r *Repository) Update(id int64, chatroom *entity.Chatroom) error {
	chatroom.Id = id
	result := r.db.Save(chatroom)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id int64) error {
	result := r.db.Delete(&entity.Chatroom{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
