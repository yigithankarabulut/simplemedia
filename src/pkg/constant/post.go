package constant

const (
	SuccessCreatePost = "Post created successfully."
	SuccessDeletePost = "Post deleted successfully."
	SuccessUpdatePost = "Post updated successfully."
)

const (
	FailedPostCreate         = "Post create failed. Please try again later. Error: %s"
	FailedPostDelete         = "Post delete failed. Please try again later. Error: %s"
	FailedDeleteUnauthorized = "Post delete failed. You are not authorized to delete this post."
	FailedGetPost            = "You are not authorized to get this post."
	FailedFriendPost         = "You are not authorized to get this post."
	FailedPostGet            = "Post get failed. Please try again later. Error: %s"
	FailedPostUpdate         = "Post update failed. Please try again later. Error: %s"
	PostNotFound             = "Post not found."
	PostAlreadyLiked         = "Post already liked."
)
