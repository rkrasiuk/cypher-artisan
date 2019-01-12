package art

import (
	"fmt"
	"strings"
)

// Node ...
type Node struct {
	name   string
	labels []string
	props  Props
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

func (n Node) String() (res string) {
	res += n.name

	if len(n.labels) > 0 {
		var labels string
		labels = fmt.Sprintf(":%v", strings.Join(n.labels, ":"))
		res += labels
	}

	if len(n.props) > 0 {
		res += " " + n.props.String()
	}

	res = fmt.Sprintf("(%v)", res)
	return
}
