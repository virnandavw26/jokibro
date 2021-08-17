package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound                 = errors.New("id_not_found")
	ErrRecordNotFound             = errors.New("record_not_found")
	ErrDuplicateData              = errors.New("duplicate_data")
	ErrDataAlreadyExist           = errors.New("data_already_exist")
	ErrInvalidBearerToken         = errors.New("invalid_bearer_token")
	ErrExpiredToken               = errors.New("expired_token")
	ErrInvalidRole                = errors.New("invalid_role")
	ErrInvalidCred                = errors.New("invalid_credential")
	ErrInternalServer             = errors.New("something gone wrong, contact administrator")
	ErrUsernamePasswordNotFound   = errors.New("(Username) or (Password) empty")
	ErrWorkerIntersectWorkingDate = errors.New("worker_intersect_working_sate")
	ErrInvalidStartWorkingAt      = errors.New("invalid_start_working_at")
	ErrInvalidEndWorkingAt        = errors.New("invalid_end_working_at")

	//Modular
	BaseResponseMessageSuccess  = "success"
	BaseResponseMessageFailed   = "something not right"
	BaseResponseMessageInserted = "data_inserted"
	BaseResponseMessageUpdated  = "data_updated"
)
