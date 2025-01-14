package schema

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	schemadmt "github.com/ipld/go-ipld-prime/schema/dmt"
)

type TypeList struct {
	name TypeName
	dmt  schemadmt.TypeList
	ts   *TypeSystem
}

// -- schema.Type interface satisfaction -->

var _ Type = (*TypeList)(nil)

func (t *TypeList) _Type() {}

func (t *TypeList) TypeSystem() *TypeSystem {
	return t.ts
}

func (TypeList) TypeKind() TypeKind {
	return TypeKind_Map
}

func (t *TypeList) Name() TypeName {
	return t.name
}

func (t TypeList) RepresentationBehavior() datamodel.Kind {
	return datamodel.Kind_Map
}

// -- specific to TypeList -->

// ValueType returns the Type of the map values.
func (t *TypeList) ValueType() Type {
	return t.ts.types[t.dmt.FieldValueType().TypeReference()]
}

// ValueIsNullable returns a bool describing if the map values are permitted
// to be null.
func (t *TypeList) ValueIsNullable() bool {
	return t.dmt.FieldValueNullable().Bool()
}
