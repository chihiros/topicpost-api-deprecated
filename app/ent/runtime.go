// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/profile"
	"app/ent/recreation"
	"app/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileFields[3].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileFields[4].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	recreationFields := schema.Recreation{}.Fields()
	_ = recreationFields
	// recreationDescPublish is the schema descriptor for publish field.
	recreationDescPublish := recreationFields[8].Descriptor()
	// recreation.DefaultPublish holds the default value on creation for the publish field.
	recreation.DefaultPublish = recreationDescPublish.Default.(bool)
	// recreationDescCreatedAt is the schema descriptor for created_at field.
	recreationDescCreatedAt := recreationFields[9].Descriptor()
	// recreation.DefaultCreatedAt holds the default value on creation for the created_at field.
	recreation.DefaultCreatedAt = recreationDescCreatedAt.Default.(func() time.Time)
	// recreationDescUpdatedAt is the schema descriptor for updated_at field.
	recreationDescUpdatedAt := recreationFields[10].Descriptor()
	// recreation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	recreation.DefaultUpdatedAt = recreationDescUpdatedAt.Default.(func() time.Time)
	// recreation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	recreation.UpdateDefaultUpdatedAt = recreationDescUpdatedAt.UpdateDefault.(func() time.Time)
}
