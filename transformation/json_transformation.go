package transformation

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type JsonTransformation struct {
}

func (jsonTransformation JsonTransformation) MarshalToJSON(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).Error("Failed to serialize structure to JSON")
	}
	return bytes
}

func (jsonTransformation JsonTransformation) UnmarshalFromJSON(jsonBytes []byte, data interface{}) {
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.WithError(err).Error("Failed to deserialize structure from JSON")
	}
}
