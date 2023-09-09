package dpfm_api_output_formatter

import (
	"data-platform-api-quotations-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationDate,
			&pm.QuotationType,
			&pm.QuotationStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.ContractType,
			&pm.BindingPeriodValidityStartDate,
			&pm.BindingPeriodValidityEndDate,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.HeaderOrderIsDefined,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.OrderProbabilityInPercent,
			&pm.ExpectedOrderNetAmount,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderIsClosed,
			&pm.HeaderBlockStatus,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			Quotation:							data.Quotation,
			QuotationDate:						data.QuotationDate,
			QuotationType:						data.QuotationType,
			QuotationStatus:					data.QuotationStatus,
			SupplyChainRelationshipID:			data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID:	data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:	data.SupplyChainRelationshipPaymentID,
			Buyer:								data.Buyer,
			Seller:								data.Seller,
			BillToParty:						data.BillToParty,
			BillFromParty:						data.BillFromParty,
			BillToCountry:						data.BillToCountry,
			BillFromCountry:					data.BillFromCountry,
			Payer:								data.Payer,
			Payee:								data.Payee,
			ContractType:						data.ContractType,
			BindingPeriodValidityStartDate:		data.BindingPeriodValidityStartDate,
			BindingPeriodValidityEndDate:		data.BindingPeriodValidityEndDate,
			OrderValidityStartDate:				data.OrderValidityStartDate,
			OrderValidityEndDate:				data.OrderValidityEndDate,
			InvoicePeriodStartDate:				data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:				data.InvoicePeriodEndDate,
			TotalNetAmount:						data.TotalNetAmount,
			TotalTaxAmount:						data.TotalTaxAmount,
			TotalGrossAmount:					data.TotalGrossAmount,
			HeaderOrderIsDefined:				data.HeaderOrderIsDefined,
			TransactionCurrency:				data.TransactionCurrency,
			PricingDate:						data.PricingDate,
			PriceDetnExchangeRate:				data.PriceDetnExchangeRate,
			RequestedDeliveryDate:				data.RequestedDeliveryDate,
			OrderProbabilityInPercent:			data.OrderProbabilityInPercent,
			ExpectedOrderNetAmount:				data.ExpectedOrderNetAmount,
			Incoterms:							data.Incoterms,
			PaymentTerms:						data.PaymentTerms,
			PaymentMethod:						data.PaymentMethod,
			ReferenceDocument:					data.ReferenceDocument,
			AccountAssignmentGroup:				data.AccountAssignmentGroup,
			AccountingExchangeRate:				data.AccountingExchangeRate,
			InvoiceDocumentDate:				data.InvoiceDocumentDate,
			IsExportImport:						data.IsExportImport,
			HeaderText:							data.HeaderText,
			HeaderIsClosed:						data.HeaderIsClosed,
			HeaderBlockStatus:					data.HeaderBlockStatus,
			CreationDate:						data.CreationDate,
			LastChangeDate:						data.LastChangeDate,
			IsCancelled:						data.IsCancelled,
			IsMarkedForDeletion:				data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil

}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.QuotationItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.QuotationItemText,
			&pm.QuotationItemTextByBuyer,
			&pm.QuotationItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.DeliveryUnit,
			&pm.ServicesRenderingDate,
			&pm.QuotationQuantityInBaseUnit,
			&pm.QuotationQuantityInDeliveryUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ItemBlockStatus,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		item = append(item, Item{
			Quotation:									data.Quotation,
			QuotationItem:								data.QuotationItem,
			QuotationItemCategory:						data.QuotationItemCategory,
			SupplyChainRelationshipID:					data.SupplyChainRelationshipID,
			Buyer:										data.Buyer,
			Seller:										data.Seller,
			QuotationItemText:							data.QuotationItemText,
			QuotationItemTextByBuyer:					data.QuotationItemTextByBuyer,
			QuotationItemTextBySeller:					data.QuotationItemTextBySeller,
			Product:									data.Product,
			ProductStandardID:							data.ProductStandardID,
			ProductGroup:								data.ProductGroup,
			BaseUnit:									data.BaseUnit,
			PricingDate:								data.PricingDate,
			PriceDetnExchangeRate:						data.PriceDetnExchangeRate,
			RequestedDeliveryDate:						data.RequestedDeliveryDate,
			DeliveryUnit:								data.DeliveryUnit,
			ServicesRenderingDate:						data.ServicesRenderingDate,
			QuotationQuantityInBaseUnit:				data.QuotationQuantityInBaseUnit,
			QuotationQuantityInDeliveryUnit:			data.QuotationQuantityInDeliveryUnit,
			ItemWeightUnit:								data.ItemWeightUnit,
			ProductGrossWeight:							data.ProductGrossWeight,
			ItemGrossWeight:							data.ItemGrossWeight,
			ProductNetWeight:							data.ProductNetWeight,
			ItemNetWeight:								data.ItemNetWeight,
			InternalCapacityQuantity:					data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:				data.InternalCapacityQuantityUnit,
			NetAmount:									data.NetAmount,
			TaxAmount:									data.TaxAmount,
			GrossAmount:								data.GrossAmount,
			Incoterms:									data.Incoterms,
			TransactionTaxClassification:				data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:		data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:	data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:					data.DefinedTaxClassification,
			AccountAssignmentGroup:						data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:				data.ProductAccountAssignmentGroup,
			PaymentTerms:								data.PaymentTerms,
			PaymentMethod:								data.PaymentMethod,
			Project:									data.Project,
			WBSElement:									data.WBSElement,
			AccountingExchangeRate:						data.AccountingExchangeRate,
			ReferenceDocument:							data.ReferenceDocument,
			ReferenceDocumentItem:						data.ReferenceDocumentItem,
			TaxCode:									data.TaxCode,
			TaxRate:									data.TaxRate,
			CountryOfOrigin:							data.CountryOfOrigin,
			CountryOfOriginLanguage:					data.CountryOfOriginLanguage,
			ItemBlockStatus:							data.ItemBlockStatus,
			CreationDate:								data.CreationDate,
			LastChangeDate:								data.LastChangeDate,
			IsCancelled:								data.IsCancelled,
			IsMarkedForDeletion:						data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}
	return &item, nil
}

