package usecases

import (
	"context"
	"errors"
	domain "task_manager/domain"
)

type UserUseCase struct {
	Repo    domain.UserRepository
	PassSvc domain.PasswordSvc
	JwtSvc  domain.JwtSvc
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserUseCase(repo domain.UserRepository, ps domain.PasswordSvc, jw domain.JwtSvc) *UserUseCase {
	return &UserUseCase{
		Repo:    repo,
		PassSvc: ps,
		JwtSvc:  jw,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	user.Password = uc.PassSvc.HashPassword(*user.Password)
	return uc.Repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, user domain.User) (string, error) {
	foundUser, err := uc.Repo.GetUser(ctx, user.Username)
	if err != nil {
		return "", err
	}
	if foundUser.Username == "" || foundUser.UserType == "" {
		return "", errors.New("user data incomplete")
	}
	valid, err := uc.PassSvc.VerifyPassword(*user.Password, *foundUser.Password)
	if err != nil || !valid {
		return "", errors.New("invalid password")
	}

	token, err := uc.JwtSvc.GenerateToken(foundUser.Username, foundUser.UserType)
	if err != nil {
		return "", err
	}
	return token, nil
}
