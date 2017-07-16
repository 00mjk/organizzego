package transformation

import (
	"encoding/json"
	"fmt"
)

type JsonTransformation struct {
}

func (jsonTransformation JsonTransformation) MarshalToJSON(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}
