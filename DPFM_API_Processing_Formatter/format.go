package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		OrderID:                          data.OrderID,
		OrderDate:                        data.OrderDate,
		OrderStatus:                      data.OrderStatus,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		Payer:                            data.Payer,
		Payee:                            data.Payee,
		OrderValidityStartDate:           data.OrderValidityStartDate,
		OrderValidityEndDate:             data.OrderValidityEndDate,
		InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
		TotalNetAmount:                   data.TotalNetAmount,
		TotalTaxAmount:                   data.TotalTaxAmount,
		TotalGrossAmount:                 data.TotalGrossAmount,
		TransactionCurrency:              data.TransactionCurrency,
		PricingDate:                      data.PricingDate,
		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
		RequestedDeliveryDate:            data.RequestedDeliveryDate,
		RequestedDeliveryTime:            data.RequestedDeliveryTime,
		Incoterms:                        data.Incoterms,
		PaymentTerms:                     data.PaymentTerms,
		PaymentMethod:                    data.PaymentMethod,
		AccountingExchangeRate:           data.AccountingExchangeRate,
		InvoiceDocumentDate:              data.InvoiceDocumentDate,
		HeaderText:                       data.HeaderText,
		HeaderBlockStatus:                data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus:        data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:         data.HeaderBillingBlockStatus,
		ExternalReferenceDocument:        data.ExternalReferenceDocument,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item
	return &ItemUpdates{
		OrderID:                     dataHeader.OrderID,
		OrderItem:                   data.OrderItem,
		OrderStatus:                 data.OrderStatus,
		RequestedDeliveryDate:       data.RequestedDeliveryDate,
		RequestedDeliveryTime:       data.RequestedDeliveryTime,
		OrderQuantityInBaseUnit:     data.OrderQuantityInBaseUnit,
		OrderQuantityInDeliveryUnit: data.OrderQuantityInDeliveryUnit,
	}
}

func ConvertToItemPricingElementUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	dataHeader := header
	dataItem := item
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		OrderID:                 dataHeader.OrderID,
		OrderItem:               dataItem.OrderItem,
		PricingProcedureCounter: data.PricingProcedureCounter,
		ConditionRateValue:      data.ConditionRateValue,
		ConditionAmount:         data.ConditionAmount,
	}
}

func ConvertToItemScheduleLineUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemScheduleLine dpfm_api_input_reader.ItemScheduleLine) *ItemScheduleLineUpdates {
	dataHeader := header
	dataItem := item
	data := itemScheduleLine

	return &ItemScheduleLineUpdates{
		OrderID:                                   dataHeader.OrderID,
		OrderItem:                                 dataItem.OrderItem,
		ScheduleLine:                              data.ScheduleLine,
		RequestedDeliveryDate:                     data.RequestedDeliveryDate,
		RequestedDeliveryTime:                     data.RequestedDeliveryTime,
		ScheduleLineOrderQuantityInBaseUnit:       data.ScheduleLineOrderQuantityInBaseUnit,
		ExternalReferenceDocument:                 data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem:             data.ExternalReferenceDocumentItem,
		ExternalReferenceDocumentItemScheduleLine: data.ExternalReferenceDocumentItemScheduleLine,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		OrderID:                 dataHeader.OrderID,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		OrderID:     dataHeader.OrderID,
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
