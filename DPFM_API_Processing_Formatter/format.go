package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		OrderID:                         data.OrderID,
		BillToParty:                     data.BillToParty,
		BillFromParty:                   data.BillFromParty,
		BillToCountry:                   data.BillToCountry,
		BillFromCountry:                 data.BillFromCountry,
		TotalNetAmount:                  *data.TotalNetAmount,
		TotalTaxAmount:                  *data.TotalTaxAmount,
		TotalGrossAmount:                *data.TotalGrossAmount,
		TransactionCurrency:             *data.TransactionCurrency,
		PricingDate:                     *data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           *data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		Incoterms:                       data.Incoterms,
		Payer:                           data.Payer,
		Payee:                           data.Payee,
		PaymentTerms:                    *data.PaymentTerms,
		PaymentMethod:                   *data.PaymentMethod,
		InvoiceDocumentDate:             *data.InvoiceDocumentDate,
		HeaderText:                      data.HeaderText,
		HeaderBlockStatus:               data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus:       data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}
}

func ConvertToAddressUpdates(address dpfm_api_input_reader.Address) *AddressUpdates {
	data := address

	return &AddressUpdates{
		PostalCode:  *data.PostalCode,
		LocalRegion: *data.LocalRegion,
		Country:     *data.Country,
		District:    *data.District,
		StreetName:  *data.StreetName,
		CityName:    *data.CityName,
		Building:    *data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
		OrderID:                 data.OrderID,
		OrderItem:               data.OrderItem,
		OrderItemText:           *data.OrderItemText,
		OrderItemTextByBuyer:    *data.OrderItemTextByBuyer,
		OrderItemTextBySeller:   *data.OrderItemTextBySeller,
		Product:                 *data.Product,
		ProductStandardID:       *data.ProductStandardID,
		ProductGroup:            data.ProductGroup,
		RequestedDeliveryDate:   *data.RequestedDeliveryDate,
		DeliverToParty:          data.DeliverToParty,
		DeliverFromParty:        data.DeliverFromParty,
		DueCalculationBaseDate:  data.DueCalculationBaseDate,
		PaymentDueDate:          data.PaymentDueDate,
		NetPaymentDays:          data.NetPaymentDays,
		ItemDeliveryBlockStatus: data.ItemDeliveryBlockStatus,
		ItemBillingBlockStatus:  data.ItemBillingBlockStatus,
	}
}

func ConvertToItemPricingElementUpdates(itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		ConditionRateValue:         data.ConditionRateValue,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}
}

func ConvertToItemScheduleLineUpdates(itemScheduleLine dpfm_api_input_reader.ItemScheduleLine) *ItemScheduleLineUpdates {
	data := itemScheduleLine

	return &ItemScheduleLineUpdates{
		RequestedDeliveryDate:               data.RequestedDeliveryDate,
		ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
	}
}
