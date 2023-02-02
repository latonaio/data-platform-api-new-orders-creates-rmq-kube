package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) itemPlantGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		plant, bpID, err := getItemPlantGeneralExistenceConfKey(mapper, &item, exconfErrMsg)
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
			res, err := c.plantGeneralExistenceConfRequest(plant, bpID, queueName, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) plantGeneralExistenceConfRequest(plant string, bpID int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
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
	req.PlantGeneralReturn.Plant = plant
	req.PlantGeneralReturn.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getItemPlantGeneralExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (string, int, error) {
	var plant string
	var bpID int
	var err error

	switch mapper.Field {
	case "DeliverToPlant":
		if item.DeliverToPlant == nil || item.DeliverToParty == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return "", 0, err
		}
		if item.DeliverToPlant != nil {
			if len(*item.DeliverToPlant) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return "", 0, err
			}
		}
		plant = *item.DeliverToPlant
		bpID = *item.DeliverToParty
	case "DeliverFromPlant":
		if item.DeliverFromPlant == nil || item.DeliverFromParty == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return "", 0, err
		}
		if item.DeliverFromPlant != nil {
			if len(*item.DeliverFromPlant) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return "", 0, err
			}
		}
		plant = *item.DeliverFromPlant
		bpID = *item.DeliverFromParty
	}

	return plant, bpID, nil
}
