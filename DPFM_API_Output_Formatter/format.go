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
			&pm.QuotationItemText,
			&pm.Product,
			&pm.QuotationQuantity,
			&pm.QuotationQuantityUnit,
			&pm.ItemOrderProbabilityInPercent,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.TransactionCurrency,
			&pm.BussinessParnterCurrency,
			&pm.NetAmount,
			&pm.ProductGroup,
			&pm.ProductPricingGroup,
			&pm.Batch,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.ProductTaxClassification1,
			&pm.Project,
			&pm.ProfitCenter,
			&pm.ReferenceInquiry,
			&pm.ReferenceInquiryItem,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		item = append(item, Item{
			Quotation:                     data.Quotation,
			QuotationItem:                 data.QuotationItem,
			QuotationItemCategory:         data.QuotationItemCategory,
			QuotationItemText:             data.QuotationItemText,
			Product:                       data.Product,
			QuotationQuantity:             data.QuotationQuantity,
			QuotationQuantityUnit:         data.QuotationQuantityUnit,
			ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
			ItemGrossWeight:               data.ItemGrossWeight,
			ItemNetWeight:                 data.ItemNetWeight,
			ItemWeightUnit:                data.ItemWeightUnit,
			TransactionCurrency:           data.TransactionCurrency,
			BussinessParnterCurrency:      data.BussinessParnterCurrency,
			NetAmount:                     data.NetAmount,
			ProductGroup:                  data.ProductGroup,
			ProductPricingGroup:           data.ProductPricingGroup,
			Batch:                         data.Batch,
			IssuingPlant:                  data.IssuingPlant,
			IssuingPlantStorageLocation:   data.IssuingPlantStorageLocation,
			ReceivingPlant:                data.ReceivingPlant,
			ReceivingPlantStorageLocation: data.ReceivingPlantStorageLocation,
			Incoterms:                     data.Incoterms,
			PaymentTerms:                  data.PaymentTerms,
			ProductTaxClassification1:     data.ProductTaxClassification1,
			Project:                       data.Project,
			ProfitCenter:                  data.ProfitCenter,
			ReferenceInquiry:              data.ReferenceInquiry,
			ReferenceInquiryItem:          data.ReferenceInquiryItem,
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
			&pm.BusinessPartner,
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PricingProcedureStep,
			&pm.PricingProcedureCounter,
			&pm.ConditionType,
			&pm.PriceConditionDeterminationDte,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.TaxCode,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			BusinessPartner:                data.BusinessPartner,
			Quotation:                      data.Quotation,
			QuotationItem:                  data.QuotationItem,
			PricingProcedureStep:           data.PricingProcedureStep,
			PricingProcedureCounter:        data.PricingProcedureCounter,
			ConditionType:                  data.ConditionType,
			PriceConditionDeterminationDte: data.PriceConditionDeterminationDte,
			ConditionRateValue:             data.ConditionRateValue,
			ConditionCurrency:              data.ConditionCurrency,
			ConditionRecord:                data.ConditionRecord,
			ConditionSequentialNumber:      data.ConditionSequentialNumber,
			TaxCode:                        data.TaxCode,
			TransactionCurrency:            data.TransactionCurrency,
			ConditionIsManuallyChanged:     data.ConditionIsManuallyChanged,
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
