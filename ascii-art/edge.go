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

func (e Edge) String() string {
	var labels string

	if len(e.labels) > 0 {
		labels = fmt.Sprintf(":%v", strings.Join(e.labels, ":"))
	}

	res := fmt.Sprintf("-[%v%v %v]-", e.name, labels, e.props.String())

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
