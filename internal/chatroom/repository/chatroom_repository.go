package repository

import (
	"github.com/pkg/errors"
	"go-api-server/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found record")
)

type MySqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MySqlRepository {
	return &MySqlRepository{db: db}

}

func (m MySqlRepository) Insert(chatroom *domain.Chatroom) error {
	result := m.db.Create(chatroom)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (m MySqlRepository) FindById(id int64) (*domain.Chatroom, error) {
	var chatroom domain.Chatroom
	result := m.db.First(&chatroom, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &chatroom, nil
}

func (m MySqlRepository) FindAll() ([]*domain.Chatroom, error) {
	var chatrooms []*domain.Chatroom
	result := m.db.Model(&domain.Chatroom{}).Scan(&chatrooms)
	if err := result.Error; err != nil {
		return nil, err
	}
	return chatrooms, nil
}

func (m MySqlRepository) Update(id int64, chatroom *domain.Chatroom) error {
	chatroom.Id = id
	result := m.db.Save(chatroom)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (m MySqlRepository) Delete(id int64) error {
	result := m.db.Delete(&domain.Chatroom{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
