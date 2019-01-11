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

func (n Node) String() string {
	var labels string

	if len(n.labels) > 0 {
		labels = fmt.Sprintf(":%v", strings.Join(n.labels, ":"))
	}

	return fmt.Sprintf("(%v%v %v)", n.name, labels, n.props.String())
}
