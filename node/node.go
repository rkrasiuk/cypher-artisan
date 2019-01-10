package node

import (
	"fmt"
	"strings"
)

// Node ...
type Node struct {
	name   string
	labels []string
	props  map[string]interface{}
}

// Prop ...
type Prop struct {
	Key   string
	Value interface{}
}

// NewNode ...
func NewNode(name string) *Node {
	return &Node{
		name,
		[]string{},
		make(map[string]interface{}),
	}
}

// Labels ...
func (n *Node) Labels(labels ...string) *Node {
	for _, label := range labels {
		n.labels = append(n.labels, label)
	}
	return n
}

// Props ...
func (n *Node) Props(props ...Prop) *Node {
	for _, prop := range props {
		n.props[prop.Key] = prop.Value
	}
	return n
}

// ToString ...
func (n Node) String() string {
	var labels, props string

	if len(n.labels) > 0 {
		labels = fmt.Sprintf(":%v", strings.Join(n.labels, ":"))
	}

	if len(n.props) > 0 {
		var propsArr []string
		for key, prop := range n.props {
			switch prop.(type) {
			case string:
				propsArr = append(propsArr, fmt.Sprintf("%v: '%v'", key, prop))
			default:
				propsArr = append(propsArr, fmt.Sprintf("%v: %v", key, prop))
			}
		}
		props = fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))
	}

	return fmt.Sprintf("(%v%v %v)", n.name, labels, props)
}
