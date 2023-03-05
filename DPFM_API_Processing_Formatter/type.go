package dpfm_api_processing_formatter

type HeaderUpdates struct {
	BillToParty               *int     `json:"BillToParty"`
	BillFromParty             *int     `json:"BillFromParty"`
	BillToCountry             *string  `json:"BillToCountry"`
	BillFromCountry           *string  `json:"BillFromCountry"`
	Payer                     *int     `json:"Payer"`
	Payee                     *int     `json:"Payee"`
	ContractType              *string  `json:"ContractType"`
	OrderValidityStartDate    *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate      *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate    *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate      *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount            float32  `json:"TotalNetAmount"`
	TotalTaxAmount            float32  `json:"TotalTaxAmount"`
	TotalGrossAmount          float32  `json:"TotalGrossAmount"`
	TransactionCurrency       string   `json:"TransactionCurrency"`
	PricingDate               string   `json:"PricingDate"`
	PriceDetnExchangeRate     *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate     string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime     string   `json:"RequestedDeliveryTime"`
	Incoterms                 *string  `json:"Incoterms"`
	PaymentTerms              string   `json:"PaymentTerms"`
	PaymentMethod             string   `json:"PaymentMethod"`
	AccountingExchangeRate    *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate       string   `json:"InvoiceDocumentDate"`
	HeaderText                *string  `json:"HeaderText"`
	HeaderBlockStatus         *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus  *bool    `json:"HeaderBillingBlockStatus"`
}

type AddressUpdates struct {
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type ItemUpdates struct {
	OrderItem                                int      `json:"OrderItem"`
	OrderItemCategory                        string   `json:"OrderItemCategory"`
	SupplyChainRelationshipID                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID  *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                            string   `json:"OrderItemText"`
	OrderItemTextByBuyer                     string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                    string   `json:"OrderItemTextBySeller"`
	Product                                  string   `json:"Product"`
	ProductStandardID                        string   `json:"ProductStandardID"`
	ProductGroup                             *string  `json:"ProductGroup"`
	RequestedDeliveryDate                    string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                    string   `json:"RequestedDeliveryTime"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation            *string  `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantBatch                      *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate     *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime     *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate       *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime       *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation          *string  `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantBatch                    *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartTime   *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliveryUnit                             string   `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner         *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                   *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch              *string  `json:"StockConfirmationPlantBatch"`
	OrderQuantityInBaseUnit                  float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit              float32  `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                  *string  `json:"StockConfirmationPolicy"`
	ItemWeightUnit                           *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                       *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                          *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                         *float32 `json:"ProductNetWeight"`
	ItemNetWeight                            *float32 `json:"ItemNetWeight"`
	TaxAmount                                *float32 `json:"TaxAmount"`
	GrossAmount                              *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                      *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation           *string  `json:"ProductionPlantStorageLocation"`
	ProductionPlantBatch                     *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate    *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime    *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate      *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime      *string  `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlan                           *int     `json:"InspectionPlan"`
	InspectionPlant                          *string  `json:"InspectionPlant"`
	InspectionOrder                          *int     `json:"InspectionOrder"`
	Incoterms                                *string  `json:"Incoterms"`
	TransactionTaxClassification             string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry    string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry  string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                 string   `json:"DefinedTaxClassification"`
	PaymentTerms                             string   `json:"PaymentTerms"`
	DueCalculationBaseDate                   *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                           *string  `json:"PaymentDueDate"`
	NetPaymentDays                           *int     `json:"NetPaymentDays"`
	Project                                  *string  `json:"Project"`
	ItemBlockStatus                          *bool    `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                  *bool    `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                   *bool    `json:"ItemBillingBlockStatus"`
}

type ItemPricingElementUpdates struct {
	ConditionRateValue *float32 `json:"ConditionRateValue"`
	ConditionAmount    *float32 `json:"ConditionAmount"`
}

type ItemScheduleLineUpdates struct {
	RequestedDeliveryDate               string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime               string  `json:"RequestedDeliveryTime"`
	ScheduleLineOrderQuantity           float32 `json:"ScheduleLineOrderQuantity"`
	ItemScheduleLineDeliveryBlockStatus *bool   `json:"ItemScheduleLineDeliveryBlockStatus"`
}

type PartnerUpdates struct {
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}
