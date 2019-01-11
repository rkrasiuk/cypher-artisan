package art

import (
	"fmt"
	"strings"
)

// Prop ...
type Prop struct {
	Key   string
	Value interface{}
}

// Props ...
type Props map[string]interface{}

func (p Props) String() string {
	var props string

	if len(p) > 0 {
		var propsArr []string
		for key, prop := range p {
			switch prop.(type) {
			case string:
				propsArr = append(propsArr, fmt.Sprintf("%v: '%v'", key, prop))
			default:
				propsArr = append(propsArr, fmt.Sprintf("%v: %v", key, prop))
			}
		}
		props = fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))
	}

	return props
}
