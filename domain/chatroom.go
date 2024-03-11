package domain

type Chatroom struct {
	Title     string
	CreatedAt string
	UpdatedAt string
}

type ChatroomRepository interface {
	Create(chatroom Chatroom) (*Chatroom, error)
	GetById(id int64) (*Chatroom, error)
	Fetch(offset, limit int) ([]*Chatroom, error)
	Update(id int64, chatroom Chatroom) (*Chatroom, error)
	Delete(id int64) error
}

type ChatroomUseCase interface {
	Create(chatroom Chatroom) (*Chatroom, error)
	GetById(id int64) (*Chatroom, error)
	Fetch(offset, limit int) ([]*Chatroom, error)
	Update(id int64, chatroom Chatroom) (*Chatroom, error)
	Delete(id int64) error
}
