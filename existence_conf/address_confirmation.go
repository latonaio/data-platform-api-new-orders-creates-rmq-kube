package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) addressExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	address := input.Header.Address
	for _, address := range address {
		addressID, validityEndDate, err := getAddressExistenceConfKey(mapper, &input.Header, &address, exconfErrMsg)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		queueName, err := getQueueName(mapper)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		wg2.Add(1)
		exReqTimes++
		go func() {
			res, err := c.addressExistenceConfRequest(addressID, validityEndDate, queueName, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) addressExistenceConfRequest(addressID int, validityEndDate string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"AddressID":       addressID,
		"ValidityEndDate": validityEndDate,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.AddressReturn.AddressID = addressID
	req.AddressReturn.ValidityEndDate = validityEndDate

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getAddressExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, address *dpfm_api_input_reader.Address, exconfErrMsg *string) (int, string, error) {
	var addressID int
	var validityEndDate string
	var err error

	if header.OrderValidityEndDate == nil {
		err = xerrors.Errorf("cannot specify null keys")
		return 0, "", err
	}
	if header.OrderValidityEndDate != nil {
		if len(*header.OrderValidityEndDate) == 0 {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, "", err
		}
	}

	addressID = address.AddressID
	validityEndDate = *header.OrderValidityEndDate

	return addressID, validityEndDate, nil
}
