package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"Quotations"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	Quotation                      int             `json:"Quotation"`
	QuotationDate                  *string         `json:"QuotationDate"`
	DistributionChannel            *string         `json:"DistributionChannel"`
	BusinessArea                   *string         `json:"BusinessArea"`
	District                       *string         `json:"District"`
	CreationDate                   *string         `json:"CreationDate"`
	LastChangeDate                 *string         `json:"LastChangeDate"`
	ContractType                   *string         `json:"ContractType"`
	VaridityStartDate              *string         `json:"VaridityStartDate"`
	ValidityEndDate                *string         `json:"ValidityEndDate"`
	InvoiceScheduleStartDate       *string         `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string         `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                 *float32        `json:"TotalNetAmount"`
	TransactionCurrency            *string         `json:"TransactionCurrency"`
	PricingDate                    *string         `json:"PricingDate"`
	RequestedDeliveryDate          *string         `json:"RequestedDeliveryDate"`
	BindingPeriodValidityStartDate *string         `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   *string         `json:"BindingPeriodValidityEndDate"`
	OrderProbabilityInPercent      *float32        `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount         *float32        `json:"ExpectedOrderNetAmount"`
	Incoterms                      *string         `json:"Incoterms"`
	PaymentTerms                   *string         `json:"PaymentTerms"`
	PaymentMethod                  *string         `json:"PaymentMethod"`
	TaxClassification              *string         `json:"TaxClassification"`
	ReferenceInquiry               *int            `json:"ReferenceInquiry"`
	HeaderPartner                  []HeaderPartner `json:"HeaderPartner"`
	Item                           []Item          `json:"Item"`
	Address                        []Address       `json:"Address"`
}

type HeaderPartner struct {
	Quotation               int                    `json:"Quotation"`
	PartnerFunction         string                 `json:"PartnerFunction"`
	BusinessPartner         *int                   `json:"BusinessPartner"`
	BusinessPartnerFullName *string                `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string                `json:"BusinessPartnerName"`
	QuotationType           *string                `json:"QuotationType"`
	Organization            *string                `json:"Organization"`
	Language                *string                `json:"Language"`
	Currency                *string                `json:"Currency"`
	ExternalDocumentID      *string                `json:"ExternalDocumentID"`
	AddressID               *int                   `json:"AddressID"`
	HeaderPartnerContact    []HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant      []HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
}

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

type HeaderPartnerPlant struct {
	Quotation       int     `json:"Quotation"`
	PartnerFunction string  `json:"PartnerFunction"`
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}

type Item struct {
	Quotation                     int                  `json:"Quotation"`
	QuotationItem                 int                  `json:"QuotationItem"`
	QuotationItemCategory         *string              `json:"QuotationItemCategory"`
	QuotationItemText             *string              `json:"QuotationItemText"`
	Product                       *string              `json:"Product"`
	QuotationQuantity             *float32             `json:"QuotationQuantity"`
	QuotationQuantityUnit         *string              `json:"QuotationQuantityUnit"`
	ItemOrderProbabilityInPercent *float32             `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               *float32             `json:"ItemGrossWeight"`
	ItemNetWeight                 *float32             `json:"ItemNetWeight"`
	ItemWeightUnit                *string              `json:"ItemWeightUnit"`
	TransactionCurrency           *string              `json:"TransactionCurrency"`
	BussinessParnterCurrency      *string              `json:"BussinessParnterCurrency"`
	NetAmount                     *float32             `json:"NetAmount"`
	ProductGroup                  *string              `json:"ProductGroup"`
	ProductPricingGroup           *string              `json:"ProductPricingGroup"`
	Batch                         *string              `json:"Batch"`
	IssuingPlant                  *string              `json:"IssuingPlant"`
	IssuingPlantStorageLocation   *string              `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                *string              `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation *string              `json:"ReceivingPlantStorageLocation"`
	Incoterms                     *string              `json:"Incoterms"`
	PaymentTerms                  *string              `json:"PaymentTerms"`
	ProductTaxClassification1     *string              `json:"ProductTaxClassification1"`
	Project                       *string              `json:"Project"`
	ProfitCenter                  *string              `json:"ProfitCenter"`
	ReferenceInquiry              *int                 `json:"ReferenceInquiry"`
	ReferenceInquiryItem          *int                 `json:"ReferenceInquiryItem"`
	ItemPartner                   []ItemPartner        `json:"ItemPartner"`
	ItemPricingElement            []ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPartner struct {
	Quotation                      int     `json:"Quotation"`
	QuotationItem                  int     `json:"QuotationItem"`
	PartnerFunction                string  `json:"PartnerFunction"`
	BusinessPartner                *int    `json:"BusinessPartner"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	BusinessPartnerFullName        *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName            *string `json:"BusinessPartnerName"`
	AddressID                      *int    `json:"AddressID"`
}

type ItemPricingElement struct {
	BusinessPartner                int      `json:"BusinessPartner"`
	Quotation                      int      `json:"Quotation"`
	QuotationItem                  int      `json:"QuotationItem"`
	PricingProcedureStep           int      `json:"PricingProcedureStep"`
	PricingProcedureCounter        int      `json:"PricingProcedureCounter"`
	ConditionType                  *string  `json:"ConditionType"`
	PriceConditionDeterminationDte *string  `json:"PriceConditionDeterminationDte"`
	ConditionRateValue             *float32 `json:"ConditionRateValue"`
	ConditionCurrency              *string  `json:"ConditionCurrency"`
	ConditionRecord                *int     `json:"ConditionRecord"`
	ConditionSequentialNumber      *int     `json:"ConditionSequentialNumber"`
	TaxCode                        *string  `json:"TaxCode"`
	TransactionCurrency            *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged     *bool    `json:"ConditionIsManuallyChanged"`
}

type Address struct {
	Quotation   int     `json:"Quotation"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
