package users

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	"github.com/dedenfarhanhub/smart-koi-be/business"
	"github.com/dedenfarhanhub/smart-koi-be/helper/encrypt"
	"github.com/dedenfarhanhub/smart-koi-be/helper/logging"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"strings"
	"time"
)

type UserUseCase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
	logger         logging.Logger
}

func (u UserUseCase) Destroy(ctx context.Context, id int) (Domain, error) {
	existedUser, err := u.userRepository.GetByID(ctx, id)

	if existedUser.Role == "ADMIN" {
		return Domain{}, business.ErrRequestNotValid
	}

	result, err := u.userRepository.Destroy(ctx, id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return result, nil
}

func (u UserUseCase) Login(ctx context.Context, username, password string, sso bool) (string, Domain, error) {
	request := map[string]interface{}{
		"username": username,
		"sso":   sso,
	}

	existedUser, err := u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		u.logger.LogEntry(request, result).Error(err.Error())
		return "", Domain{}, err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", Domain{}, business.ErrEmailPasswordNotFound
	}

	token := u.jwtAuth.GenerateToken(existedUser.ID, existedUser.Role)
	result := map[string]interface{}{
		"success": "true",
	}
	u.logger.LogEntry(request, result).Info("incoming request")
	return token, existedUser, nil
}

func (u UserUseCase) Store(ctx context.Context, data *Domain, sso bool) (Domain, error) {

	request := map[string]interface{}{
		"username": data.Username,
		"name":  data.Name,
	}

	existedUser, err := u.userRepository.GetByUsername(ctx, data.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			u.logger.LogEntry(request, result).Error(err.Error())
			return Domain{}, err
		}
	}
	if existedUser != (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if !sso {
		data.Password, _ = encrypt.Hash(data.Password)
	}

	user, err := u.userRepository.Store(ctx, data)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		u.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	u.logger.LogEntry(request, result).Info("incoming request")

	return user, nil
}

func (u UserUseCase) Fetch(ctx context.Context, pagination pagination.Pagination, keyword string) (pagination.Pagination, []Domain, error) {
	result, allUsers, err := u.userRepository.Fetch(ctx, pagination, keyword)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		u.logger.LogEntry("can't get all data users", result).Error(err.Error())

		return pagination, []Domain{}, err
	}

	u.logger.LogEntry("success to get all data users", nil).Info("incoming request")

	return result, allUsers, nil
}

func (u UserUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return result, nil
}

func (u UserUseCase) Update(ctx context.Context, data *Domain, id int) (Domain, error) {
	request := map[string]interface{}{
		"id":       	id,
		"name":     	data.Name,
		"username":    	data.Username,
		"phone_number":	data.PhoneNumber,
		"password": 	data.Password,
		"role":  		data.Role,
	}

	exitedUser, err := u.userRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	if data.Password != "" {
		data.Password, _ = encrypt.Hash(data.Password)
	}
	if data.Password == "" {
		data.Password = exitedUser.Password
	}

	if data.Role == "" {
		data.Role = exitedUser.Role
	}

	if data.PhoneNumber == "" {
		data.PhoneNumber = exitedUser.PhoneNumber
	}

	data.Username = exitedUser.Username
	savedUser, err := u.userRepository.Update(ctx, data, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		u.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	u.logger.LogEntry(request, result).Info("incoming request")

	return savedUser, nil
}


func NewUserUseCase(ur Repository, jwtAuth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) UseCase {
	return &UserUseCase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
		contextTimeout: timeout,
		logger:         logger,
	}
}