package requests

type HeaderPartnerContact struct {
	Quotation         int     `json:"Quotation"`
	PartnerFunction   string  `json:"PartnerFunction"`
	ContactID         int     `json:"ContactID"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	ContactPersonName *string `json:"ContactPersonName"`
	EmailAddress      *string `json:"EmailAddress"`
	PhoneNumber       *string `json:"PhoneNumber"`
	MobilePhoneNumber *string `json:"MobilePhoneNumber"`
	FaxNumber         *string `json:"FaxNumber"`
	ContactTag1       *string `json:"ContactTag1"`
	ContactTag2       *string `json:"ContactTag2"`
	ContactTag3       *string `json:"ContactTag3"`
	ContactTag4       *string `json:"ContactTag4"`
}
