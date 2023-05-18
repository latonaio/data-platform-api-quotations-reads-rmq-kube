package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-quotations-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
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
	var header *[]dpfm_api_output_formatter.Header
	var headerPartner *[]dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *[]dpfm_api_output_formatter.HeaderPartnerContact
	var headerPartnerPlant *[]dpfm_api_output_formatter.HeaderPartnerPlant
	var address *[]dpfm_api_output_formatter.Address
	var item *[]dpfm_api_output_formatter.Item
	var itemPartner *[]dpfm_api_output_formatter.ItemPartner
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
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
) *[]dpfm_api_output_formatter.Header {
	quotation := input.Header.Quotation

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_data
		WHERE Quotation = ?;`, quotation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.HeaderPartner {
	var args []interface{}
	quotation := input.Header.Quotation
	partner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range partner {
		args = append(args, quotation, v.PartnerFunction)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_data
		WHERE (Quotation, PartnerFunction) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.HeaderPartnerContact {
	var args []interface{}
	quotation := input.Header.Quotation
	partner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range partner {
		partnercontact := v.HeaderPartnerContact
		for _, w := range partnercontact {
			args = append(args, quotation, v.PartnerFunction, w.ContactID)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_contact_data
		WHERE (Quotation, PartnerFunction, ContactID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.HeaderPartnerPlant {
	var args []interface{}
	quotation := input.Header.Quotation
	partner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range partner {
		partnerplant := v.HeaderPartnerPlant
		for _, w := range partnerplant {
			args = append(args, quotation, v.PartnerFunction, w.BusinessPartner)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_partner_plant_data
		WHERE (Quotation, PartnerFunction, BusinessPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	quotation := input.Header.Quotation
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, quotation, v.AddressID)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_address_data
		WHERE (Quotation, AddressID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	quotation := input.Header.Quotation
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, quotation, v.QuotationItem)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_data
		WHERE (Quotation, QuotationItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.ItemPartner {
	var args []interface{}
	quotation := input.Header.Quotation
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemPartner := v.ItemPartner
		for _, w := range itemPartner {
			args = append(args, quotation, v.QuotationItem, w.PartnerFunction)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_partner_data
		WHERE (Quotation, QuotationItem, PartnerFunction)  IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

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
) *[]dpfm_api_output_formatter.ItemPricingElement {
	var args []interface{}
	quotation := input.Header.Quotation
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemPricingElement := v.ItemPricingElement
		for _, w := range itemPricingElement {
			args = append(args, quotation, v.QuotationItem, w.PricingProcedureStep, w.PricingProcedureCounter)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_pricing_element_data
		WHERE (Quotation, QuotationItem, PricingProcedureStep, PricingProcedureCounter) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
