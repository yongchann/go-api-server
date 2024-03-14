package member

import "go-api-server/internal/entity"

type UseCase struct {
	repo entity.MemberRepository
}

func NewUseCase(repo entity.MemberRepository) entity.MemberUseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Join(member *entity.Member) error {
	if err := u.repo.Insert(member); err != nil {
		return err
	}
	return nil
}

func (u *UseCase) FindByNickname(nickname string) (*entity.Member, error) {
	member, err := u.repo.FindByNickname(nickname)
	if err != nil {
		return nil, err
	}
	return member, nil
}
