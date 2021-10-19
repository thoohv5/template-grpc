// Code generated by entc, DO NOT EDIT.

package userextend

const (
	// Label holds the string label denoting the userextend type in the database.
	Label = "user_extend"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserIdentity holds the string denoting the user_identity field in the database.
	FieldUserIdentity = "user_identity"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// Table holds the table name of the userextend in the database.
	Table = "user_extend"
)

// Columns holds all SQL columns for userextend fields.
var Columns = []string{
	FieldID,
	FieldUserIdentity,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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