// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/datumforge/entx/vanilla/_example/ent/organization"
	"github.com/datumforge/entx/vanilla/_example/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orgmembershipFields := schema.OrgMembership{}.Fields()
	_ = orgmembershipFields
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinHooks0 := organizationMixin[0].Hooks()
	organization.Hooks[0] = organizationMixinHooks0[0]
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationMixinFields0[0].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescUpdatedAt is the schema descriptor for updated_at field.
	organizationDescUpdatedAt := organizationMixinFields0[1].Descriptor()
	// organization.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organization.DefaultUpdatedAt = organizationDescUpdatedAt.Default.(func() time.Time)
	// organization.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organization.UpdateDefaultUpdatedAt = organizationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[1].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)