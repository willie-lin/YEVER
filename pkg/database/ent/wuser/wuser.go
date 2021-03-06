// Code generated by entc, DO NOT EDIT.

package wuser

const (
	// Label holds the string label denoting the wuser type in the database.
	Label = "wuser"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// Table holds the table name of the wuser in the database.
	Table = "wuser"
)

// Columns holds all SQL columns for wuser fields.
var Columns = []string{
	FieldID,
	FieldPhone,
	FieldUID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
