package request

type PublisherNameRequest struct {
	Name string `validate:"required" json:"name"`
}
