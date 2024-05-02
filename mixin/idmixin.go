package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/pkg/utils/ulids"
)

// IDMixin holds the schema definition for the ID
type IDMixin struct {
	mixin.Schema
	// ExcludeMappingID to exclude the mapping ID field to the schema that can be used without exposing the primary ID
	// by default, it is included in any schema that uses this mixin.
	ExcludeMappingID bool
}

// Fields of the IDMixin.
func (i IDMixin) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("id").
			Immutable().
			Annotations(entoas.Annotation{ReadOnly: true}).
			DefaultFunc(func() string { return ulids.New().String() }),
	}

	if !i.ExcludeMappingID {
		fields = append(fields,
			field.String("mapping_id").
				Immutable().
				Annotations(
					entoas.Skip(true),
					entgql.Skip(),
				).
				DefaultFunc(func() string { return ulids.New().String() }),
		)
	}

	return fields
}
