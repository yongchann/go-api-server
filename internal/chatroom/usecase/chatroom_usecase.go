package usecase

import (
	"go-api-server/internal/domain"
)

type ChatroomUseCase struct {
	chatroomRepository domain.ChatroomRepository
}

func NewChatroomUseCase(repo domain.ChatroomRepository) *ChatroomUseCase {
	return &ChatroomUseCase{
		chatroomRepository: repo,
	}
}

func (c *ChatroomUseCase) Create(chatroom *domain.Chatroom) error {
	if err := c.chatroomRepository.Insert(chatroom); err != nil {
		return err
	}
	return nil
}

func (c *ChatroomUseCase) FindById(id int64) (*domain.Chatroom, error) {
	chatroom, err := c.chatroomRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return chatroom, nil
}

func (c *ChatroomUseCase) FindAll() ([]*domain.Chatroom, error) {
	chatrooms, err := c.chatroomRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return chatrooms, nil
}

func (c *ChatroomUseCase) Update(id int64, chatroom *domain.Chatroom) error {
	if err := c.chatroomRepository.Update(id, chatroom); err != nil {
		return err
	}
	return nil
}

func (c *ChatroomUseCase) Delete(id int64) error {
	if err := c.chatroomRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
