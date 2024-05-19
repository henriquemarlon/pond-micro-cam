package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/domain/dto"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository entity.UserRepository
}

func NewUserUseCase(userRepository entity.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(input *dto.CreateUserInputDTO) (*dto.CreateUserOutputDTO, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	//TODO: Implement update that does not require all fields of input DTO (Maybe i can do this only in the repository?)
	user := entity.NewUser(input.Name, input.Email, string(hashedPassword))
	res, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserOutputDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (u *UserUseCase) FindUserById(id string) (*dto.FindUserOutputDTO, error) {
	res, err := u.UserRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return &dto.FindUserOutputDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (u *UserUseCase) FindAllUsers() ([]*dto.FindUserOutputDTO, error) {
	res, err := u.UserRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	var output []*dto.FindUserOutputDTO
	for _, user := range res {
		output = append(output, &dto.FindUserOutputDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return output, nil
}

func (u *UserUseCase) LoginUser(input *dto.LoginInputDTO) (*dto.LoginOutputDTO, error) {
	res, err := u.UserRepository.FindUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password)); err != nil {
		return nil, err
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Id:        res.ID,
		Issuer:    "authorization-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, fmt.Errorf("error signing token: %v", err)
	}

	return &dto.LoginOutputDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		Token:     tokenString,
	}, nil
}