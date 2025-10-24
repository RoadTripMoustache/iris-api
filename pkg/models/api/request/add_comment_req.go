package request

type AddCommentReq struct {
	Message string   `json:"message"`
	Images  []string `json:"images"`
}
