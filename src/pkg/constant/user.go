package constant

const (
	AlreadyExistsEmail    = "Email already exists, please try another one."
	AlreadyExistsPhone    = "Phone already exists, please try another one."
	AlreadyExistsUsername = "Username already exists, please try another one."

	LogoutFailed         = "Logout failed. Please try again later. Error: %s"
	RegisterFailed       = "Register failed. Please try again later. Error: %s"
	FailedUserNameOrPass = "Username or password is wrong, please try again."
	FailedLogin          = "Login failed, please try again. error: %s"
	FailedChangePwd      = "Password change failed, please try again. error: %s"
	FailedSavePicture    = "Profile picture cannot be saved, please try again. error: %s"
	FailedUserNotFound   = "User not found."
	FailedUNameMismatch  = "Username mismatch."
)

const (
	SuccessChangePass    = "Password changed successfully."
	SuccessUpdatePicture = "Profile picture updated successfully."
)
