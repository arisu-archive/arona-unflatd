package state

import "github.com/arisu-archive/arona-unflatd/pkg/parser/analyzer"

func DetermineState(state *analyzer.ParsingState) NodeState {
	switch {
	case state.Namespace != "":
		return &NamespaceState{}
	case state.Name != "" && state.Enum.Name != "":
		return &EnumState{}
	case state.Name != "" && state.Field.Name != "":
		return &FieldState{}
	case state.Name != "" && state.Method.Name != "":
		return &MethodState{}
	case state.Name != "":
		return &StructState{}
	default:
		return &EmptyState{}
	}
}
