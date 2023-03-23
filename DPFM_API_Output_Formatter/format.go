package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

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
