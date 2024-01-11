package constant

const (
	CreateCommentSuccess = "Comment created successfully."
	DeleteCommentSuccess = "Comment deleted successfully."
	UpdateCommentSuccess = "Comment updated successfully."

	CommentNotUseTogether = "CommentID and PostID cannot be used together."
	CommentNotFound       = "Comment not found."
	CommentAlreadyLiked   = "Comment already liked."
	CommentParentNotFound = "Parent comment not found."
	CommentNotOwner       = "You are not authorized to delete this comment."
	CommentCreateFail     = "Comment create failed. Please try again later. Error: %s"
	CommentDeleteFail     = "Comment delete failed. Please try again later. Error: %s"
	CommentUpdateFail     = "Comment update failed. Please try again later. Error: %s"
)
