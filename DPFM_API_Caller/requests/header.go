package requests

type Header struct {
	Quotation                      int      `json:"Quotation"`
	QuotationDate                  *string  `json:"QuotationDate"`
	DistributionChannel            *string  `json:"DistributionChannel"`
	BusinessArea                   *string  `json:"BusinessArea"`
	District                       *string  `json:"District"`
	CreationDate                   *string  `json:"CreationDate"`
	LastChangeDate                 *string  `json:"LastChangeDate"`
	ContractType                   *string  `json:"ContractType"`
	VaridityStartDate              *string  `json:"VaridityStartDate"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	InvoiceScheduleStartDate       *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string  `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                 *float32 `json:"TotalNetAmount"`
	TransactionCurrency            *string  `json:"TransactionCurrency"`
	PricingDate                    *string  `json:"PricingDate"`
	RequestedDeliveryDate          *string  `json:"RequestedDeliveryDate"`
	BindingPeriodValidityStartDate *string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   *string  `json:"BindingPeriodValidityEndDate"`
	OrderProbabilityInPercent      *float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount         *float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                      *string  `json:"Incoterms"`
	PaymentTerms                   *string  `json:"PaymentTerms"`
	PaymentMethod                  *string  `json:"PaymentMethod"`
	TaxClassification              *string  `json:"TaxClassification"`
	ReferenceInquiry               *int     `json:"ReferenceInquiry"`
}
