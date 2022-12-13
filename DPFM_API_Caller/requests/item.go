package requests

type Item struct {
	Quotation                      int      `json:"Quotation"`
	QuotationItem                  int      `json:"QuotationItem"`
	QuotationItemCategory          *string  `json:"QuotationItemCategory"`
	QuotationItemText              *string  `json:"QuotationItemText"`
	Product                        *string  `json:"Product"`
	ProductStandardID              *string  `json:"ProductStandardID"`
	PricingDate                    *string  `json:"PricingDate"`
	QuotationQuantity              *float32 `json:"QuotationQuantity"`
	QuotationQuantityUnit          *string  `json:"QuotationQuantityUnit"`
	ItemGrossWeight                *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                 *string  `json:"ItemWeightUnit"`
	NetAmount                      *float32 `json:"NetAmount"`
	TaxAmount                      *float32 `json:"TaxAmount"`
	GrossAmount                    *float32 `json:"GrossAmount"`
	ProductGroup                   *string  `json:"ProductGroup"`
	BillingDocumentDate            *string  `json:"BillingDocumentDate"`
	ProductionPlantBusinessPartner *string  `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone        *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation *string  `json:"ProductionPlantStorageLocation"`
	IssuingPlantBusinessPartner    *string  `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string  `json:"IssuingPlant"`
	IssuingPlantTimeZone           *string  `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation    *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantBusinessPartner  *string  `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string  `json:"ReceivingPlant"`
	ReceivingPlantTimeZone         *string  `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation  *string  `json:"ReceivingPlantStorageLocation"`
	Incoterms                      *string  `json:"Incoterms"`
	BPTaxClassification            *string  `json:"BPTaxClassification"`
	ProductTaxClassification       *string  `json:"ProductTaxClassification"`
	ProductAccountAssignmentGroup  *string  `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                   *string  `json:"PaymentTerms"`
	PaymentMethod                  *string  `json:"PaymentMethod"`
	Project                        *string  `json:"Project"`
	AccountingExchangeRate         *float32 `json:"AccountingExchangeRate"`
	ReferenceInquiry               *int     `json:"ReferenceInquiry"`
	ReferenceDocumentItem          *int     `json:"ReferenceDocumentItem"`
	ItemDeliveryStatus             *string  `json:"ItemDeliveryStatus"`
	IssuingStatus                  *string  `json:"IssuingStatus"`
	ReceivingStatus                *string  `json:"ReceivingStatus"`
	BillingStatus                  *string  `json:"BillingStatus"`
	TaxCode                        *string  `json:"TaxCode"`
	TaxRate                        *float32 `json:"TaxRate"`
	CountryOfOrigin                *string  `json:"CountryOfOrigin"`
}
