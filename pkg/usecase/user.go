package usecase

import (
	"context"
	"errors"
	"github/jonathanmoreiraa/planejja/pkg/api/middleware"
	"github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (useCase *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := useCase.userRepo.FindByID(ctx, id)
	return user, err
}

func (useCase *userUseCase) FindByEmail(ctx context.Context, email string) (domain.Users, error) {
	user, err := useCase.userRepo.FindByEmail(ctx, email)
	return user, err
}

func (useCase *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	passwordHashed, err := HashPassword(user.Password)
	if err != nil {
		return domain.Users{}, errors.New("erro ao processar a senha")
	}
	user.Password = passwordHashed

	user, err = useCase.userRepo.Save(ctx, user)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				err = errors.New("o e-mail informado já foi cadastrado anteriormente")
			}
		}
		return domain.Users{}, err
	}

	return user, nil
}

func (useCase *userUseCase) Update(ctx context.Context, user domain.Users) error {
	_, err := useCase.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}

	err = useCase.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *userUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := useCase.userRepo.Delete(ctx, user)

	return err
}

func (useCase *userUseCase) Login(ctx context.Context, email string, password string) (map[string]any, error) {
	user, err := useCase.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !middleware.CheckPasswordHash(user.Password, password) {
		return nil, errors.New("senha incorreta")
	}

	token, err := middleware.CreateToken(email)
	if err != nil {
		return nil, errors.New("erro ao gerar o token")
	}

	return map[string]any{
		"token":  token,
		"userId": user.ID,
	}, nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
