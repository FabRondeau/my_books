package request

/* Switch between error and kind + totalItems */
type GoogleBookRequest struct {
	Kind       string           `validate:"required" json:"kind"`
	TotalItems int              `validate:"required" json:"totalItems"`
	Items      []GoogleBookItem `json:"items"`
}

type GoogleBookErrorRequest struct {
	Error GoogleBookError `validate:"required" json:"error"`
}

type GoogleBookError struct {
	Code    int    `validate:"required" json:"code"`
	Message string `validate:"required" json:"message"`
}
type GoogleBookItem struct {
	Kind       string                   `validate:"required" json:"kind"`
	GoogleID   string                   `validate:"required" json:"id"`
	GoogleEtag string                   `validate:"required" json:"etag"`
	SelfLink   string                   `validate:"required" json:"selfLink"`
	VolumeInfo GoogleBookItemVolumeInfo `validate:"required" json:"volumeInfo"`
}

type GoogleBookItemVolumeInfo struct {
	Title               string                         `validate:"required" json:"title"`
	SubTitle            string                         `validate:"required" json:"subtitle"`
	Authors             []string                       `validate:"required" json:"authors"`
	PublishedDate       string                         `validate:"required" json:"publishedDate"`
	IndustryIdentifiers []GoogleBookIndustryIdentifier `validate:"required" json:"industryIdentifiers"`
	ReadingModes        GoogleBookReadingMode          `validate:"required" json:"readingModes"`
	PageCount           int                            `validate:"required" json:"pageCount"`
	PrintType           string                         `validate:"required" json:"printType"`
	MaturityRating      string                         `validate:"required" json:"maturityRating"`
	AllowAnonLogging    bool                           `validate:"required" json:"allowAnonLogging"`
	ContentVersion      string                         `validate:"required" json:"contentVersion"`
	Language            string                         `validate:"required" json:"language"`
	PreviewLink         string                         `validate:"required" json:"previewLink"`
	InfoLink            string                         `validate:"required" json:"infoLink"`
	CanonicalVolumeLink string                         `validate:"required" json:"canonicalVolumeLink"`
}
type GoogleBookItemSaleInfo struct {
	Country     string `validate:"required" json:"country"`
	Saleability string `validate:"required" json:"saleability"`
	IsEbook     bool   `validate:"required" json:"isEbook"`
}
type GoogleBookItemAccessInfo struct {
	Country             string         `validate:"required" json:"country"`
	Viewability         string         `validate:"required" json:"viewability"`
	Embeddable          bool           `validate:"required" json:"embeddable"`
	PublicDomain        bool           `validate:"required" json:"publicDomain"`
	Epub                GoogleBookEPub `validate:"required" json:"epub"`
	Pdf                 GoogleBookPDF  `validate:"required" json:"pdf"`
	AccessViewStatus    string         `validate:"required" json:"accessViewStatus"`
	WebReaderLink       string         `validate:"required" json:"webReaderLink"`
	QuoteSharingAllowed bool           `validate:"required" json:"quoteSharingAllowed"`
}

// type GoogleBookAuthor struct {
// 	Name string
// }

type GoogleBookIndustryIdentifier struct {
	Type       string `validate:"required" json:"type"`
	Identifier string `validate:"required" json:"identifier"`
}

type GoogleBookReadingMode struct {
	Text  bool `validate:"required" json:"text"`
	Image bool `validate:"required" json:"image"`
}

type GoogleBookEPub struct {
	IsAvailable bool `validate:"required" json:"isAvailable"`
}

type GoogleBookPDF struct {
	IsAvailable bool `validate:"required" json:"isAvailable"`
}
