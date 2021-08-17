package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound               = errors.New("id_not_found")
	ErrRecordNotFound           = errors.New("record_not_found")
	ErrDuplicateData            = errors.New("duplicate_data")
	ErrDataAlreadyExist         = errors.New("data_already_exist")
	ErrInvalidBearerToken       = errors.New("invalid_bearer_token")
	ErrExpiredToken             = errors.New("expired_token")
	ErrInvalidRole              = errors.New("invalid_role")
	ErrInvalidCred              = errors.New("invalid_credential")
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")

	//Modular
	BaseResponseMessageSuccess  = "success"
	BaseResponseMessageFailed   = "something not right"
	BaseResponseMessageInserted = "data_inserted"
	BaseResponseMessageUpdated  = "data_updated"
)
