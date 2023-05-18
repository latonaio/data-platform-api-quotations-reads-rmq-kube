package requests

type ItemPartner struct {
	Quotation                      int     `json:"Quotation"`
	QuotationItem                  int     `json:"QuotationItem"`
	PartnerFunction                string  `json:"PartnerFunction"`
	BusinessPartner                int     `json:"BusinessPartner"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	BusinessPartnerFullName        *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName            *string `json:"BusinessPartnerName"`
	AddressID                      *int    `json:"AddressID"`
}
