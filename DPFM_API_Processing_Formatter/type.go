package dpfm_api_processing_formatter

type HeaderUpdates struct {
	OrderID                         int      `json:"OrderID"`
	BillToParty                     *int     `json:"BillToParty"`
	BillFromParty                   *int     `json:"BillFromParty"`
	BillToCountry                   *string  `json:"BillToCountry"`
	BillFromCountry                 *string  `json:"BillFromCountry"`
	Payer                           *int     `json:"Payer"`
	Payee                           *int     `json:"Payee"`
	TotalNetAmount                  float32  `json:"TotalNetAmount"`
	TotalTaxAmount                  float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                float32  `json:"TotalGrossAmount"`
	TransactionCurrency             string   `json:"TransactionCurrency"`
	PricingDate                     string   `json:"PricingDate"`
	PriceDetnExchangeRate           *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           string   `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                       *string  `json:"Incoterms"`
	PaymentTerms                    string   `json:"PaymentTerms"`
	PaymentMethod                   string   `json:"PaymentMethod"`
	InvoiceDocumentDate             string   `json:"InvoiceDocumentDate"`
	HeaderText                      *string  `json:"HeaderText"`
	HeaderBlockStatus               *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus       *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus        *bool    `json:"HeaderBillingBlockStatus"`
}

type AddressUpdates struct {
	PostalCode  string `json:"PostalCode"`
	LocalRegion string `json:"LocalRegion"`
	Country     string `json:"Country"`
	District    string `json:"District"`
	StreetName  string `json:"StreetName"`
	CityName    string `json:"CityName"`
	Building    string `json:"Building"`
	Floor       *int   `json:"Floor"`
	Room        *int   `json:"Room"`
}

type ItemUpdates struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	OrderItemText           string  `json:"OrderItemText"`
	OrderItemTextByBuyer    string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller   string  `json:"OrderItemTextBySeller"`
	Product                 string  `json:"Product"`
	ProductStandardID       string  `json:"ProductStandardID"`
	ProductGroup            *string `json:"ProductGroup"`
	RequestedDeliveryDate   string  `json:"RequestedDeliveryDate"`
	DeliverToParty          *int    `json:"DeliverToParty"`
	DeliverFromParty        *int    `json:"DeliverFromParty"`
	DueCalculationBaseDate  *string `json:"DueCalculationBaseDate"`
	PaymentDueDate          *string `json:"PaymentDueDate"`
	NetPaymentDays          *int    `json:"NetPaymentDays"`
	ItemDeliveryBlockStatus *bool   `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus  *bool   `json:"ItemBillingBlockStatus"`
}

type ItemPricingElementUpdates struct {
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ItemScheduleLineUpdates struct {
	RequestedDeliveryDate               *string `json:"RequestedDeliveryDate"`
	ItemScheduleLineDeliveryBlockStatus *bool   `json:"ItemScheduleLineDeliveryBlockStatus"`
}
