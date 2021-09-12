package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Mixin of the Card.
func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("suit").
			NamedValues(
				"Hearts", "HEARTS",
				"Diamonds", "DIAMONDS",
				"Clubs", "CLUBS",
				"Spades", "SPADES",
			),
		field.Int("value").
			Min(1).
			Max(13),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("cards").
			Unique(),
	}
}
