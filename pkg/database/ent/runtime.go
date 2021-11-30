// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/willie-lin/YEVER/pkg/database/ent/schema"
	"github.com/willie-lin/YEVER/pkg/database/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[1].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[3].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[6].Descriptor()
	// user.DefaultDescription holds the default value on creation for the description field.
	user.DefaultDescription = userDescDescription.Default.(string)
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = func() func(string) error {
		validators := userDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[7].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescUpdated is the schema descriptor for updated field.
	userDescUpdated := userFields[8].Descriptor()
	// user.DefaultUpdated holds the default value on creation for the updated field.
	user.DefaultUpdated = userDescUpdated.Default.(func() time.Time)
	// user.UpdateDefaultUpdated holds the default value on update for the updated field.
	user.UpdateDefaultUpdated = userDescUpdated.UpdateDefault.(func() time.Time)
}
