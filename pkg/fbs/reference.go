package fbs

type SchemaReference map[string][]*Schema

func (s SchemaReference) HasNamespace(schemaName, namespace string) bool {
	// Check if the path is in the reference map
	references, ok := s[schemaName]
	if !ok {
		return false
	}

	// Check if the namespace is in the reference map
	for _, schema := range references {
		if schema.Namespace == namespace {
			return true
		}
	}
	return false
}
