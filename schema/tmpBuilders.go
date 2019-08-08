package schema

// Everything in this file is __a temporary hack__ and will be __removed__.
//
// These methods will only hang around until more of the "ast" packages are finished;
// thereafter, building schema.Type and schema.TypeSystem values will only be
// possible through first constructing a schema AST, and *then* using Reify(),
// which will validate things correctly, cycle-check, cross-link, etc.
//
// (Meanwhile, we're using these methods in the codegen prototypes.)

func SpawnString(name TypeName) TypeString {
	return TypeString{anyType{name, nil}}
}

func SpawnStruct(name TypeName, fields []StructField, repr StructRepresentation) TypeStruct {
	return TypeStruct{anyType{name, nil}, fields, repr}
}
func SpawnStructField(name string, typ Type, optional bool, nullable bool) StructField {
	return StructField{name, typ, optional, nullable}
}