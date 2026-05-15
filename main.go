package main

import (
	"context"
	"fmt"

	user_repo "practices.com/clean_arch_go/infrastructure/repository/user"
	"practices.com/clean_arch_go/usecase/user/create_user"
)

func main() {
	// 创建 Repository 实例
	repo := user_repo.MockUserRepository{}
	// 创建 PasswordHasher 实例
	passwordHasher := MockPasswordHasher{}

	caseExecutor := create_user.NewCreateUserUseCase(repo, passwordHasher)
	in := create_user.Input{
		Name:     "Copyandpaste",
		Email:    "cap-0.example.com",
		Password: "example",
	}
	err := caseExecutor.Invoke(context.Background(), in)
	if err != nil {
		fmt.Printf("执行创建用户用例是发生系统错误, err: %s\n", err.Error())
		return
	}

	fmt.Println("用户创建成功")
}

type MockPasswordHasher struct{}

func (MockPasswordHasher) Hash(string) (string, error) {
	return "", nil
}

func (MockPasswordHasher) Compare(plainText, hashedText string) bool {
	return false
}
