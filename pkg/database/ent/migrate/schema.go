// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 120},
		{Name: "age", Type: field.TypeInt},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Size: 120},
		{Name: "description", Type: field.TypeString, Size: 200, Default: ""},
		{Name: "created", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ImagesTable,
		UsersTable,
	}
)

func init() {
	ImagesTable.Annotation = &entsql.Annotation{
		Table: "images",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
}
