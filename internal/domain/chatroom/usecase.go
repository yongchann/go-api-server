package chatroom

import (
	"go-api-server/internal/entity"
)

type UseCase struct {
	repo entity.ChatroomRepository
}

func NewUseCase(repo entity.ChatroomRepository) entity.ChatroomUseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Create(chatroom *entity.Chatroom) error {
	if err := u.repo.Insert(chatroom); err != nil {
		return err
	}
	return nil
}

func (u *UseCase) FindById(id int64) (*entity.Chatroom, error) {
	chatroom, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return chatroom, nil
}

func (u *UseCase) FindAll() ([]*entity.Chatroom, error) {
	chatrooms, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return chatrooms, nil
}

func (u *UseCase) Update(id int64, chatroom *entity.Chatroom) error {
	if err := u.repo.Update(id, chatroom); err != nil {
		return err
	}
	return nil
}

func (u *UseCase) Delete(id int64) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
