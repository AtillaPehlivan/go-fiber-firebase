package helper

import "encoding/json"

func StructToMapString(in interface{}) (map[string]interface{}, error) {

	var out map[string]interface{}
	byteData, _ := json.Marshal(in)
	err := json.Unmarshal(byteData, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}

type PaginateArgs struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}
