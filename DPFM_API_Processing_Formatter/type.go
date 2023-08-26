package dpfm_api_processing_formatter

type HeaderUpdates struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        string   `json:"OrderDate"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            string   `json:"RequestedDeliveryTime"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
}

type ItemUpdates struct {
	OrderID                                       int      `json:"OrderID"`
	OrderItem                                     int      `json:"OrderItem"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	OrderItemText                                 string   `json:"OrderItemText"`
	OrderItemTextByBuyer                          string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         string   `json:"OrderItemTextBySeller"`
	Product                                       string   `json:"Product"`
	BillOfMaterial                                *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int     `json:"BillOfMaterialItem"`
	RequestedDeliveryDate                         string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                         string   `json:"RequestedDeliveryTime"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   float32  `json:"OrderQuantityInDeliveryUnit"`
	QuantityPerPackage                            float32  `json:"QuantityPerPackage"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	TaxAmount                                     float32  `json:"TaxAmount"`
	GrossAmount                                   float32  `json:"GrossAmount"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	InspectionPlan                                *int     `json:"InspectionPlan"`
	InspectionPlant                               *string  `json:"InspectionPlant"`
	InspectionOrder                               *int     `json:"InspectionOrder"`
	Incoterms                                     *string  `json:"Incoterms"`
	TransactionTaxClassification                  string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      string   `json:"DefinedTaxClassification"`
	PaymentTerms                                  string   `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 string   `json:"PaymentMethod"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	Equipment				                      *int	   `json:"Equipment"`
	PlannedFreight				                  *int	   `json:"PlannedFreight"`
	FreightOrder				                  *int	   `json:"FreightOrder"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
}

type ItemPricingElementUpdates struct {
	OrderID                   int      `json:"OrderID"`
	OrderItem                 int      `json:"OrderItem"`
	SupplyChainRelationshipID int      `json:"SupplyChainRelationshipID"`
	Buyer                     int      `json:"Buyer"`
	Seller                    int      `json:"Seller"`
	PricingProcedureCounter   int      `json:"PricingProcedureCounter"`
	ConditionRateValue        *float32 `json:"ConditionRateValue"`
	ConditionAmount           *float32 `json:"ConditionAmount"`
}

type ItemScheduleLineUpdates struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string   `json:"RequestedDeliveryTime"`
	ScheduleLineOrderQuantityInBaseUnit             float32  `json:"ScheduleLineOrderQuantityInBaseUnit"`
}

type PartnerUpdates struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
	OrderID     int     `json:"OrderID"`
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
