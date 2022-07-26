package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrRequestNotValid = errors.New("request not valid")

	ErrHistoryProductionResource = errors.New("history production not found or empty, you must create history production")

	ErrDuplicateData = errors.New("duplicate data")

	ErrEmailPasswordNotFound = errors.New("(Email) or (Password) not found")
)
