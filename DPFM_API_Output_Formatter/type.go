package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header             *Header               `json:"Header"`
	Item               *[]Item               `json:"Item"`
	ItemPricingElement *[]ItemPricingElement `json:"ItemPricingElement"`
	ItemScheduleLine   *[]ItemScheduleLine   `json:"ItemScheduleLine"`
	Partner            *[]Partner            `json:"Partner"`
	Address            *[]Address            `json:"Address"`
}

type Header struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        string   `json:"OrderDate"`
	OrderType                        string   `json:"OrderType"`
	OrderStatus                      string   `json:"OrderStatus"`
	SupplyChainRelationshipID        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int      `json:"Buyer"`
	Seller                           int      `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *string  `json:"BillToCountry"`
	BillFromCountry                  *string  `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	ContractType                     *string  `json:"ContractType"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             string   `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            string   `json:"RequestedDeliveryTime"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	Contract                         *int     `json:"Contract"`
	ContractItem                     *int     `json:"ContractItem"`
	ProductionVersion                *int     `json:"ProductionVersion"`
	ProductionVersionItem            *int     `json:"ProductionVersionItem"`
	ProductionOrder                  *int     `json:"ProductionOrder"`
	ProductionOrderItem              *int     `json:"ProductionOrderItem"`
	Operations                       *int     `json:"Operations"`
	OperationsItem                   *int     `json:"OperationsItem"`
	OperationID                      *int     `json:"OperationID"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain        *string  `json:"CertificateAuthorityChain"`
	UsageControlChain                *string  `json:"UsageControlChain"`
	CreationDate                     string   `json:"CreationDate"`
	CreationTime                     string   `json:"CreationTime"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	LastChangeTime                   string   `json:"LastChangeTime"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderStatus                 *string  `json:"OrderStatus"`
	RequestedDeliveryDate       *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       *string  `json:"RequestedDeliveryTime"`
	OrderQuantityInBaseUnit     *float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit"`
}

type ItemPricingElement struct {
	OrderID                    int     `json:"OrderID"`
	OrderItem                  int     `json:"OrderItem"`
	PricingProcedureCounter    int     `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  int     `json:"SupplyChainRelationshipID"`
	Buyer                      int     `json:"Buyer"`
	Seller                     int     `json:"Seller"`
	ConditionRecord            int     `json:"ConditionRecord"`
	ConditionSequentialNumber  int     `json:"ConditionSequentialNumber"`
	ConditionType              string  `json:"ConditionType"`
	PricingDate                string  `json:"PricingDate"`
	ConditionRateValue         float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          string  `json:"ConditionCurrency"`
	ConditionQuantity          float32 `json:"ConditionQuantity"`
	TaxCode                    *string `json:"TaxCode"`
	ConditionAmount            float32 `json:"ConditionAmount"`
	TransactionCurrency        string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool   `json:"ConditionIsManuallyChanged"`
	CreationDate               string  `json:"CreationDate"`
	CreationTime               string  `json:"CreationTime"`
	LastChangeDate             string  `json:"LastChangeDate"`
	LastChangeTime             string  `json:"LastChangeTime"`
	IsCancelled                *bool   `json:"IsCancelled"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

type ItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	StockConfirmationBussinessPartner               int      `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string   `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                           *string  `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryTime                           *string  `json:"ConfirmedDeliveryTime"`
	ScheduleLineOrderQuantityInBaseUnit             float32  `json:"ScheduleLineOrderQuantityInBaseUnit"`
	OriginalOrderQuantityInBaseUnit                 float32  `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                   *string  `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	ExternalReferenceDocument                       *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                   *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemScheduleLine       *string  `json:"ExternalReferenceDocumentItemScheduleLine"`
	CreationDate                                    string   `json:"CreationDate"`
	CreationTime                                    string   `json:"CreationTime"`
	LastChangeDate                                  string   `json:"LastChangeDate"`
	LastChangeTime                                  string   `json:"LastChangeTime"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
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
	EmailAddress            *string `json:"EmailAddress"`
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
