package converter

import (
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetID(entity.ID.Hex())

	return domain
}
