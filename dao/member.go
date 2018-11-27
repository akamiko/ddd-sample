package dao

import (
	"fmt"

	"github.com/akamiko/ddd-sample/repository"
)

type Member struct{}

func ProvideMember() repository.Member {
	return Member{}
}

func (m Member) Foo() {
	fmt.Println("Foo")
}
