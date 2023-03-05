package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
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
		TotalNetAmount:            *data.TotalNetAmount,
		TotalTaxAmount:            *data.TotalTaxAmount,
		TotalGrossAmount:          *data.TotalGrossAmount,
		TransactionCurrency:       *data.TransactionCurrency,
		PricingDate:               *data.PricingDate,
		PriceDetnExchangeRate:     data.PriceDetnExchangeRate,
		RequestedDeliveryDate:     *data.RequestedDeliveryDate,
		RequestedDeliveryTime:     *data.RequestedDeliveryTime,
		Incoterms:                 data.Incoterms,
		PaymentTerms:              *data.PaymentTerms,
		PaymentMethod:             *data.PaymentMethod,
		AccountingExchangeRate:    data.AccountingExchangeRate,
		InvoiceDocumentDate:       *data.InvoiceDocumentDate,
		HeaderText:                data.HeaderText,
		HeaderBlockStatus:         data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus: data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:  data.HeaderBillingBlockStatus,
	}
}

func ConvertToAddressUpdates(address dpfm_api_input_reader.Address) *AddressUpdates {
	data := address

	return &AddressUpdates{
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

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
		OrderItem:                                data.OrderItem,
		OrderItemCategory:                        *data.OrderItemCategory,
		SupplyChainRelationshipID:                *data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipStockConfPlantID:  data.SupplyChainRelationshipStockConfPlantID,
		SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
		OrderItemText:                            *data.OrderItemText,
		OrderItemTextByBuyer:                     *data.OrderItemTextByBuyer,
		OrderItemTextBySeller:                    *data.OrderItemTextBySeller,
		Product:                                  *data.Product,
		ProductStandardID:                        *data.ProductStandardID,
		ProductGroup:                             data.ProductGroup,
		RequestedDeliveryDate:                    *data.RequestedDeliveryDate,
		RequestedDeliveryTime:                    *data.RequestedDeliveryTime,
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
		DeliveryUnit:                             *data.DeliveryUnit,
		StockConfirmationBusinessPartner:         data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:                   data.StockConfirmationPlant,
		StockConfirmationPlantBatch:              data.StockConfirmationPlantBatch,
		OrderQuantityInBaseUnit:                  *data.OrderQuantityInBaseUnit,
		OrderQuantityInDeliveryUnit:              *data.OrderQuantityInDeliveryUnit,
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
		TransactionTaxClassification:             *data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:    *data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry:  *data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                 *data.DefinedTaxClassification,
		PaymentTerms:                             *data.PaymentTerms,
		DueCalculationBaseDate:                   data.DueCalculationBaseDate,
		PaymentDueDate:                           data.PaymentDueDate,
		NetPaymentDays:                           data.NetPaymentDays,
		Project:                                  data.Project,
		ItemBlockStatus:                          data.ItemBlockStatus,
		ItemDeliveryBlockStatus:                  data.ItemDeliveryBlockStatus,
		ItemBillingBlockStatus:                   data.ItemBillingBlockStatus,
	}
}

func ConvertToItemPricingElementUpdates(itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		ConditionRateValue: data.ConditionRateValue,
		ConditionAmount:    data.ConditionAmount,
	}
}

func ConvertToItemScheduleLineUpdates(itemScheduleLine dpfm_api_input_reader.ItemScheduleLine) *ItemScheduleLineUpdates {
	data := itemScheduleLine

	return &ItemScheduleLineUpdates{
		RequestedDeliveryDate:               *data.RequestedDeliveryDate,
		RequestedDeliveryTime:               *data.RequestedDeliveryTime,
		ScheduleLineOrderQuantity:           *data.ScheduleLineOrderQuantity,
		ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
	}
}

func ConvertToPartnerUpdates(partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	data := partner

	return &PartnerUpdates{
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}
