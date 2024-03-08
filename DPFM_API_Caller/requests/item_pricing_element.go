package requests

type ItemPricingElement struct {
	Quotation					int		`json:"Quotation"`
	QuotationItem				int		`json:"QuotationItem"`
	PricingProcedureCounter		int		`json:"PricingProcedureCounter"`
	SupplyChainRelationshipID	int		`json:"SupplyChainRelationshipID"`
	Buyer						int		`json:"Buyer"`
	Seller						int		`json:"Seller"`
	ConditionRecord				int		`json:"ConditionRecord"`
	ConditionSequentialNumber	int		`json:"ConditionSequentialNumber"`
	ConditionType				string	`json:"ConditionType"`
	PricingDate					string	`json:"PricingDate"`
	ConditionRateValue			float32	`json:"ConditionRateValue"`
	ConditionRateValueUnit		int		`json:"ConditionRateValueUnit"`
	ConditionScaleQuantity		int		`json:"ConditionScaleQuantity"`
	ConditionCurrency			string	`json:"ConditionCurrency"`
	ConditionQuantity			float32	`json:"ConditionQuantity"`
	TaxCode						*string	`json:"TaxCode"`
	ConditionAmount				float32	`json:"ConditionAmount"`
	TransactionCurrency			string	`json:"TransactionCurrency"`
	ConditionIsManuallyChanged	*bool	`json:"ConditionIsManuallyChanged"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsCancelled					*bool	`json:"IsCancelled"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
