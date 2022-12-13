package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var headerPartner *dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *dpfm_api_output_formatter.HeaderPartnerContact
	var headerPartnerPlant *dpfm_api_output_formatter.HeaderPartnerPlant
	var address *dpfm_api_output_formatter.Address
	var item *dpfm_api_output_formatter.Item
	var itemPartner *dpfm_api_output_formatter.ItemPartner
	var itemPricingElement *dpfm_api_output_formatter.ItemPricingElement
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderPartner":
			func() {
				headerPartner = c.HeaderPartner(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerContact":
			func() {
				headerPartnerContact = c.HeaderPartnerContact(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerPlant":
			func() {
				headerPartnerPlant = c.HeaderPartnerPlant(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "ItemPartner":
			func() {
				itemPartner = c.ItemPartner(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		HeaderPartner:        headerPartner,
		HeaderPartnerContact: headerPartnerContact,
		HeaderPartnerPlant:   headerPartnerPlant,
		Address:              address,
		Item:                 item,
		ItemPartner:          itemPartner,
		ItemPricingElement:   itemPricingElement,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	quotation := input.Header.Quotation

	rows, err := c.db.Query(
		`SELECT Quotation, QuotationDate, QuoationType, Buyer, Seller, CreationDate, LastChangeDate, 
		ContractType, VaridityStartDate, ValidityEndDate, InvoiceScheduleStartDate, InvoiceScheduleEndDate, 
		TotalNetAmount, TotalTaxAmount, TotalGrossAmount, TransactionCurrency, PricingDate, RequestedDeliveryDate, 
		BindingPeriodValidityStartDate, BindingPeriodValidityEndDate, Incoterms, PaymentTerms, 
		PaymentMethod, AccountingExchangeRate, BillingDocumentDate, HeaderText, ReferenceInquiry
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_data
		WHERE Quotations = ?;`, quotation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartner {
	quotation := input.Header.Quotation
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT Quotation, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, 
		Language, Organization, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_data
		WHERE (Quotation, PartnerFunction, BusinessPartner) = (?, ?, ?);`, quotation, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerContact {
	quotation := input.Header.Quotation
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	contactID := input.Header.HeaderPartner.HeaderPartnerContact.ContactID

	rows, err := c.db.Query(
		`SELECT Quotation, PartnerFunction, BusinessPartner, ContactID, ContactPersonName, EmailAddress, PhoneNumber, 
		MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_contact_data
		WHERE (Quotation, PartnerFunction, BusinessPartner, ContactID) = (?, ?, ?, ?);`, quotation, partnerFunction, businessPartner, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerPlant {
	quotation := input.Header.Quotation
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	plant := input.Header.HeaderPartner.HeaderPartnerPlant.Plant

	rows, err := c.db.Query(
		`SELECT Quotation, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_plant_data
		WHERE (Quotation, PartnerFunction, BusinessPartner, Plant) = (?, ?, ?, ?);`, quotation, partnerFunction, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Address {
	quotation := input.Header.Quotation
	addressID := input.Header.Address.AddressID

	rows, err := c.db.Query(
		`SELECT Quotation, AddressID, PostalCode, LocalRegion, Country, District, StreetName, 
		CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_address_data
		WHERE (Quotation, AddressID) = (?, ?);`, quotation, addressID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Item {
	quotation := input.Header.Quotation
	quotationItem := input.Header.Item.QuotationItem

	rows, err := c.db.Query(
		`SELECT Quotation, QuotationItem, QuotationItemCategory, QuotationItemText, Product, ProductStandardID, 
		PricingDate, QuotationQuantity, QuotationQuantityUnit, ItemGrossWeight, ItemNetWeight, ItemWeightUnit, NetAmount, 
		TaxAmount, GrossAmount, ProductGroup, BillingDocumentDate, ProductionPlantBusinessPartner, ProductionPlant, 
		ProductionPlantTimeZone, ProductionPlantStorageLocation, IssuingPlantBusinessPartner, IssuingPlant, IssuingPlantTimeZone, 
		IssuingPlantStorageLocation, ReceivingPlantBusinessPartner, ReceivingPlant, ReceivingPlantTimeZone, 
		ReceivingPlantStorageLocation, Incoterms, BPTaxClassification, ProductTaxClassification, ProductAccountAssignmentGroup, 
		PaymentTerms, PaymentMethod, Project, AccountingExchangeRate, ReferenceInquiry, ReferenceDocumentItem, ItemDeliveryStatus, 
		IssuingStatus, ReceivingStatus, BillingStatus, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_data
		WHERE (Quotation, QuotationItem) = (?, ?);`, quotation, quotationItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartner {
	quotation := input.Header.Quotation
	quotationItem := input.Header.Item.QuotationItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT Quotation, QuotationItem, PartnerFunction, BusinessPartner, BusinessPartnerFullName, 
		BusinessPartnerName, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_partner_data
		WHERE (Quotation, QuotationItem, PartnerFunction, BusinessPartner) = (?, ?, ?, ?);`, quotation, quotationItem, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPricingElement {
	quotation := input.Header.Quotation
	quotationItem := input.Header.Item.QuotationItem
	pricingProcedureStep := input.Header.Item.ItemPricingElement.PricingProcedureStep
	pricingProcedureCounter := input.Header.Item.ItemPricingElement.PricingProcedureCounter

	rows, err := c.db.Query(
		`SELECT Quotation, QuotationItem, PricingProcedureStep, PricingProcedureCounter, ConditionType, 
		PricingDate, ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit, 
		ConditionRecord, ConditionSequentialNumber, TaxCode, ConditionAmount, TransactionCurrency, ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_pricing_element_data
		WHERE (Quotation, QuotationItem, PricingProcedureStep, PricingProcedureCounter) = (?, ?, ?, ?);`, quotation, quotationItem, pricingProcedureStep, pricingProcedureCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
