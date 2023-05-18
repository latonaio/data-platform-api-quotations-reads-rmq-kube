package requests

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
