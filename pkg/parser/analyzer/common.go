package analyzer

func (state *ParsingState) HandleModifier(modifier string) {
	state.Modifiers = append(state.Modifiers, modifier)
}

func (state *ParsingState) HandleName(name string) {
	state.Name = name
}

func (state *ParsingState) HandleNamespace(namespace string) {
	state.Namespace = namespace
}

func (state *ParsingState) HandleInterface(name string) {
	state.StructBaseList = name
}