func ConvertToItemPricingElement(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			Quotation:					data.Quotation,
			QuotationItem:				data.QuotationItem,
			PricingProcedureCounter:	data.PricingProcedureCounter,
			SupplyChainRelationshipID:	data.SupplyChainRelationshipID,
			Buyer:						data.Buyer,
			Seller:						data.Seller,
			ConditionRecord:			data.ConditionRecord,
			ConditionSequentialNumber:	data.ConditionSequentialNumber,
			ConditionType:				data.ConditionType,
			PricingDate:				data.PricingDate,
			ConditionRateValue:			data.ConditionRateValue,
			ConditionRateValueUnit:		data.ConditionRateValueUnit,
			ConditionScaleQuantity:		data.ConditionScaleQuantity,
			ConditionCurrency:			data.ConditionCurrency,
			ConditionQuantity:			data.ConditionQuantity,
			TaxCode:					data.TaxCode,
			ConditionAmount:			data.ConditionAmount,
			TransactionCurrency:		data.TransactionCurrency,
			ConditionIsManuallyChanged:	data.ConditionIsManuallyChanged,
			CreationDate:				data.CreationDate,
			LastChangeDate:				data.LastChangeDate,
			IsCancelled:				data.IsCancelled,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}
	return &itemPricingElement, nil
}

func ConvertToPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.QuotationType,
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

		data := pm
		partner = append(partner, Partner{
			Quotation:               data.Quotation,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			QuotationType:           data.QuotationType,
			Organization:            data.Organization,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &partner, nil

}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

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

		data := pm
		address = append(address, Address{
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
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &address, nil
}
