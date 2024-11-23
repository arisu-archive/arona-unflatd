package analyzer

func (state *ParsingState) HandleEnumName(name string) {
	state.Name = name
	state.Enum.Name = name
}

func (state *ParsingState) HandleEnumBaseType(baseType string) {
	state.Enum.BaseType = baseType
}

func (state *ParsingState) HandleEnumMemberName(name string) {
	state.Enum.Keys = append(state.Enum.Keys, name)
}

func (state *ParsingState) HandleEnumMemberValue(value string) {
	state.Enum.Values = append(state.Enum.Values, value)
}
