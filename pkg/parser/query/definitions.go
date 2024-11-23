package query

const StructParsingQuery = `
	; Namespace
	(file_scoped_namespace_declaration
		(identifier) @namespace)
	; Enum
	(enum_declaration
		(modifier)* @enum_modifier
		name: (identifier) @enum_name
		(base_list
			(predefined_type) @enum_base_type)?
		body: (enum_member_declaration_list
			(enum_member_declaration
				name: (identifier) @enum_member_name
				value: (_) @enum_member_value)))
	; Struct
    (struct_declaration
        name: (identifier) @struct_name
        (base_list
            (identifier) @interface)
        (declaration_list
            [
                (property_declaration
                    (modifier)* @modifier
                    type: (_) @field_type
                    name: (identifier) @field_name)
                (method_declaration
                    (modifier)* @method_modifier
                    returns: (_) @method_return_type
                    name: (identifier) @method_name
                    (parameter_list
                        (parameter
                            type: (_) @param_type
                            name: (identifier) @param_name)))
            ]))
`
