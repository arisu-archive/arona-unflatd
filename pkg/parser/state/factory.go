package state

import "github.com/arisu-archive/arona-unflatd/pkg/parser/analyzer"

func DetermineState(state *analyzer.ParsingState) NodeState {
	switch {
	case state.Namespace != "":
		return &NamespaceState{}
	case state.Enum.Name != "":
		return &EnumState{}
	case state.Field.Name != "":
		return &FieldState{}
	case state.Method.Name != "":
		return &MethodState{}
	case state.Name != "" && len(state.StructBaseList) > 0:
		return &StructState{}
	default:
		return &EmptyState{}
	}
}
