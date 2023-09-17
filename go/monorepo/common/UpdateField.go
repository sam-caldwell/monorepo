package criCommon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/viper"
)

// UpdateField - Update a field value with given criteria
func UpdateField(section, recordName, recordCriteria, fieldName string, fieldValue any) error {
	//Lookup the section
	dataObject := viper.Get(section).([]interface{})
	// Find the matching recordName using RecordCriteria and update fieldName with fieldValue
	for i, rt := range dataObject {
		rtMap := rt.(map[string]interface{})
		if rtMap[recordName].(string) == recordCriteria {
			dataObject[i].(map[string]interface{})[fieldName] = fieldValue
			viper.Set(words.ContainerRuntimes, dataObject)
			return nil
		}
	}
	return fmt.Errorf("not found (%s)", recordCriteria)
}
