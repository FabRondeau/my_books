package request

type BookTitleRequest struct {
	Title string `validate:"required" json:"title"`
}
type BookISBN13Request struct {
	ISBN13 string `validate:"required" json:"isbn13"`
}
