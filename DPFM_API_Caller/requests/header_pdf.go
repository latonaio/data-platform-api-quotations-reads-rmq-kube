package requests

type HeaderPDF struct {
	Quotation                *int    `json:"Quotation"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FileName                 *string `json:"FileName"`
}
