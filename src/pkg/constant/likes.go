package constant

const (
	PostOrCommentIDNotFound = "PostID or CommentID must be provided."
	AddLikeSuccess          = "Like successfully added."
	DeleteLikeSuccess       = "Like successfully deleted."

	FailedLikeAdd      = "Like add failed. Please try again later. Error: %s"
	FailedLikeNotFound = "Like not found."
	FailedNotBelong    = "You are not authorized to delete this like."
	FailedLikeDelete   = "Like delete failed. Please try again later. Error: %s"
	FailedLikeList     = "Like list failed. Please try again later. Error: %s"
)
