package request

type AuthorFullNameRequest struct {
	FullName string `validate:"required" json:"fullname"`
}
