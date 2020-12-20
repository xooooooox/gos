package fast

import (
	"encoding/json"
)

func JsonMarshal(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func JsonUnmarshal(s string,v interface{}) {
	_ = json.Unmarshal([]byte(s),&v)
}
