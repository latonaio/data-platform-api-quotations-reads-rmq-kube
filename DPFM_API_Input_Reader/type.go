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
	Header            Header   `json:"Orders"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	Quotation                      int           `json:"Quotation"`
	QuotationDate                  *string       `json:"QuotationDate"`
	QuoationType                   *string       `json:"QuoationType"`
	Buyer                          *int          `json:"Buyer"`
	Seller                         *int          `json:"Seller"`
	CreationDate                   *string       `json:"CreationDate"`
	LastChangeDate                 *string       `json:"LastChangeDate"`
	ContractType                   *string       `json:"ContractType"`
	VaridityStartDate              *string       `json:"VaridityStartDate"`
	ValidityEndDate                *string       `json:"ValidityEndDate"`
	InvoiceScheduleStartDate       *string       `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string       `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                 *float32      `json:"TotalNetAmount"`
	TotalTaxAmount                 *float32      `json:"TotalTaxAmount"`
	TotalGrossAmount               *float32      `json:"TotalGrossAmount"`
	TransactionCurrency            *string       `json:"TransactionCurrency"`
	PricingDate                    *string       `json:"PricingDate"`
	RequestedDeliveryDate          *string       `json:"RequestedDeliveryDate"`
	BindingPeriodValidityStartDate *string       `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   *string       `json:"BindingPeriodValidityEndDate"`
	Incoterms                      *string       `json:"Incoterms"`
	PaymentTerms                   *string       `json:"PaymentTerms"`
	PaymentMethod                  *string       `json:"PaymentMethod"`
	AccountingExchangeRate         *float32      `json:"AccountingExchangeRate"`
	BillingDocumentDate            *string       `json:"BillingDocumentDate"`
	HeaderText                     *string       `json:"HeaderText"`
	ReferenceInquiry               *int          `json:"ReferenceInquiry"`
	HeaderPartner                  HeaderPartner `json:"HeaderPartner"`
	Item                           Item          `json:"Item"`
	Address                        Address       `json:"Address"`
}

type HeaderPartner struct {
	Quotation               int                  `json:"Quotation"`
	PartnerFunction         string               `json:"PartnerFunction"`
	BusinessPartner         int                  `json:"BusinessPartner"`
	BusinessPartnerFullName *string              `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string              `json:"BusinessPartnerName"`
	Language                *string              `json:"Language"`
	Organization            *string              `json:"Organization"`
	Currency                *string              `json:"Currency"`
	ExternalDocumentID      *string              `json:"ExternalDocumentID"`
	AddressID               *int                 `json:"AddressID"`
	HeaderPartnerContact    HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant      HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
}

type HeaderPartnerContact struct {
	Quotation         int     `json:"Quotation"`
	PartnerFunction   string  `json:"PartnerFunction"`
	BusinessPartner   int     `json:"BusinessPartner"`
	ContactID         int     `json:"ContactID"`
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
	Quotation       int    `json:"Quotation"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type Item struct {
	Quotation                      int                `json:"Quotation"`
	QuotationItem                  int                `json:"QuotationItem"`
	QuotationItemCategory          *string            `json:"QuotationItemCategory"`
	QuotationItemText              *string            `json:"QuotationItemText"`
	Product                        *string            `json:"Product"`
	ProductStandardID              *string            `json:"ProductStandardID"`
	PricingDate                    *string            `json:"PricingDate"`
	QuotationQuantity              *float32           `json:"QuotationQuantity"`
	QuotationQuantityUnit          *string            `json:"QuotationQuantityUnit"`
	ItemGrossWeight                *float32           `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32           `json:"ItemNetWeight"`
	ItemWeightUnit                 *string            `json:"ItemWeightUnit"`
	NetAmount                      *float32           `json:"NetAmount"`
	TaxAmount                      *float32           `json:"TaxAmount"`
	GrossAmount                    *float32           `json:"GrossAmount"`
	ProductGroup                   *string            `json:"ProductGroup"`
	BillingDocumentDate            *string            `json:"BillingDocumentDate"`
	ProductionPlantBusinessPartner *string            `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string            `json:"ProductionPlant"`
	ProductionPlantTimeZone        *string            `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation *string            `json:"ProductionPlantStorageLocation"`
	IssuingPlantBusinessPartner    *string            `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string            `json:"IssuingPlant"`
	IssuingPlantTimeZone           *string            `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation    *string            `json:"IssuingPlantStorageLocation"`
	ReceivingPlantBusinessPartner  *string            `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string            `json:"ReceivingPlant"`
	ReceivingPlantTimeZone         *string            `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation  *string            `json:"ReceivingPlantStorageLocation"`
	Incoterms                      *string            `json:"Incoterms"`
	BPTaxClassification            *string            `json:"BPTaxClassification"`
	ProductTaxClassification       *string            `json:"ProductTaxClassification"`
	ProductAccountAssignmentGroup  *string            `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                   *string            `json:"PaymentTerms"`
	PaymentMethod                  *string            `json:"PaymentMethod"`
	Project                        *string            `json:"Project"`
	AccountingExchangeRate         *float32           `json:"AccountingExchangeRate"`
	ReferenceInquiry               *int               `json:"ReferenceInquiry"`
	ReferenceDocumentItem          *int               `json:"ReferenceDocumentItem"`
	ItemDeliveryStatus             *string            `json:"ItemDeliveryStatus"`
	IssuingStatus                  *string            `json:"IssuingStatus"`
	ReceivingStatus                *string            `json:"ReceivingStatus"`
	BillingStatus                  *string            `json:"BillingStatus"`
	TaxCode                        *string            `json:"TaxCode"`
	TaxRate                        *float32           `json:"TaxRate"`
	CountryOfOrigin                *string            `json:"CountryOfOrigin"`
	ItemPartner                    ItemPartner        `json:"ItemPartner"`
	ItemPricingElement             ItemPricingElement `json:"ItemPricingElement"`
}

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

type ItemPricingElement struct {
	Quotation                  int      `json:"Quotation"`
	QuotationItem              int      `json:"QuotationItem"`
	PricingProcedureStep       int      `json:"PricingProcedureStep"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
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
