package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var itemScheduleLine *[]dpfm_api_output_formatter.ItemScheduleLine
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner

	handleAccepter := referenceDocumentHandle(
		input,
		subfuncSDC,
		accepter,
	)

	var calculateOrderIDQueryGets *sub_func_complementer.CalculateOrderIDQueryGets
	var orderIssuedID int

	if len(handleAccepter) > 0 {
		calculateOrderIDQueryGets = c.CalculateOrderID(errs)

		if calculateOrderIDQueryGets == nil {
			err := xerrors.Errorf("calculateOrderIDQueryGets is nil")
			*errs = append(*errs, err)
			return nil
		}

		orderIssuedID = calculateOrderIDQueryGets.OrderIDLatestNumber + 1

		subfuncSDCHeader := c.getQuotationsHeader(
			input,
			errs,
			orderIssuedID,
		)

		if len(*subfuncSDCHeader) > 0 {
			subfuncSDC.Message.Header = &(*subfuncSDCHeader)[0]
		}

		subfuncSDCItem := c.getQuotationsItem(
			input,
			errs,
			orderIssuedID,
		)

		if len(*subfuncSDCItem) > 0 {
			subfuncSDC.Message.Item = subfuncSDCItem
		}

		subfuncSDCItemPricingElement := c.getQuotationsItemPricingElement(
			input,
			errs,
			orderIssuedID,
		)

		if len(*subfuncSDCItemPricingElement) > 0 {
			subfuncSDC.Message.ItemPricingElement = subfuncSDCItemPricingElement
		}

		subfuncSDCPartner := c.getQuotationsPartner(
			input,
			errs,
			orderIssuedID,
		)

		if len(*subfuncSDCPartner) > 0 {
			subfuncSDC.Message.Partner = subfuncSDCPartner
		}

		// itemScheduleLine
		subfuncSDCItemScheduleLine := c.setOrdersItemDataForItemScheduleLine(
			subfuncSDCItem,
			errs,
		)

		if len(*subfuncSDCItemScheduleLine) > 0 {
			subfuncSDC.Message.ItemScheduleLine = subfuncSDCItemScheduleLine
		}
	}

	for _, fn := range handleAccepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemPricingElement":
			itemPricingElement = c.itemPricingElementCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemScheduleLine":
			itemScheduleLine = c.itemScheduleLineCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Partner":
			partner = c.partnerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Address":
			address = c.addressCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	if calculateOrderIDQueryGets != nil {
		err := c.UpdateLatestNumber(errs, orderIssuedID)
		if err != nil {
			*errs = append(*errs, err)
			return nil
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		ItemScheduleLine:   itemScheduleLine,
		Address:            address,
		Partner:            partner,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var itemScheduleLine *[]dpfm_api_output_formatter.ItemScheduleLine
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "ItemPricingElement":
			itemPricingElement = c.itemPricingElementUpdateSql(mtx, input, output, errs, log)
		case "ItemScheduleLine":
			itemScheduleLine = c.itemScheduleLineUpdateSql(mtx, input, output, errs, log)
		case "Partner":
			partner = c.partnerUpdateSql(mtx, input, output, errs, log)
		case "Address":
			address = c.addressUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		ItemScheduleLine:   itemScheduleLine,
		Partner:            partner,
		Address:            address,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_header_dataの更新
	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_item_dataの更新
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "OrdersItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemPricingElementCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_item_pricing_element_dataの更新
	for _, itemPricingElementData := range *subfuncSDC.Message.ItemPricingElement {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemPricingElementData, "function": "OrdersItemPricingElement", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Pricing Element Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElementCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemScheduleLineCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemScheduleLine {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_item_schedule_line_dataの更新
	for _, itemScheduleLineData := range *subfuncSDC.Message.ItemScheduleLine {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemScheduleLineData, "function": "OrdersItemScheduleLine", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Schedule Line Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemScheduleLineCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_partner_dataの更新
	for _, partnerData := range *subfuncSDC.Message.Partner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "OrdersPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Partner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_address_dataの更新
	for _, addressData := range *subfuncSDC.Message.Address {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "OrdersAddress", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Address Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	req := make([]dpfm_api_processing_formatter.ItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		itemData := *dpfm_api_processing_formatter.ConvertToItemUpdates(header, item)

		if itemIsUpdate(&itemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "OrdersItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot update"
				return nil
			}
		}
		req = append(req, itemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemPricingElementUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	req := make([]dpfm_api_processing_formatter.ItemPricingElementUpdates, 0)
	//sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemPricingElement := range item.ItemPricingElement {
			itemPricingElementData := *dpfm_api_processing_formatter.ConvertToItemPricingElementUpdates(header, item, itemPricingElement)

			//if itemPricingElementIsUpdate(&itemPricingElementData) {
			//	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemPricingElementData, "function": "OrdersItemPricingElement", "runtime_session_id": sessionID})
			//	if err != nil {
			//		err = xerrors.Errorf("rmq error: %w", err)
			//		*errs = append(*errs, err)
			//		return nil
			//	}
			//	res.Success()
			//	if !checkResult(res) {
			//		output.SQLUpdateResult = getBoolPtr(false)
			//		output.SQLUpdateError = "Item Pricing Element Data cannot update"
			//		return nil
			//	}
			//}
			req = append(req, itemPricingElementData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElementUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemScheduleLineUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemScheduleLine {
	req := make([]dpfm_api_processing_formatter.ItemScheduleLineUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemScheduleLine := range item.ItemScheduleLine {
			itemScheduleLineData := *dpfm_api_processing_formatter.ConvertToItemScheduleLineUpdates(header, item, itemScheduleLine)

			if itemScheduleLineIsUpdate(&itemScheduleLineData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemScheduleLineData, "function": "OrdersItemScheduleLine", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Item Schedule Line Data cannot update"
					return nil
				}
			}
			req = append(req, itemScheduleLineData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemScheduleLineUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	req := make([]dpfm_api_processing_formatter.PartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, partner := range header.Partner {
		partnerData := *dpfm_api_processing_formatter.ConvertToPartnerUpdates(header, partner)

		if partnerIsUpdate(&partnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "OrdersPartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Partner Data cannot update"
				return nil
			}
		}
		req = append(req, partnerData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	req := make([]dpfm_api_processing_formatter.AddressUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, address := range header.Address {
		addressData := *dpfm_api_processing_formatter.ConvertToAddressUpdates(header, address)

		if addressIsUpdate(&addressData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "OrdersAddress", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Address Data cannot update"
				return nil
			}
		}
		req = append(req, addressData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	orderID := header.OrderID

	return !(orderID == 0)
}

func itemIsUpdate(item *dpfm_api_processing_formatter.ItemUpdates) bool {
	orderID := item.OrderID
	orderItem := item.OrderItem

	return !(orderID == 0 || orderItem == 0)
}

func itemPricingElementIsUpdate(itemPricingElement *dpfm_api_processing_formatter.ItemPricingElementUpdates) bool {
	orderID := itemPricingElement.OrderID
	orderItem := itemPricingElement.OrderItem
	pricingProcedureCounter := itemPricingElement.PricingProcedureCounter

	return !(orderID == 0 || orderItem == 0 || pricingProcedureCounter == 0)
}

func itemScheduleLineIsUpdate(itemScheduleLine *dpfm_api_processing_formatter.ItemScheduleLineUpdates) bool {
	orderID := itemScheduleLine.OrderID
	orderItem := itemScheduleLine.OrderItem
	scheduleLine := itemScheduleLine.ScheduleLine

	return !(orderID == 0 || orderItem == 0 || scheduleLine == 0)
}

func partnerIsUpdate(partner *dpfm_api_processing_formatter.PartnerUpdates) bool {
	orderID := partner.OrderID
	partnerFunction := partner.PartnerFunction
	businessPartner := partner.BusinessPartner

	return !(orderID == 0 || partnerFunction == "" || businessPartner == 0)
}

func addressIsUpdate(address *dpfm_api_processing_formatter.AddressUpdates) bool {
	orderID := address.OrderID
	addressID := address.AddressID

	return !(orderID == 0 || addressID == 0)
}

func referenceDocumentHandle(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
) []string {
	var handleAccepter []string

	if input.InputParameters.ReferenceDocument != nil {
		handleAccepter = append(handleAccepter, "Header")
		handleAccepter = append(handleAccepter, "Item")
		handleAccepter = append(handleAccepter, "ItemPricingElement")
		handleAccepter = append(handleAccepter, "ItemScheduleLine")
		handleAccepter = append(handleAccepter, "Partner")
	} else {
		handleAccepter = accepter
	}

	return handleAccepter
}

func (c *DPFMAPICaller) getQuotationsHeader(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	orderIssuedID int,
) *[]sub_func_complementer.Header {
	quotation := input.InputParameters.ReferenceDocument

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

	data, err := dpfm_api_output_formatter.ConvertToHeaderFromQuotations(rows, orderIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) getQuotationsItem(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	orderIssuedID int,
) *[]sub_func_complementer.Item {
	quotation := input.InputParameters.ReferenceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_data
		WHERE Quotation = ?;`, quotation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemFromQuotations(rows, orderIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) getQuotationsItemPricingElement(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	orderIssuedID int,
) *[]sub_func_complementer.ItemPricingElement {
	quotation := input.InputParameters.ReferenceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_item_pricing_element_data
		WHERE Quotation = ?;`, quotation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElementFromQuotations(rows, orderIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) getQuotationsPartner(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	orderIssuedID int,
) *[]sub_func_complementer.Partner {
	quotation := input.InputParameters.ReferenceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_partner_data
		WHERE Quotation = ?;`, quotation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartnerFromQuotations(rows, orderIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) setOrdersItemDataForItemScheduleLine(
	complementerOrdersItem *[]sub_func_complementer.Item,
	errs *[]error,
) *[]sub_func_complementer.ItemScheduleLine {
	data, err := dpfm_api_output_formatter.ConvertToItemScheduleLineFromQuotations(complementerOrdersItem)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) CalculateOrderID(
	errs *[]error,
) *sub_func_complementer.CalculateOrderIDQueryGets {
	pm := &sub_func_complementer.CalculateOrderIDQueryGets{}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, "ORDERS", "OrderID",
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				*errs = append(*errs, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。"))
				return nil
			} else {
				break
			}
		}
		err = rows.Scan(
			&pm.NumberRangeID,
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.OrderIDLatestNumber,
		)
		if err != nil {
			*errs = append(*errs, err)
			return nil
		}
	}

	return pm
}

func (c *DPFMAPICaller) UpdateLatestNumber(
	errs *[]error,
	orderIssuedID int,
) error {
	//rows, err := c.db.Query(
	//	`SELECT *
	//	FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
	//	WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, "ORDERS", "OrderID",
	//)

	_, err := c.db.Exec(`
			UPDATE data_platform_number_range_latest_number_data SET LatestNumber=(?)
			WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`,
		orderIssuedID,
		"ORDERS",
		"OrderID",
	)
	if err != nil {
		return xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルの更新に失敗しました。")
	}

	return nil
}
