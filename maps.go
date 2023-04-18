package pathify

type Processor interface {
	Set(path string, value any) Processor
	Map() map[string]any
}

type processor struct {
	errorMsgs []string
	root      command
}

func (p *processor) HashErrors() bool {
	return len(p.errorMsgs) > 0
}

func New() Processor {
	return &processor{
		root:      command{children: make(map[string]command)},
		errorMsgs: make([]string, 0),
	}
}

func (p *processor) Set(expr string, value any) Processor {
	path := newPath(expr)
	if path.validate() {
		p.root.append(path, value)
	} else {
		p.errorMsgs = append(p.errorMsgs, "the provided path '%s' is not valid", expr)
	}
	return p
}

func (p *processor) Map() map[string]any {
	return p.root.Get().(map[string]any)
}
