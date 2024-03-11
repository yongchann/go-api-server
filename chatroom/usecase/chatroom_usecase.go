package usecase

import "go-api-server/domain"

type ChatroomUseCase struct {
	chatroomRepository domain.ChatroomRepository
}

func NewChatroomUseCase(repo domain.ChatroomRepository) domain.ChatroomUseCase {
	return &ChatroomUseCase{chatroomRepository: repo}
}

func (c *ChatroomUseCase) Create(chatroom domain.Chatroom) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatroomUseCase) GetById(id int64) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatroomUseCase) Fetch(offset, limit int) ([]*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatroomUseCase) Update(id int64, chatroom domain.Chatroom) (*domain.Chatroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatroomUseCase) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
