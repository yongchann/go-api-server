package member

import (
	"github.com/pkg/errors"
	"go-api-server/internal/entity"
	"go-api-server/internal/util"
	"strings"
)

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

func (u *UseCase) Login(nickname string, password string) (string, error) {
	findMember, err := u.repo.FindByNickname(nickname)
	if err != nil {
		return "", err
	}
	if strings.Compare(findMember.Password, password) != 0 {
		return "", errors.New("wrong password")
	}

	return util.CretaeJwt(findMember), nil
}
