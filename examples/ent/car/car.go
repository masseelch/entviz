// Code generated by entc, DO NOT EDIT.

package car

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldModelYear holds the string denoting the model_year field in the database.
	FieldModelYear = "model_year"
	// Table holds the table name of the car in the database.
	Table = "cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldNickname,
	FieldBrand,
	FieldModelYear,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cars"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_cars",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
