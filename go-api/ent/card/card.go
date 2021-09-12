// Code generated by entc, DO NOT EDIT.

package card

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSuit holds the string denoting the suit field in the database.
	FieldSuit = "suit"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgePlayer holds the string denoting the player edge name in mutations.
	EdgePlayer = "player"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// PlayerTable is the table that holds the player relation/edge.
	PlayerTable = "cards"
	// PlayerInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayerInverseTable = "players"
	// PlayerColumn is the table column denoting the player relation/edge.
	PlayerColumn = "player_cards"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSuit,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"player_cards",
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(int) error
)

// Suit defines the type for the "suit" enum field.
type Suit string

// Suit values.
const (
	SuitHearts   Suit = "HEARTS"
	SuitDiamonds Suit = "DIAMONDS"
	SuitClubs    Suit = "CLUBS"
	SuitSpades   Suit = "SPADES"
)

func (s Suit) String() string {
	return string(s)
}

// SuitValidator is a validator for the "suit" field enum values. It is called by the builders before save.
func SuitValidator(s Suit) error {
	switch s {
	case SuitHearts, SuitDiamonds, SuitClubs, SuitSpades:
		return nil
	default:
		return fmt.Errorf("card: invalid enum value for suit field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (s Suit) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (s *Suit) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*s = Suit(str)
	if err := SuitValidator(*s); err != nil {
		return fmt.Errorf("%s is not a valid Suit", str)
	}
	return nil
}