package dpfm_api_output_formatter

import (
	"data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemPricingElementCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemPricingElement, error) {
	itemPricingElements := make([]ItemPricingElement, 0)

	for _, data := range *subfuncSDC.Message.ItemPricingElement {
		itemPricingElement, err := TypeConverter[*ItemPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemPricingElements = append(itemPricingElements, *itemPricingElement)
	}

	return &itemPricingElements, nil
}

func ConvertToItemScheduleLineCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemScheduleLine, error) {
	itemScheduleLines := make([]ItemScheduleLine, 0)

	for _, data := range *subfuncSDC.Message.ItemScheduleLine {
		itemScheduleLine, err := TypeConverter[*ItemScheduleLine](data)
		if err != nil {
			return nil, err
		}

		itemScheduleLines = append(itemScheduleLines, *itemScheduleLine)
	}

	return &itemScheduleLines, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemPricingElementUpdates(itemPricingElementUpdates *[]dpfm_api_processing_formatter.ItemPricingElementUpdates) (*[]ItemPricingElement, error) {
	itemPricingElements := make([]ItemPricingElement, 0)

	for _, data := range *itemPricingElementUpdates {
		itemPricingElement, err := TypeConverter[*ItemPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemPricingElements = append(itemPricingElements, *itemPricingElement)
	}

	return &itemPricingElements, nil
}

func ConvertToItemScheduleLineUpdates(itemScheduleLineUpdates *[]dpfm_api_processing_formatter.ItemScheduleLineUpdates) (*[]ItemScheduleLine, error) {
	itemScheduleLines := make([]ItemScheduleLine, 0)

	for _, data := range *itemScheduleLineUpdates {
		itemScheduleLine, err := TypeConverter[*ItemScheduleLine](data)
		if err != nil {
			return nil, err
		}

		itemScheduleLines = append(itemScheduleLines, *itemScheduleLine)
	}

	return &itemScheduleLines, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func ConvertToHeaderFromQuotations(
	rows *sql.Rows,
	orderIssuedID int,
) (*[]sub_func_complementer.Header, error) {
	defer rows.Close()
	header := make([]sub_func_complementer.Header, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsHeader{}

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
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.ReferenceDocument,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderIsClosed,
			&pm.HeaderBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
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
		headerCompleteDeliveryIsDefined := false

		header = append(header, sub_func_complementer.Header{
			OrderID:                         orderIssuedID,
			OrderDate:                       formattedDate,
			OrderType:                       data.QuotationType,
			OrderStatus:                     "DFT",
			SupplyChainRelationshipID:       data.SupplyChainRelationshipID,
			Buyer:                           data.Buyer,
			Seller:                          data.Seller,
			TotalGrossAmount:                data.TotalGrossAmount,
			HeaderDeliveryStatus:            "NP",
			HeaderBillingStatus:             "NP",
			HeaderDocReferenceStatus:        "NP",
			TransactionCurrency:             data.TransactionCurrency,
			PricingDate:                     data.PricingDate,
			RequestedDeliveryDate:           data.RequestedDeliveryDate,
			RequestedDeliveryTime:           "00:00:00",
			HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
			Incoterms:                       data.Incoterms,
			PaymentTerms:                    data.PaymentTerms,
			PaymentMethod:                   data.PaymentMethod,
			Contract:                        data.Contract,
			ContractItem:                    data.ContractItem,
			Project:                         data.Project,
			WBSElement:                      data.WBSElement,
			AccountAssignmentGroup:          data.AccountAssignmentGroup,
			InvoiceDocumentDate:             *data.InvoiceDocumentDate,
			HeaderText:                      data.HeaderText,
			CreationDate:                    formattedDate,
			CreationTime:                    formattedTime,
			LastChangeDate:                  formattedDate,
			LastChangeTime:                  formattedTime,
			IsCancelled:                     &isCancelled,
			IsMarkedForDeletion:             &isMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil
}

func ConvertToItemFromQuotations(
	rows *sql.Rows,
	orderIssuedID int,
) (*[]sub_func_complementer.Item, error) {
	defer rows.Close()
	item := make([]sub_func_complementer.Item, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsItem{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.QuotationItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.QuotationItemText,
			&pm.QuotationItemTextByBuyer,
			&pm.QuotationItemTextBySeller,
			&pm.Product,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.ProductSpecification,
			&pm.MarkingOfMaterial,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.DeliveryUnit,
			&pm.ServicesRenderingDate,
			&pm.QuotationQuantityInBaseUnit,
			&pm.QuotationQuantityInDeliveryUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ItemBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.InspectionPlan,
			&pm.InspectionLot,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemCompleteDeliveryIsDefined := false
		itemDeliveryStatus := "NP"
		itemDeliveryBlockStatus := false
		deliverToPlant := "TE01"
		deliverFromPlant := "TE01"
		productWeightUnit := "KG"

		item = append(item, sub_func_complementer.Item{
			OrderID:                                 orderIssuedID,
			OrderItem:                               data.QuotationItem,
			OrderItemCategory:                       data.QuotationItemCategory,
			OrderStatus:                             "DFT",
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       &data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryPlantID:  &data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: &data.SupplyChainRelationshipID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			DeliverToParty:                          &data.Buyer,
			DeliverFromParty:                        &data.Seller,
			DeliverToPlant:                          &deliverToPlant,
			DeliverFromPlant:                        &deliverFromPlant,
			OrderItemText:                           data.QuotationItemText,
			OrderItemTextByBuyer:                    data.QuotationItemTextByBuyer,
			OrderItemTextBySeller:                   data.QuotationItemTextBySeller,
			Product:                                 data.Product,
			SizeOrDimensionText:                     data.SizeOrDimensionText,
			ProductSpecification:                    data.ProductSpecification,
			MarkingOfMaterial:                       data.MarkingOfMaterial,
			BaseUnit:                                data.BaseUnit,
			DeliveryUnit:                            data.DeliveryUnit,
			PricingDate:                             data.PricingDate,
			RequestedDeliveryDate:                   data.RequestedDeliveryDate,
			RequestedDeliveryTime:                   "00:00:00",
			OrderQuantityInBaseUnit:                 data.QuotationQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:             data.QuotationQuantityInDeliveryUnit,
			QuantityPerPackage:                      1,
			ItemGrossWeight:                         data.ProductNetWeight,
			ItemNetWeight:                           data.ProductNetWeight,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			InspectionPlantBusinessPartner:          data.InspectionPlantBusinessPartner,
			InspectionPlant:                         data.InspectionPlant,
			InspectionPlan:                          data.InspectionPlan,
			InspectionLot:                           data.InspectionLot,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   *data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: *data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                *data.DefinedTaxClassification,
			AccountAssignmentGroup:                  data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:           data.ProductAccountAssignmentGroup,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			ProductNetWeight:                        data.ProductNetWeight,
			ProductWeightUnit:                       &productWeightUnit,
			ItemCompleteDeliveryIsDefined:           &itemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                      &itemDeliveryStatus,
			ItemDeliveryBlockStatus:                 &itemDeliveryBlockStatus,
			CreationDate:                            formattedDate,
			CreationTime:                            formattedTime,
			LastChangeDate:                          formattedDate,
			LastChangeTime:                          formattedTime,
			IsCancelled:                             &isCancelled,
			IsMarkedForDeletion:                     &isMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &item, nil
}

func ConvertToItemPricingElementFromQuotations(
	rows *sql.Rows,
	orderIssuedID int,
) (*[]sub_func_complementer.ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]sub_func_complementer.ItemPricingElement, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsItemPricingElement{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
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

		itemPricingElement = append(itemPricingElement, sub_func_complementer.ItemPricingElement{
			OrderID:                   orderIssuedID,
			OrderItem:                 data.QuotationItem,
			PricingProcedureCounter:   data.PricingProcedureCounter,
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			ConditionRecord:           data.ConditionRecord,
			ConditionSequentialNumber: data.ConditionSequentialNumber,
			ConditionType:             data.ConditionType,
			PricingDate:               data.PricingDate,
			ConditionRateValue:        data.ConditionRateValue,
			ConditionRateValueUnit:    data.ConditionRateValueUnit,
			ConditionScaleQuantity:    data.ConditionScaleQuantity,
			ConditionCurrency:         data.ConditionCurrency,
			ConditionQuantity:         data.ConditionQuantity,
			ConditionAmount:           data.ConditionAmount,
			TransactionCurrency:       data.TransactionCurrency,
			CreationDate:              formattedDate,
			CreationTime:              formattedTime,
			LastChangeDate:            formattedDate,
			LastChangeTime:            formattedTime,
			IsCancelled:               &isCancelled,
			IsMarkedForDeletion:       &isMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &itemPricingElement, nil
}

func ConvertToItemScheduleLineFromQuotations(
	complementerOrdersItems *[]sub_func_complementer.Item,
) (*[]sub_func_complementer.ItemScheduleLine, error) {
	itemScheduleLine := make([]sub_func_complementer.ItemScheduleLine, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	for _, data := range *complementerOrdersItems {
		itemScheduleLine = append(itemScheduleLine, sub_func_complementer.ItemScheduleLine{
			OrderID:                                 data.OrderID,
			OrderItem:                               data.OrderItem,
			ScheduleLine:                            data.OrderItem,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipID,
			Product:                                 data.Product,
			StockConfirmationBussinessPartner:       data.Seller,
			//StockConfirmationPlant:                  *data.DeliverToPlant,
			StockConfirmationPlant:              "TE01",
			RequestedDeliveryDate:               data.RequestedDeliveryDate,
			RequestedDeliveryTime:               "00:00:00",
			ScheduleLineOrderQuantityInBaseUnit: data.OrderQuantityInBaseUnit,
			OriginalOrderQuantityInBaseUnit:     data.OrderQuantityInBaseUnit,
			CreationDate:                        formattedDate,
			CreationTime:                        formattedTime,
			LastChangeDate:                      formattedDate,
			LastChangeTime:                      formattedTime,
			IsCancelled:                         &isCancelled,
			IsMarkedForDeletion:                 &isMarkedForDeletion,
		})
	}

	return &itemScheduleLine, nil
}

func ConvertToPartnerFromQuotations(
	rows *sql.Rows,
	orderIssuedID int,
) (*[]sub_func_complementer.Partner, error) {
	defer rows.Close()
	partner := make([]sub_func_complementer.Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsPartner{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
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

		partner = append(partner, sub_func_complementer.Partner{
			OrderID:                 orderIssuedID,
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &partner, nil
}

func ConvertToAddressFromQuotations(
	rows *sql.Rows,
	orderIssuedID int,
) (*[]sub_func_complementer.Address, error) {
	defer rows.Close()
	address := make([]sub_func_complementer.Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsAddress{}

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

		address = append(address, sub_func_complementer.Address{
			OrderID:     orderIssuedID, // TODO 加算された自動採番
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
