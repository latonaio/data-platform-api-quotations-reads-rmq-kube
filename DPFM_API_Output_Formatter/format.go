package dpfm_api_output_formatter

import (
	"data-platform-api-quotations-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*Header, error) {
	pm := &requests.Header{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationDate,
			&pm.QuoationType,
			&pm.Buyer,
			&pm.Seller,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.VaridityStartDate,
			&pm.ValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.RequestedDeliveryDate,
			&pm.BindingPeriodValidityStartDate,
			&pm.BindingPeriodValidityEndDate,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.AccountingExchangeRate,
			&pm.BillingDocumentDate,
			&pm.HeaderText,
			&pm.ReferenceInquiry,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
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
	return header, nil
}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartner, error) {
	pm := &requests.HeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartner := &HeaderPartner{
		Quotation:               data.Quotation,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
	return headerPartner, nil
}

func ConvertToHeaderPartnerContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartnerContact, error) {
	pm := &requests.HeaderPartnerContact{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.ContactID,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartnerContact := &HeaderPartnerContact{
		Quotation:         data.Quotation,
		PartnerFunction:   data.PartnerFunction,
		BusinessPartner:   data.BusinessPartner,
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
	return headerPartnerContact, nil
}

func ConvertToHeaderPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartnerPlant, error) {
	pm := &requests.HeaderPartnerPlant{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartnerPlant := &HeaderPartnerPlant{
		Quotation:       data.Quotation,
		PartnerFunction: data.PartnerFunction,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
	return headerPartnerPlant, nil
}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*Address, error) {
	pm := &requests.Address{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	address := &Address{
		Quotation:   data.Quotation,
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
	return address, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*Item, error) {
	pm := &requests.Item{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.QuotationItemCategory,
			&pm.QuotationItemText,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.PricingDate,
			&pm.QuotationQuantity,
			&pm.QuotationQuantityUnit,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.ProductGroup,
			&pm.BillingDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantTimeZone,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantTimeZone,
			&pm.ReceivingPlantStorageLocation,
			&pm.Incoterms,
			&pm.BPTaxClassification,
			&pm.ProductTaxClassification,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.AccountingExchangeRate,
			&pm.ReferenceInquiry,
			&pm.ReferenceDocumentItem,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.BillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	item := &Item{
		Quotation:                      data.Quotation,
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
	return item, nil
}

func ConvertToItemPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPartner, error) {
	pm := &requests.ItemPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPartner := &ItemPartner{
		Quotation:               data.Quotation,
		QuotationItem:           data.QuotationItem,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
	return itemPartner, nil
}

func ConvertToItemPricingElement(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPricingElement, error) {
	pm := &requests.ItemPricingElement{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PricingProcedureStep,
			&pm.PricingProcedureStep,
			&pm.PricingProcedureCounter,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPricingElement := &ItemPricingElement{
		Quotation:                  data.Quotation,
		QuotationItem:              data.QuotationItem,
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
	return itemPricingElement, nil
}
