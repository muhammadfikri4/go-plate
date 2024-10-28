package mappers

import (
	"github.com/muhammadfikri4/go-plate/app/dto"
	"github.com/muhammadfikri4/go-plate/app/models"
)

func UserDTOMapper(data *models.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Age:   data.Age,
	}
}

func UserModelMapper(data *dto.UserDTO) *models.User {
	return &models.User{
		Name:  data.Name,
		Email: data.Email,
		Age:   data.Age,
	}
}

func UpdateUserDTOMapper(data *dto.UpdateUserDTO) *models.User {
	return &models.User{
		Name:  data.Name,
		Email: data.Email,
		Age:   data.Age,
	}
}

func UsersDTOMapper(datas []*models.User) []*dto.UserDTO {
	dtos := make([]*dto.UserDTO, len(datas))
	for i, item := range datas {
		dtos[i] = &dto.UserDTO{
			ID:    item.ID,
			Name:  item.Name,
			Email: item.Email,
			Age:   item.Age,
		}
	}
	return dtos
}

func CreateUserDTOMapper(data *dto.CreateUserDTO) *models.User {
	return &models.User{
		Name:  data.Name,
		Email: data.Email,
		Age:   data.Age,
	}
}
