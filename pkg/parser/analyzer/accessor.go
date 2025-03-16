package analyzer

import (
	"strconv"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// HandleAccessorAttribute processes accessor attributes and stores their parameters.
// It extracts attribute name, parameters, and values from the AST node.
func (state *ParsingState) HandleAccessorAttribute(node *sitter.Node, content []byte) {
	// Get attribute name
	nameNode := node.ChildByFieldName("name")
	if nameNode == nil {
		return
	}
	attributeName := nameNode.Content(content)

	// Process attribute arguments
	processAttributeArguments(state, node, content, attributeName)
}

// processAttributeArguments finds and processes all argument lists in an attribute.
func processAttributeArguments(state *ParsingState, node *sitter.Node, content []byte, attributeName string) {
	// Find attribute_argument_list nodes
	for i := range node.NamedChildCount() {
		child := node.NamedChild(int(i))
		if child.Type() != "attribute_argument_list" {
			continue
		}

		// Process each argument in the list
		processArguments(state, child, content, attributeName)
	}
}

// processArguments processes all arguments in an attribute_argument_list.
func processArguments(state *ParsingState, argListNode *sitter.Node, content []byte, attributeName string) {
	for i := range argListNode.NamedChildCount() {
		arg := argListNode.NamedChild(int(i))
		if arg.Type() != "attribute_argument" {
			continue
		}

		// Process each assignment expression in the argument
		processAssignmentExpressions(state, arg, content, attributeName)
	}
}

// processAssignmentExpressions handles assignment expressions within an attribute argument.
func processAssignmentExpressions(state *ParsingState, argNode *sitter.Node, content []byte, attributeName string) {
	for i := range argNode.NamedChildCount() {
		assignExpr := argNode.NamedChild(int(i))
		if assignExpr.Type() != "assignment_expression" {
			continue
		}

		// Extract parameter name and value
		leftNode := assignExpr.ChildByFieldName("left")
		rightNode := assignExpr.ChildByFieldName("right")

		if leftNode == nil || rightNode == nil {
			continue
		}

		argName := leftNode.Content(content)
		argValue := rightNode.Content(content)

		// Handle string literals
		unquoted, err := strconv.Unquote(argValue)
		if err != nil {
			continue
		}

		// Initialize accessor maps if needed
		ensureAccessorMapsExist(state, attributeName)

		// Store the parameter value
		state.Field.Accessors[attributeName][argName] = strings.TrimPrefix(unquoted, "0x")
	}
}

// ensureAccessorMapsExist initializes the accessor map structure if it doesn't exist.
func ensureAccessorMapsExist(state *ParsingState, attributeName string) {
	if state.Field.Accessors == nil {
		state.Field.Accessors = make(map[string]map[string]string)
	}
	if state.Field.Accessors[attributeName] == nil {
		state.Field.Accessors[attributeName] = make(map[string]string)
	}
}
