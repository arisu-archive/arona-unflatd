package analyzer

func (state *ParsingState) HandleFieldType(fieldType string) {
	state.Field.Type = fieldType
}

func (state *ParsingState) HandleFieldName(fieldName string) {
	state.Field.Name = fieldName
}
