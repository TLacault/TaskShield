package label

type Priority string

const (
	MANDATORY Priority = "Mandatory"
	URGENT    Priority = "Urgent"
	ASAP      Priority = "ASAP"
	OPTIONAL  Priority = "Optional"
)

type Label struct {
	name     string
	priority Priority
}

/* ********************************************************************************************* */

func New(name string, priority Priority) Label {
	label := Label{
		name:     name,
		priority: priority,
	}
	return label
}

/* ********************************************************************************************* */

func (l *Label) GetName() string     { return l.name }
func (l *Label) GetPriority() string { return string(l.priority) }

/* ********************************************************************************************* */

func (l *Label) ToString() string {
	return l.name + ";" + string(l.priority)
}

func (l *Label) ToStringFormat() string {
	return "[" + l.name + "] " + string(l.priority)
}
