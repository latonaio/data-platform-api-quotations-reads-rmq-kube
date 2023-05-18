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
			&pm.DistributionChannel,
			&pm.BusinessArea,
			&pm.District,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.VaridityStartDate,
			&pm.ValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.TotalNetAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.RequestedDeliveryDate,
			&pm.BindingPeriodValidityStartDate,
			&pm.BindingPeriodValidityEndDate,
			&pm.OrderProbabilityInPercent,
			&pm.ExpectedOrderNetAmount,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.TaxClassification,
			&pm.ReferenceInquiry,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			Quotation:                      data.Quotation,
			QuotationDate:                  data.QuotationDate,
			DistributionChannel:            data.DistributionChannel,
			BusinessArea:                   data.BusinessArea,
			District:                       data.District,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ContractType:                   data.ContractType,
			VaridityStartDate:              data.VaridityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			InvoiceScheduleStartDate:       data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:         data.InvoiceScheduleEndDate,
			TotalNetAmount:                 data.TotalNetAmount,
			TransactionCurrency:            data.TransactionCurrency,
			PricingDate:                    data.PricingDate,
			RequestedDeliveryDate:          data.RequestedDeliveryDate,
			BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
			BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
			OrderProbabilityInPercent:      data.OrderProbabilityInPercent,
			ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
			Incoterms:                      data.Incoterms,
			PaymentTerms:                   data.PaymentTerms,
			PaymentMethod:                  data.PaymentMethod,
			TaxClassification:              data.TaxClassification,
			ReferenceInquiry:               data.ReferenceInquiry,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil

}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartner, error) {
	defer rows.Close()
	headerPartner := make([]HeaderPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeaderPartner{}

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
		headerPartner = append(headerPartner, HeaderPartner{
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

	return &headerPartner, nil

}

func ConvertToHeaderPartnerContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartnerContact, error) {
	defer rows.Close()
	headerPartnerContact := make([]HeaderPartnerContact, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeaderPartnerContact{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.ContactID,
			&pm.BusinessPartner,
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

		data := pm
		headerPartnerContact = append(headerPartnerContact, HeaderPartnerContact{
			Quotation:         data.Quotation,
			PartnerFunction:   data.PartnerFunction,
			ContactID:         data.ContactID,
			BusinessPartner:   data.BusinessPartner,
			ContactPersonName: data.ContactPersonName,
			EmailAddress:      data.EmailAddress,
			PhoneNumber:       data.PhoneNumber,
			MobilePhoneNumber: data.MobilePhoneNumber,
			FaxNumber:         data.FaxNumber,
			ContactTag1:       data.ContactTag1,
			ContactTag2:       data.ContactTag2,
			ContactTag3:       data.ContactTag3,
			ContactTag4:       data.ContactTag4,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}
	return &headerPartnerContact, nil
}

func ConvertToHeaderPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartnerPlant, error) {
	defer rows.Close()
	headerPartnerPlant := make([]HeaderPartnerPlant, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeaderPartnerPlant{}

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

		data := pm
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			Quotation:       data.Quotation,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &headerPartnerPlant, nil
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

func ConvertToItemPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemPartner, error) {
	defer rows.Close()
	itemPartner := make([]ItemPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPartner{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PartnerFunction,
			&pm.PartnerFunctionBusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPartner = append(itemPartner, ItemPartner{
			BusinessPartner:                data.BusinessPartner,
			Quotation:                      data.Quotation,
			QuotationItem:                  data.QuotationItem,
			PartnerFunction:                data.PartnerFunction,
			PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
			BusinessPartnerFullName:        data.BusinessPartnerFullName,
			BusinessPartnerName:            data.BusinessPartnerName,
			AddressID:                      data.AddressID,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}
	return &itemPartner, nil
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
