package analyzer

func (state *ParsingState) HandleMethodName(methodName string) {
	state.Method.Name = methodName
}

func (state *ParsingState) HandleMethodReturnType(returnType string) {
	state.Method.ReturnType = returnType
}

func (state *ParsingState) HandleParamName(paramName string) {
	state.Method.ParameterNames = append(state.Method.ParameterNames, paramName)
}

func (state *ParsingState) HandleParamType(paramType string) {
	state.Method.ParameterTypes = append(state.Method.ParameterTypes, paramType)
}
