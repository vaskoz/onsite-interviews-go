package schema

// Field represents a field that has aliases used by other fields.
type Field struct {
	name    string
	aliases []string
}

// Schema is a list of Fields.
type Schema struct {
	fields []Field
}

// Mapping takes two schemas and returns a map of field names between schemas.
func Mapping(src, target Schema) map[string]string {
	aliasMap := make(map[string]string)
	for _, field := range target.fields {
		aliasMap[field.name] = field.name
		for _, alias := range field.aliases {
			aliasMap[alias] = field.name
		}
	}
	result := make(map[string]string)
	for _, field := range src.fields {
		if _, found := aliasMap[field.name]; found {
			result[field.name] = aliasMap[field.name]
		}
	}
	return result
}
