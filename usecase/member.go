package usecase

import (
	"fmt"

	"github.com/akamiko/ddd-sample/dao"
	"github.com/akamiko/ddd-sample/repository"
)

// Member
type Member struct {
	repo memberRepository
}

// MemberRepository ここで利用するリポジトリを定義
type memberRepository struct {
	member repository.Member
}

func ProvideMember() *Member {
	memberRepo := dao.ProvideMember()
	return &Member{
		repo: memberRepository{
			member: memberRepo,
		},
	}
}
func (m Member) Hoge() {

	fmt.Println("hogehoge!!")
	// dao -> リポジトリ
	m.repo.member.Foo()

	//m.repo.member.Foo()
}
