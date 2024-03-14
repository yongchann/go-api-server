package entity

type Chatroom struct {
	Id    int64 `gorm:"primaryKey, autoIncrement"`
	Title string
}

func (Chatroom) TableName() string {
	return "chatroom"
}

type ChatroomRepository interface {
	Insert(chatroom *Chatroom) error
	FindById(id int64) (*Chatroom, error)
	FindAll() ([]*Chatroom, error)
	Update(id int64, chatroom *Chatroom) error
	Delete(id int64) error
}

type ChatroomUseCase interface {
	Create(chatroom *Chatroom) error
	FindById(id int64) (*Chatroom, error)
	FindAll() ([]*Chatroom, error)
	Update(id int64, chatroom *Chatroom) error
	Delete(id int64) error
}
