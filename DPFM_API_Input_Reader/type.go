package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	Filepath              string                `json:"filepath"`
	APIStatusCode         int                   `json:"api_status_code"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartner       *int                  `json:"business_partner"`
	ServiceLabel          string                `json:"service_label"`
	APIType               string                `json:"api_type"`
	OrdersInputParameters OrdersInputParameters `json:"OrdersInputParameters"`
	Header                Header                `json:"Orders"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type OrdersInputParameters struct {
	ReferenceDocument     *int `json:"ReferenceDocument"`
	ReferenceDocumentItem *int `json:"ReferenceDocumentItem"`
}

type Header struct {
	OrderID                          int         `json:"OrderID"`
	OrderDate                        *string     `json:"OrderDate"`
	OrderType                        *string     `json:"OrderType"`
	SupplyChainRelationshipID        *int        `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int        `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int        `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            *int        `json:"Buyer"`
	Seller                           *int        `json:"Seller"`
	BillToParty                      *int        `json:"BillToParty"`
	BillFromParty                    *int        `json:"BillFromParty"`
	BillToCountry                    *string     `json:"BillToCountry"`
	BillFromCountry                  *string     `json:"BillFromCountry"`
	Payer                            *int        `json:"Payer"`
	Payee                            *int        `json:"Payee"`
	CreationDate                     *string     `json:"CreationDate"`
	LastChangeDate                   *string     `json:"LastChangeDate"`
	ContractType                     *string     `json:"ContractType"`
	OrderValidityStartDate           *string     `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string     `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string     `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string     `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32    `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32    `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32    `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             *string     `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              *string     `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         *string     `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              *string     `json:"TransactionCurrency"`
	PricingDate                      *string     `json:"PricingDate"`
	PriceDetnExchangeRate            *float32    `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            *string     `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined  *bool       `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string     `json:"Incoterms"`
	PaymentTerms                     *string     `json:"PaymentTerms"`
	PaymentMethod                    *string     `json:"PaymentMethod"`
	ReferenceDocument                *int        `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int        `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           *string     `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32    `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string     `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool       `json:"IsExportImport"`
	HeaderText                       *string     `json:"HeaderText"`
	HeaderBlockStatus                *bool       `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool       `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool       `json:"HeaderBillingBlockStatus"`
	Item                             []Item      `json:"Item"`
	HeaderDoc                        []HeaderDoc `json:"HeaderDoc"`
	Partner                          []Partner   `json:"Partner"`
	Address                          []Address   `json:"Address"`
}

type Partner struct {
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
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
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

type HeaderDoc struct {
	OrderID                  int     `json:"OrderID"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Item struct {
	OrderID                                       int                  `json:"OrderID"`
	OrderItem                                     int                  `json:"OrderItem"`
	OrderItemCategory                             *string              `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     *int                 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int                 `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int                 `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int                 `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int                 `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 *string              `json:"OrderItemText"`
	OrderItemTextByBuyer                          *string              `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         *string              `json:"OrderItemTextBySeller"`
	Product                                       *string              `json:"Product"`
	ProductStandardID                             *string              `json:"ProductStandardID"`
	ProductGroup                                  *string              `json:"ProductGroup"`
	BaseUnit                                      *string              `json:"BaseUnit"`
	PricingDate                                   *string              `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32             `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string              `json:"RequestedDeliveryDate"`
	DeliverToParty                                *int                 `json:"DeliverToParty"`
	DeliverFromParty                              *int                 `json:"DeliverFromParty"`
	CreationDate                                  *string              `json:"CreationDate"`
	LastChangeDate                                *string              `json:"LastChangeDate"`
	DeliverToPlant                                *string              `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string              `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string              `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool                `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string              `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string              `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string              `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate            *string              `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverFromPlant                              *string              `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string              `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string              `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool                `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string              `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string              `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string              `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate          *string              `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliveryUnit                                  *string              `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int                 `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string              `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string              `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool                `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyStockConfirmationInPlant       *string              `json:"BatchMgmtPolicyStockConfirmationInPlant"`
	StockConfirmationPlantBatch                   *string              `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string              `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string              `json:"StockConfirmationPlantBatchValidityEndDate"`
	ServicesRenderingDate                         *string              `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       *float32             `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   *float32             `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                       *string              `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string              `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32             `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string              `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32             `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32             `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32             `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32             `json:"ItemNetWeight"`
	NetAmount                                     *float32             `json:"NetAmount"`
	TaxAmount                                     *float32             `json:"TaxAmount"`
	GrossAmount                                   *float32             `json:"GrossAmount"`
	InvoiceDocumentDate                           *string              `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string              `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string              `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string              `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool                `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string              `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string              `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string              `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string              `json:"ProductionPlantBatchValidityEndDate"`
	Incoterms                                     *string              `json:"Incoterms"`
	TransactionTaxClassification                  *string              `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string              `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string              `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string              `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string              `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string              `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  *string              `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string              `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string              `json:"PaymentDueDate"`
	NetPaymentDays                                *int                 `json:"NetPaymentDays"`
	PaymentMethod                                 *string              `json:"PaymentMethod"`
	Project                                       *string              `json:"Project"`
	AccountingExchangeRate                        *float32             `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int                 `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int                 `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool                `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string              `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string              `json:"IssuingStatus"`
	ReceivingStatus                               *string              `json:"ReceivingStatus"`
	ItemBillingStatus                             *string              `json:"ItemBillingStatus"`
	TaxCode                                       *string              `json:"TaxCode"`
	TaxRate                                       *float32             `json:"TaxRate"`
	CountryOfOrigin                               *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string              `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                               *bool                `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool                `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool                `json:"ItemBillingBlockStatus"`
	ItemPricingElement                            []ItemPricingElement `json:"ItemPricingElement"`
	ItemScheduleLine                              []ItemScheduleLine   `json:"ItemScheduleLine"`
}

type ItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ItemScheduleLine struct {
	OrderID                                      int      `json:"OrderID"`
	OrderItem                                    int      `json:"OrderItem"`
	ScheduleLine                                 int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                    *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID      *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                      *string  `json:"Product"`
	StockConfirmationBussinessPartner            *int     `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                       *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone               *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                        *string  `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate                        *string  `json:"ConfirmedDeliveryDate"`
	OrderQuantityInBaseUnit                      *float32 `json:"OrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck        *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	DeliveredQuantityInBaseUnit                  *float32 `json:"DeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit              *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                        *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                *string  `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus          *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
}
