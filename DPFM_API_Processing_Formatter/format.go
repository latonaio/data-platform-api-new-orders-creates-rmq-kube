package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		OrderID:                   data.OrderID,
		BillToParty:               data.BillToParty,
		BillFromParty:             data.BillFromParty,
		BillToCountry:             data.BillToCountry,
		BillFromCountry:           data.BillFromCountry,
		Payer:                     data.Payer,
		Payee:                     data.Payee,
		ContractType:              data.ContractType,
		OrderValidityStartDate:    data.OrderValidityStartDate,
		OrderValidityEndDate:      data.OrderValidityEndDate,
		InvoicePeriodStartDate:    data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:      data.InvoicePeriodEndDate,
		TotalNetAmount:            data.TotalNetAmount,
		TotalTaxAmount:            data.TotalTaxAmount,
		TotalGrossAmount:          data.TotalGrossAmount,
		TransactionCurrency:       data.TransactionCurrency,
		PricingDate:               data.PricingDate,
		PriceDetnExchangeRate:     data.PriceDetnExchangeRate,
		RequestedDeliveryDate:     data.RequestedDeliveryDate,
		RequestedDeliveryTime:     data.RequestedDeliveryTime,
		Incoterms:                 data.Incoterms,
		PaymentTerms:              data.PaymentTerms,
		PaymentMethod:             data.PaymentMethod,
		AccountingExchangeRate:    data.AccountingExchangeRate,
		InvoiceDocumentDate:       data.InvoiceDocumentDate,
		HeaderText:                data.HeaderText,
		HeaderBlockStatus:         data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus: data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:  data.HeaderBillingBlockStatus,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item
	return &ItemUpdates{
		OrderID:                                  dataHeader.OrderID,
		OrderItem:                                data.OrderItem,
		OrderItemCategory:                        data.OrderItemCategory,
		SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipStockConfPlantID:  data.SupplyChainRelationshipStockConfPlantID,
		SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
		OrderItemText:                            data.OrderItemText,
		OrderItemTextByBuyer:                     data.OrderItemTextByBuyer,
		OrderItemTextBySeller:                    data.OrderItemTextBySeller,
		Product:                                  data.Product,
		ProductStandardID:                        data.ProductStandardID,
		ProductGroup:                             data.ProductGroup,
		RequestedDeliveryDate:                    data.RequestedDeliveryDate,
		RequestedDeliveryTime:                    data.RequestedDeliveryTime,
		DeliverToParty:                           data.DeliverToParty,
		DeliverFromParty:                         data.DeliverFromParty,
		DeliverToPlant:                           data.DeliverToPlant,
		DeliverToPlantStorageLocation:            data.DeliverToPlantStorageLocation,
		DeliverToPlantBatch:                      data.DeliverToPlantBatch,
		DeliverToPlantBatchValidityStartDate:     data.DeliverToPlantBatchValidityStartDate,
		DeliverToPlantBatchValidityStartTime:     data.DeliverToPlantBatchValidityStartTime,
		DeliverToPlantBatchValidityEndDate:       data.DeliverToPlantBatchValidityEndDate,
		DeliverToPlantBatchValidityEndTime:       data.DeliverToPlantBatchValidityEndTime,
		DeliverFromPlant:                         data.DeliverFromPlant,
		DeliverFromPlantStorageLocation:          data.DeliverFromPlantStorageLocation,
		DeliverFromPlantBatch:                    data.DeliverFromPlantBatch,
		DeliverFromPlantBatchValidityStartTime:   data.DeliverFromPlantBatchValidityStartTime,
		DeliveryUnit:                             data.DeliveryUnit,
		StockConfirmationBusinessPartner:         data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:                   data.StockConfirmationPlant,
		StockConfirmationPlantBatch:              data.StockConfirmationPlantBatch,
		OrderQuantityInBaseUnit:                  data.OrderQuantityInBaseUnit,
		OrderQuantityInDeliveryUnit:              data.OrderQuantityInDeliveryUnit,
		StockConfirmationPolicy:                  data.StockConfirmationPolicy,
		ItemWeightUnit:                           data.ItemWeightUnit,
		ProductGrossWeight:                       data.ProductGrossWeight,
		ItemGrossWeight:                          data.ItemGrossWeight,
		ProductNetWeight:                         data.ProductNetWeight,
		ItemNetWeight:                            data.ItemNetWeight,
		TaxAmount:                                data.TaxAmount,
		GrossAmount:                              data.GrossAmount,
		InvoiceDocumentDate:                      data.InvoiceDocumentDate,
		ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
		ProductionPlant:                          data.ProductionPlant,
		ProductionPlantStorageLocation:           data.ProductionPlantStorageLocation,
		ProductionPlantBatch:                     data.ProductionPlantBatch,
		ProductionPlantBatchValidityStartDate:    data.ProductionPlantBatchValidityStartDate,
		ProductionPlantBatchValidityStartTime:    data.ProductionPlantBatchValidityStartTime,
		ProductionPlantBatchValidityEndDate:      data.ProductionPlantBatchValidityEndDate,
		ProductionPlantBatchValidityEndTime:      data.ProductionPlantBatchValidityEndTime,
		InspectionPlan:                           data.InspectionPlan,
		InspectionPlant:                          data.InspectionPlant,
		InspectionOrder:                          data.InspectionOrder,
		Incoterms:                                data.Incoterms,
		TransactionTaxClassification:             data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:    data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry:  data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                 data.DefinedTaxClassification,
		PaymentTerms:                             data.PaymentTerms,
		DueCalculationBaseDate:                   data.DueCalculationBaseDate,
		PaymentDueDate:                           data.PaymentDueDate,
		NetPaymentDays:                           data.NetPaymentDays,
		Project:                                  data.Project,
		ItemBlockStatus:                          data.ItemBlockStatus,
		ItemDeliveryBlockStatus:                  data.ItemDeliveryBlockStatus,
		ItemBillingBlockStatus:                   data.ItemBillingBlockStatus,
	}
}

func ConvertToItemPricingElementUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	dataHeader := header
	dataItem := item
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		OrderID:                   dataHeader.OrderID,
		OrderItem:                 dataItem.OrderItem,
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		PricingProcedureCounter:   data.PricingProcedureCounter,
		ConditionRateValue:        data.ConditionRateValue,
		ConditionAmount:           data.ConditionAmount,
	}
}

func ConvertToItemScheduleLineUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemScheduleLine dpfm_api_input_reader.ItemScheduleLine) *ItemScheduleLineUpdates {
	dataHeader := header
	dataItem := item
	data := itemScheduleLine

	return &ItemScheduleLineUpdates{
		OrderID:                             dataHeader.OrderID,
		OrderItem:                           dataItem.OrderItem,
		ScheduleLine:                        data.ScheduleLine,
		RequestedDeliveryDate:               data.RequestedDeliveryDate,
		RequestedDeliveryTime:               data.RequestedDeliveryTime,
		ScheduleLineOrderQuantity:           data.ScheduleLineOrderQuantity,
		ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
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
