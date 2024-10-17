package constant

import "net/http"

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	AlreadyExists
	InvalidCredentials
	NotFound
)

const (
	TokenHourLifespan = 1
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "ALREADY_EXISTS", "INVALID_CREDENTIALS", "NOT_FOUND"}[r-1]
}

func (r ResponseStatus) GetResponseStatusCode() int {
	return [...]int{http.StatusOK, http.StatusNotFound, http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusConflict, http.StatusBadRequest, http.StatusNotFound}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "Already Exists", "Invalid credentials", "Not found"}[r-1]
}
