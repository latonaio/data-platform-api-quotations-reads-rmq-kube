package requests

type ItemPartner struct {
	Quotation               int     `json:"Quotation"`
	QuotationItem           int     `json:"QuotationItem"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *string `json:"AddressID"`
}
