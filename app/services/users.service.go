package services

import (
	"errors"
	"strconv"

	"github.com/muhammadfikri4/go-plate/app/dto"
	"github.com/muhammadfikri4/go-plate/app/mappers"
	"github.com/muhammadfikri4/go-plate/app/repositories"
	"github.com/muhammadfikri4/go-plate/utils"
)

type UserService interface {
	CreateUser(dto *dto.CreateUserDTO) error
	GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error)
	GetUserById(userId string) (*dto.UserDTO, error)
	UpdateUser(userId string, dto *dto.UpdateUserDTO) (*dto.UserDTO, error)
	DeleteUser(userId string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) CreateUser(data *dto.CreateUserDTO) error {
	user := mappers.CreateUserDTOMapper(data)

	err := service.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error) {
	page, PerPage := utils.GetPaginationParams(query.Page, query.PerPage)

	users, totalItems, err := service.userRepository.GetUsers(page, PerPage)
	if err != nil {
		return nil, utils.Meta{}, err
	}

	usersDTOs := mappers.UsersDTOMapper(users)

	meta := utils.MetaPagination(
		page, PerPage, len(users), int(totalItems),
	)

	return usersDTOs, meta, err
}

func (service *userService) GetUserById(id string) (*dto.UserDTO, error) {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, nil
	}
	user, err := service.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserDTOMapper(user)

	return userDTO, err
}

func (service *userService) UpdateUser(id string, dto *dto.UpdateUserDTO) (*dto.UserDTO, error) {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, nil
	}

	user, err := service.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	userModel := mappers.UpdateUserDTOMapper(dto)

	err = service.userRepository.UpdateUser(uint(userId), userModel)
	if err != nil {
		return nil, err
	}

	updateduser, err := service.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserDTOMapper(updateduser)

	return userDTO, nil
}

func (service *userService) DeleteUser(id string) error {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}

	user, err := service.userRepository.GetUser(uint(userId))
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	err = service.userRepository.DeleteUser(uint(userId))
	if err != nil {
		return err
	}

	return nil
}
