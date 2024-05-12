package JiraTransition

import "encoding/json"

func (action *TransitionAction) Marshall() []byte {

	result, err := json.Marshal(action)

	if err != nil {
		panic("internal error")
	}

	return result
}
