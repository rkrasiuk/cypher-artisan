package art

import (
	"fmt"
	"strings"
)

// Path ...
type Path string

const (
	// Plain --
	Plain Path = "--"

	// Outgoing -->
	Outgoing Path = "-->"

	// Incoming <--
	Incoming Path = "<--"

	// Bidirectional <-->
	Bidirectional Path = "<-->"
)

// Edge ...
type Edge struct {
	name   string
	labels []string
	props  Props
	path   Path
	// variable length
}

// NewEdge ...
func NewEdge(name string) *Edge {
	return &Edge{
		name,
		[]string{},
		make(map[string]interface{}),
		"",
	}
}

// Labels ...
func (e *Edge) Labels(labels ...string) *Edge {
	for _, label := range labels {
		e.labels = append(e.labels, label)
	}
	return e
}

// Props ...
func (e *Edge) Props(props ...Prop) *Edge {
	for _, prop := range props {
		e.props[prop.Key] = prop.Value
	}
	return e
}

// Path ...
func (e *Edge) Path(path Path) *Edge {
	e.path = path
	return e
}

// Relationship ...
func (e Edge) Relationship(lnode, rnode *Node) string {
	return fmt.Sprintf("%v%v%v", lnode.String(), e.String(), rnode.String())
}

func (e Edge) String() (res string) {
	res = e.name

	var labels string
	if len(e.labels) > 0 {
		labels = fmt.Sprintf(":%v", strings.Join(e.labels, ":"))
		res += labels
	}

	if len(e.props) > 0 {
		res += " " + e.props.String()
	}

	res = fmt.Sprintf("-[%v]-", res)

	switch e.path {
	case Outgoing:
		res += ">"
	case Incoming:
		res = "<" + res
	case Bidirectional:
		res = "<" + res + ">"
	case Plain:
	default:
	}

	return res
}
