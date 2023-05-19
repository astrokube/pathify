package pathify

import (
	"log"
)

var emptyValue = struct{}{}

type pathValue struct {
	path  string
	value any
}

type pathValueList []pathValue

type sanitizer struct {
	strict bool
}

func (s *sanitizer) sanitize(args ...any) pathValueList {
	if len(args)%2 != 0 {
		args = append(args, emptyValue)
	}
	list := make(pathValueList, len(args)/2)
	arg := 0
	invalidPathValues := 0
	for i := 0; i < len(args); i += 2 {
		path, ok := args[i].(string)
		if !ok {
			if s.strict {
				log.Panicf("invalid path '%v'.  Paths must be string", args[i])
			}
			invalidPathValues += 1
			continue
		}
		list[arg] = pathValue{
			path:  path,
			value: args[i+1],
		}
		arg++
	}
	if invalidPathValues > 0 {
		return list[:len(list)-invalidPathValues]
	}
	return list
}
