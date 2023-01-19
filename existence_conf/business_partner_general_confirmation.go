package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) bpGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		bpID, err := getBPGeneralExistenceConfKey(mapper, &header, exconfErrMsg)
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
			res, err := c.bpGeneralExistenceConfRequest(bpID, queueName, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) bpGeneralExistenceConfRequest(bpID int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
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
	req.BPGeneralReturn.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func getBPGeneralExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, error) {
	var bpID int
	var err error

	switch mapper.Field {
	case "Buyer":
		if header.Buyer == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, err
		}
		bpID = *header.Buyer
	case "Seller":
		if header.Seller == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, err
		}
		bpID = *header.Seller
	}

	return bpID, nil
}
