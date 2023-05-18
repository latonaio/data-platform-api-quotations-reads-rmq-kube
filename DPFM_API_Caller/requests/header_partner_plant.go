package requests

type HeaderPartnerPlant struct {
	Quotation       int     `json:"Quotation"`
	PartnerFunction string  `json:"PartnerFunction"`
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}
