package dpfm_api_processing_formatter

type HeaderUpdates struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        *string  `json:"OrderDate"`
	OrderStatus                      *string  `json:"OrderStatus"`
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
	TotalNetAmount                   *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount"`
	TransactionCurrency              *string  `json:"TransactionCurrency"`
	PricingDate                      *string  `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            *string  `json:"RequestedDeliveryTime"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
}

type ItemUpdates struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderStatus                 *string  `json:"OrderStatus"`
	RequestedDeliveryDate       *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       *string  `json:"RequestedDeliveryTime"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit"`
}

type ItemPricingElementUpdates struct {
	OrderID                 int      `json:"OrderID"`
	OrderItem               int      `json:"OrderItem"`
	PricingProcedureCounter int      `json:"PricingProcedureCounter"`
	ConditionRateValue      *float32 `json:"ConditionRateValue"`
	ConditionAmount         *float32 `json:"ConditionAmount"`
}

type ItemScheduleLineUpdates struct {
	OrderID                                   int      `json:"OrderID"`
	OrderItem                                 int      `json:"OrderItem"`
	ScheduleLine                              int      `json:"ScheduleLine"`
	RequestedDeliveryDate                     *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                     *string  `json:"RequestedDeliveryTime"`
	ScheduleLineOrderQuantityInBaseUnit       *float32 `json:"ScheduleLineOrderQuantityInBaseUnit"`
	ExternalReferenceDocument                 *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem             *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemScheduleLine *string  `json:"ExternalReferenceDocumentItemScheduleLine"`
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
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
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
