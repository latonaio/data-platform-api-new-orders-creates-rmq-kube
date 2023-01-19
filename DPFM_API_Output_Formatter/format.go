package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderFromCreates(subfuncSDC *sub_func_complementer.SDC) *Header {
	data := subfuncSDC.Message.Header

	header := &Header{
		OrderID:                          data.OrderID,
		OrderDate:                        data.OrderDate,
		OrderType:                        data.OrderType,
		SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		Buyer:                            data.Buyer,
		Seller:                           data.Seller,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		BillToCountry:                    data.BillToCountry,
		BillFromCountry:                  data.BillFromCountry,
		Payer:                            data.Payer,
		Payee:                            data.Payee,
		CreationDate:                     data.CreationDate,
		LastChangeDate:                   data.LastChangeDate,
		ContractType:                     data.ContractType,
		OrderValidityStartDate:           data.OrderValidityStartDate,
		OrderValidityEndDate:             data.OrderValidityEndDate,
		InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
		TotalNetAmount:                   data.TotalNetAmount,
		TotalTaxAmount:                   data.TotalTaxAmount,
		TotalGrossAmount:                 data.TotalGrossAmount,
		HeaderDeliveryStatus:             data.HeaderDeliveryStatus,
		HeaderBillingStatus:              data.HeaderBillingStatus,
		HeaderDocReferenceStatus:         data.HeaderDocReferenceStatus,
		TransactionCurrency:              data.TransactionCurrency,
		PricingDate:                      data.PricingDate,
		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
		RequestedDeliveryDate:            data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined:  data.HeaderCompleteDeliveryIsDefined,
		Incoterms:                        data.Incoterms,
		PaymentTerms:                     data.PaymentTerms,
		PaymentMethod:                    data.PaymentMethod,
		ReferenceDocument:                data.ReferenceDocument,
		ReferenceDocumentItem:            data.ReferenceDocumentItem,
		AccountAssignmentGroup:           data.AccountAssignmentGroup,
		AccountingExchangeRate:           data.AccountingExchangeRate,
		InvoiceDocumentDate:              data.InvoiceDocumentDate,
		IsExportImport:                   data.IsExportImport,
		HeaderText:                       data.HeaderText,
		HeaderBlockStatus:                data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus:        data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:         data.HeaderBillingBlockStatus,
	}

	return header
}

