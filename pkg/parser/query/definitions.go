package query

const StructParsingQuery = `
; Group1: Namespace
(
  (file_scoped_namespace_declaration
    name: [(identifier) (qualified_name)] @namespace)
)

; Group2: Enum
(
	(enum_declaration
	(modifier)* @enum_modifier
	name: (identifier) @enum_name
	(base_list
		(predefined_type) @enum_base_type)?
	body: (enum_member_declaration_list
		(enum_member_declaration
			name: (identifier) @enum_member_name
			value: (_) @enum_member_value))) @enum_declaration
)

; Group3: Struct base info
(struct_declaration
	(modifier)* @modifier
	name: (identifier) @struct_name
	(base_list
		(identifier) @interface) @struct_base_list
)

; Group4: Struct methods
(method_declaration
	(modifier)* @method_modifier
	returns: (_) @method_return_type
	name: (identifier) @method_name
	(parameter_list
		(parameter
			type: (_) @param_type
			name: (identifier) @param_name))
)
`
