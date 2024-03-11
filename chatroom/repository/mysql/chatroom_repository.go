package mysql

import (
	"go-api-server/domain"
)

type MySqlRepository struct{}

func NewMysqlRepository() (*MySqlRepository, error) {
	return nil, nil
}

func (m *MySqlRepository) Create(chatroom domain.Chatroom) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySqlRepository) GetById(id int64) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySqlRepository) Fetch(offset, limit int) ([]*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySqlRepository) Update(id int64, chatroom domain.Chatroom) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySqlRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
