package constant

const (
	Validate      = "Request could not be understood. Please make sure you send a valid request. Error: %s"
	Authenticate  = "Authentication failed. Please make sure you send a valid request. Error: %s"
	UnknowError   = "An unknown error occurred. Please try again later"
	Unauthorized  = "Unauthorized"
	InvalidParam  = "Invalid parameter"
	IDNotFound    = "ID not found. Error: %s"
	ErrorOccurred = "An error occurred. Error: %s"

	StandartGet          = "There was a problem while getting data. Error: %s"
	StandartCreate       = "There was a problem while creating data. Error: %s"
	StandartDelete       = "There was a problem while deleting data. Error: %s"
	StandartUpdate       = "There was a problem while updating data. Error: %s"
	StandartProcessError = "There was a problem while processing. Error: %s"
)
