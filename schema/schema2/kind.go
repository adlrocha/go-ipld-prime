package schema

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

// TypeKind is an enum of kind in the IPLD Schema system.
//
// Note that schema.TypeKind is distinct from datamodel.Kind!
// Schema kinds include concepts such as "struct" and "enum", which are
// concepts only introduced by the Schema layer, and not present in the
// Data Model layer.
type TypeKind uint8

const (
	TypeKind_Invalid TypeKind = 0
	TypeKind_Map     TypeKind = '{'
	TypeKind_List    TypeKind = '['
	TypeKind_Unit    TypeKind = '1'
	TypeKind_Bool    TypeKind = 'b'
	TypeKind_Int     TypeKind = 'i'
	TypeKind_Float   TypeKind = 'f'
	TypeKind_String  TypeKind = 's'
	TypeKind_Bytes   TypeKind = 'x'
	TypeKind_Link    TypeKind = '/'
	TypeKind_Struct  TypeKind = '$'
	TypeKind_Union   TypeKind = '^'
	TypeKind_Enum    TypeKind = '%'
	// FUTURE: TypeKind_Any = '?'?
)

func (k TypeKind) String() string {
	switch k {
	case TypeKind_Invalid:
		return "Invalid"
	case TypeKind_Map:
		return "Map"
	case TypeKind_List:
		return "List"
	case TypeKind_Unit:
		return "Unit"
	case TypeKind_Bool:
		return "Bool"
	case TypeKind_Int:
		return "Int"
	case TypeKind_Float:
		return "Float"
	case TypeKind_String:
		return "String"
	case TypeKind_Bytes:
		return "Bytes"
	case TypeKind_Link:
		return "Link"
	case TypeKind_Struct:
		return "Struct"
	case TypeKind_Union:
		return "Union"
	case TypeKind_Enum:
		return "Enum"
	default:
		panic("invalid enumeration value!")
	}
}

// ActsLike returns a constant from the datamodel.Kind enum describing what
// this schema.TypeKind acts like at the Data Model layer.
//
// Things with similar names are generally conserved
// (e.g. "map" acts like "map");
// concepts added by the schema layer have to be mapped onto something
// (e.g. "struct" acts like "map").
//
// Note that this mapping describes how a typed Node will *act*, programmatically;
// it does not necessarily describe how it will be *serialized*
// (for example, a struct will always act like a map, even if it has a tuple
// representation strategy and thus becomes a list when serialized).
func (k TypeKind) ActsLike() datamodel.Kind {
	switch k {
	case TypeKind_Invalid:
		return datamodel.Kind_Invalid
	case TypeKind_Map:
		return datamodel.Kind_Map
	case TypeKind_List:
		return datamodel.Kind_List
	case TypeKind_Unit:
		return datamodel.Kind_Bool // maps to 'true'.
	case TypeKind_Bool:
		return datamodel.Kind_Bool
	case TypeKind_Int:
		return datamodel.Kind_Int
	case TypeKind_Float:
		return datamodel.Kind_Float
	case TypeKind_String:
		return datamodel.Kind_String
	case TypeKind_Bytes:
		return datamodel.Kind_Bytes
	case TypeKind_Link:
		return datamodel.Kind_Link
	case TypeKind_Struct:
		return datamodel.Kind_Map // clear enough: fields are keys.
	case TypeKind_Union:
		return datamodel.Kind_Map // REVIEW: unions are tricky.
	case TypeKind_Enum:
		return datamodel.Kind_String // 'AsString' is the one clear thing to define.
	default:
		panic("invalid enumeration value!")
	}
}
