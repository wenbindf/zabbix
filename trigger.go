package zabbix

type Tigger struct {
	Triggerid   string      `json:"triggerid"`
	Description string      `json:"description"`
	Expression  string      `json:"expression"`
	Priority    IntegerType `json:"priority"`
}

const (
	NotClassified IntegerType = 0
	Information   IntegerType = 1
	Warning       IntegerType = 2
	Average       IntegerType = 3
	High          IntegerType = 4
	Disaster      IntegerType = 5
)

type Tiggers []Tigger
type IntegerType int

func (api *API) TiggerCreate(tiggers Tiggers) (err error) {
	response, err := api.CallWithError("trigger.create", tiggers)
	if err != nil {
		return
	}
	result := response.Result.(map[string]interface{})
	triggerids := result["triggerids"].([]interface{})
	for i, id := range triggerids {
		tiggers[i].Triggerid = id.(string)
	}
	return
}

func (api *API) TiggerDelete(tiggers Tiggers) (err error) {
	ids := make([]string, len(tiggers))
	for i, tigger := range tiggers {
		ids[i] = tigger.Triggerid
	}

	err = api.TiggersDeleteByIds(ids)
	if err == nil {
		for i := range tiggers {
			tiggers[i].Triggerid = ""
		}
	}
	return
}
func (api *API) TiggersDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("trigger.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	tiggerids1, ok := result["triggerids"].([]interface{})
	l := len(tiggerids1)
	if !ok {
		// some versions actually return map there
		tiggerids2 := result["triggerids"].(map[string]interface{})
		l = len(tiggerids2)
	}
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}

type TiggerExitsArgs struct {
	Expression string `json:"expression"`
	Host       string `json:"host"`
	HostId     string `json:"hostid"`
	//	Description string `json:"description"`
}

func (api *API) TiggerExits(tiggerExitsArgs TiggerExitsArgs) (bool, error) {
	response, err := api.CallWithError("trigger.exists", tiggerExitsArgs)
	if err != nil {
		return false, err
	}
	result := response.Result.(bool)
	return result, nil
}
