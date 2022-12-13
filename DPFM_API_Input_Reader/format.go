package dpfm_api_input_reader

import (
	"data-platform-api-quotations-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		Quotation:                      data.Quotation,
		QuotationDate:                  data.QuotationDate,
		QuoationType:                   data.QuoationType,
		Buyer:                          data.Buyer,
		Seller:                         data.Seller,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		ContractType:                   data.ContractType,
		VaridityStartDate:              data.VaridityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		InvoiceScheduleStartDate:       data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:         data.InvoiceScheduleEndDate,
		TotalNetAmount:                 data.TotalNetAmount,
		TotalTaxAmount:                 data.TotalTaxAmount,
		TotalGrossAmount:               data.TotalGrossAmount,
		TransactionCurrency:            data.TransactionCurrency,
		PricingDate:                    data.PricingDate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
		Incoterms:                      data.Incoterms,
		PaymentTerms:                   data.PaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		AccountingExchangeRate:         data.AccountingExchangeRate,
		BillingDocumentDate:            data.BillingDocumentDate,
		HeaderText:                     data.HeaderText,
		ReferenceInquiry:               data.ReferenceInquiry,
	}
}

func (sdc *SDC) ConvertToHeaderPartner() *requests.HeaderPartner {
	dataHeader := sdc.Header
	data := sdc.Header.HeaderPartner
	return &requests.HeaderPartner{
		Quotation:               dataHeader.Quotation,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Language:                data.Language,
		Organization:            data.Organization,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerPlant() *requests.HeaderPartnerPlant {
	dataHeader := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner
	data := sdc.Header.HeaderPartner.HeaderPartnerPlant
	return &requests.HeaderPartnerPlant{
		Quotation:       dataHeader.Quotation,
		PartnerFunction: dataHeaderPartner.PartnerFunction,
		BusinessPartner: dataHeaderPartner.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerContact() *requests.HeaderPartnerContact {
	dataHeader := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner
	data := sdc.Header.HeaderPartner.HeaderPartnerContact
	return &requests.HeaderPartnerContact{
		Quotation:         dataHeader.Quotation,
		PartnerFunction:   dataHeaderPartner.PartnerFunction,
		BusinessPartner:   dataHeaderPartner.BusinessPartner,
		ContactID:         data.ContactID,
		ContactPersonName: data.ContactPersonName,
		EmailAddress:      data.EmailAddress,
		PhoneNumber:       data.PhoneNumber,
		MobilePhoneNumber: data.MobilePhoneNumber,
		FaxNumber:         data.FaxNumber,
		ContactTag1:       data.ContactTag1,
		ContactTag2:       data.ContactTag2,
		ContactTag3:       data.ContactTag3,
		ContactTag4:       data.ContactTag4,
	}
}

func (sdc *SDC) ConvertToAddress() *requests.Address {
	dataHeader := sdc.Header
	data := sdc.Header.Address
	return &requests.Address{
		Quotation:   dataHeader.Quotation,
		AddressID:   data.AddressID,
		PostalCode:  data.PostalCode,
		LocalRegion: data.LocalRegion,
		Country:     data.Country,
		District:    data.District,
		StreetName:  data.StreetName,
		CityName:    data.CityName,
		Building:    data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataHeader := sdc.Header
	data := sdc.Header.Item
	return &requests.Item{
		Quotation:                      dataHeader.Quotation,
		QuotationItem:                  data.QuotationItem,
		QuotationItemCategory:          data.QuotationItemCategory,
		QuotationItemText:              data.QuotationItemText,
		Product:                        data.Product,
		ProductStandardID:              data.ProductStandardID,
		PricingDate:                    data.PricingDate,
		QuotationQuantity:              data.QuotationQuantity,
		QuotationQuantityUnit:          data.QuotationQuantityUnit,
		ItemGrossWeight:                data.ItemGrossWeight,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		NetAmount:                      data.NetAmount,
		TaxAmount:                      data.TaxAmount,
		GrossAmount:                    data.GrossAmount,
		ProductGroup:                   data.ProductGroup,
		BillingDocumentDate:            data.BillingDocumentDate,
		ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
		ProductionPlant:                data.ProductionPlant,
		ProductionPlantTimeZone:        data.ProductionPlantTimeZone,
		ProductionPlantStorageLocation: data.ProductionPlantStorageLocation,
		IssuingPlantBusinessPartner:    data.IssuingPlantBusinessPartner,
		IssuingPlant:                   data.IssuingPlant,
		IssuingPlantTimeZone:           data.IssuingPlantTimeZone,
		IssuingPlantStorageLocation:    data.IssuingPlantStorageLocation,
		ReceivingPlantBusinessPartner:  data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                 data.ReceivingPlant,
		ReceivingPlantTimeZone:         data.ReceivingPlantTimeZone,
		ReceivingPlantStorageLocation:  data.ReceivingPlantStorageLocation,
		Incoterms:                      data.Incoterms,
		BPTaxClassification:            data.BPTaxClassification,
		ProductTaxClassification:       data.ProductTaxClassification,
		ProductAccountAssignmentGroup:  data.ProductAccountAssignmentGroup,
		PaymentTerms:                   data.PaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		Project:                        data.Project,
		AccountingExchangeRate:         data.AccountingExchangeRate,
		ReferenceInquiry:               data.ReferenceInquiry,
		ReferenceDocumentItem:          data.ReferenceDocumentItem,
		ItemDeliveryStatus:             data.ItemDeliveryStatus,
		IssuingStatus:                  data.IssuingStatus,
		ReceivingStatus:                data.ReceivingStatus,
		BillingStatus:                  data.BillingStatus,
		TaxCode:                        data.TaxCode,
		TaxRate:                        data.TaxRate,
		CountryOfOrigin:                data.CountryOfOrigin,
	}
}

func (sdc *SDC) ConvertToItemPartner() *requests.ItemPartner {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	data := sdc.Header.Item.ItemPartner
	return &requests.ItemPartner{
		Quotation:               dataHeader.Quotation,
		QuotationItem:           dataItem.QuotationItem,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToItemPricingElement() *requests.ItemPricingElement {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	data := sdc.Header.Item.ItemPricingElement
	return &requests.ItemPricingElement{
		Quotation:                  dataHeader.Quotation,
		QuotationItem:              dataItem.QuotationItem,
		PricingProcedureStep:       data.PricingProcedureStep,
		PricingProcedureCounter:    data.PricingProcedureCounter,
		ConditionType:              data.ConditionType,
		PricingDate:                data.PricingDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionCurrency:          data.ConditionCurrency,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionQuantityUnit:      data.ConditionQuantityUnit,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		TaxCode:                    data.TaxCode,
		ConditionAmount:            data.ConditionAmount,
		TransactionCurrency:        data.TransactionCurrency,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}
}
