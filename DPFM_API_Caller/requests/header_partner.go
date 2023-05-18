package requests

type HeaderPartner struct {
	Quotation               int     `json:"Quotation"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	QuotationType           *string `json:"QuotationType"`
	Organization            *string `json:"Organization"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}