func ConvertToHeaderFromUpdates(headerUpdates *dpfm_api_processing_formatter.HeaderUpdates) *Header {
	data := headerUpdates

	header := &Header{
		OrderID:                         data.OrderID,
		BillToParty:                     data.BillToParty,
		BillFromParty:                   data.BillFromParty,
		BillToCountry:                   data.BillToCountry,
		BillFromCountry:                 data.BillFromCountry,
		Payer:                           data.Payer,
		Payee:                           data.Payee,
		TotalNetAmount:                  data.TotalNetAmount,
		TotalTaxAmount:                  data.TotalTaxAmount,
		TotalGrossAmount:                data.TotalGrossAmount,
		TransactionCurrency:             data.TransactionCurrency,
		PricingDate:                     data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		Incoterms:                       data.Incoterms,
		PaymentTerms:                    data.PaymentTerms,
		PaymentMethod:                   data.PaymentMethod,
		InvoiceDocumentDate:             data.InvoiceDocumentDate,
		HeaderText:                      data.HeaderText,
		HeaderBlockStatus:               data.HeaderBlockStatus,
		HeaderDeliveryBlockStatus:       data.HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return header
}

func ConvertToPartner(subfuncSDC *sub_func_complementer.SDC) *[]Partner {
	var partner []Partner

	for _, data := range *subfuncSDC.Message.Partner {
		partner = append(partner, Partner{
			OrderID:                 data.OrderID,
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
		})
	}

	return &partner
}

func ConvertToAddress(subfuncSDC *sub_func_complementer.SDC) *[]Address {
	var address []Address

	for _, data := range *subfuncSDC.Message.Address {
		address = append(address, Address{
			OrderID:     data.OrderID,
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

	return &address
}

func ConvertToItemFromCreates(subfuncSDC *sub_func_complementer.SDC) *[]Item {
	var item []Item

	for _, data := range *subfuncSDC.Message.Item {
		item = append(item, Item{
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderItemCategory:                             data.OrderItemCategory,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			OrderItemText:                                 data.OrderItemText,
			OrderItemTextByBuyer:                          data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                         data.OrderItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			PricingDate:                                   data.PricingDate,
			PriceDetnExchangeRate:                         data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                         data.RequestedDeliveryDate,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			CreationDate:                                  data.CreationDate,
			LastChangeDate:                                data.LastChangeDate,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverToPlantTimeZone:                        data.DeliverToPlantTimeZone,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverFromPlant:                              data.DeliverFromPlant,
			DeliverFromPlantTimeZone:                      data.DeliverFromPlantTimeZone,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliveryUnit:                                  data.DeliveryUnit,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                data.StockConfirmationPlantTimeZone,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyStockConfirmationInPlant:       data.BatchMgmtPolicyStockConfirmationInPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			ServicesRenderingDate:                         data.ServicesRenderingDate,
			OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:                   data.OrderQuantityInDeliveryUnit,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			InvoiceDocumentDate:                           data.InvoiceDocumentDate,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			Incoterms:                                     data.Incoterms,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			Project:                                       data.Project,
			AccountingExchangeRate:                        data.AccountingExchangeRate,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                            data.ItemDeliveryStatus,
			IssuingStatus:                                 data.IssuingStatus,
			ReceivingStatus:                               data.ReceivingStatus,
			ItemBillingStatus:                             data.ItemBillingStatus,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			ItemBlockStatus:                               data.ItemBlockStatus,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
		})
	}

	return &item
}

func ConvertToItemFromUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) *[]Item {
	var item []Item

	for _, data := range *itemUpdates {
		item = append(item, Item{
			OrderID:                 data.OrderID,
			OrderItem:               data.OrderItem,
			OrderItemText:           data.OrderItemText,
			OrderItemTextByBuyer:    data.OrderItemTextByBuyer,
			OrderItemTextBySeller:   data.OrderItemTextBySeller,
			Product:                 data.Product,
			ProductStandardID:       data.ProductStandardID,
			ProductGroup:            data.ProductGroup,
			RequestedDeliveryDate:   data.RequestedDeliveryDate,
			DeliverToParty:          data.DeliverToParty,
			DeliverFromParty:        data.DeliverFromParty,
			DueCalculationBaseDate:  data.DueCalculationBaseDate,
			PaymentDueDate:          data.PaymentDueDate,
			NetPaymentDays:          data.NetPaymentDays,
			ItemDeliveryBlockStatus: data.ItemDeliveryBlockStatus,
			ItemBillingBlockStatus:  data.ItemBillingBlockStatus,
		})
	}

	return &item
}

func ConvertToItemPricingElementFromCreates(subfuncSDC *sub_func_complementer.SDC) *[]ItemPricingElement {
	var itemPricingElement []ItemPricingElement

	for _, data := range *subfuncSDC.Message.ItemPricingElement {

		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}

	return &itemPricingElement
}

func ConvertToItemPricingElementFromUpdates(itemPricingElementUpdates *[]dpfm_api_processing_formatter.ItemPricingElementUpdates) *[]ItemPricingElement {
	var itemPricingElement []ItemPricingElement

	for _, data := range *itemPricingElementUpdates {
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			ConditionRateValue:         data.ConditionRateValue,
			ConditionAmount:            data.ConditionAmount,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}

	return &itemPricingElement
}

func ConvertToItemScheduleLine(subfuncSDC *sub_func_complementer.SDC) *[]ItemScheduleLine {
	var itemScheduleLine []ItemScheduleLine

	for _, data := range *subfuncSDC.Message.ItemScheduleLine {

		itemScheduleLine = append(itemScheduleLine, ItemScheduleLine{
			OrderID:                                      data.OrderID,
			OrderItem:                                    data.OrderItem,
			ScheduleLine:                                 data.ScheduleLine,
			SupplyChainRelationshipID:                    data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID:      data.SupplyChainRelationshipStockConfPlantID,
			Product:                                      data.Product,
			StockConfirmationBussinessPartner:            data.StockConfirmationBussinessPartner,
			StockConfirmationPlant:                       data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:               data.StockConfirmationPlantTimeZone,
			StockConfirmationPlantBatch:                  data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
			RequestedDeliveryDate:                        data.RequestedDeliveryDate,
			ConfirmedDeliveryDate:                        data.ConfirmedDeliveryDate,
			OrderQuantityInBaseUnit:                      data.OrderQuantityInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheck:        data.ConfirmedOrderQuantityByPDTAvailCheck,
			DeliveredQuantityInBaseUnit:                  data.DeliveredQuantityInBaseUnit,
			OpenConfirmedQuantityInBaseUnit:              data.OpenConfirmedQuantityInBaseUnit,
			StockIsFullyConfirmed:                        data.StockIsFullyConfirmed,
			PlusMinusFlag:                                data.PlusMinusFlag,
			ItemScheduleLineDeliveryBlockStatus:          data.ItemScheduleLineDeliveryBlockStatus,
		})
	}

	return &itemScheduleLine
}
