package pathify

type command struct {
	value    any
	children map[string]command
}

func (cmd *command) hasChildren() bool {
	return len(cmd.children) > 0
}

func (n command) append(p path, value any) {
	key := p.Parent()
	if p.isRoot() {
		n.children[key] = command{value: value}
		return
	}
	if len(n.children) == 0 {
		n.children = map[string]command{
			key: {children: make(map[string]command)},
		}
	}
	if len(n.children[key].children) == 0 {
		n.children[key] = command{children: make(map[string]command)}
	}

	n.children[key].append(p.children(), value)
}

func (cmd *command) Get() any {
	if cmd.hasChildren() {
		out := make(map[string]any)
		for k, v := range cmd.children {
			out[k] = v.Get()
		}
		return out
	}
	return cmd.value
}
