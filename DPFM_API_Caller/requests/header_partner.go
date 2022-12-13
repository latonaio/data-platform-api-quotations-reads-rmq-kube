package requests

type HeaderPartner struct {
	Quotation               int     `json:"Quotation"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Language                *string `json:"Language"`
	Organization            *string `json:"Organization"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}
