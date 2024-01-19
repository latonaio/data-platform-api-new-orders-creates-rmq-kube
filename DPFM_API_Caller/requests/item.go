package requests

type Item struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderStatus                 *string  `json:"OrderStatus"`
	RequestedDeliveryDate       *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       *string  `json:"RequestedDeliveryTime"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit"`
}
