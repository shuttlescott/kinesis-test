package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kinesis-test/ent"
	"time"
)

func (r *cardResolver) CreatedAt(ctx context.Context, obj *ent.Card) (*time.Time, error) {
	card, err := r.client.Card.Get(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("card not found")
	}
	return &card.CreateTime, nil
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, player PlayerInput) (*ent.Player, error) {
	return r.client.Player.
		Create().
		SetAge(player.Age).
		SetNillableName(player.Name).
		Save(ctx)
}

func (r *mutationResolver) CreateCard(ctx context.Context, card CardInput) (*ent.Card, error) {
	return r.client.Card.
		Create().
		SetPlayerID(card.PlayerID).
		SetSuit(card.Suit).
		SetValue(card.Value).
		Save(ctx)
}

func (r *mutationResolver) ClearCards(ctx context.Context, playerID int) (*ent.Player, error) {
	player, err := r.client.Player.Get(ctx, playerID)
	if err != nil {
		return nil, fmt.Errorf("player not found")
	}
	return player.Update().
		ClearCards().
		Save(ctx)
}

func (r *playerResolver) CreatedAt(ctx context.Context, obj *ent.Player) (*time.Time, error) {
	player, err := r.client.Player.Get(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("player not found")
	}
	return &player.CreateTime, nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Players(ctx context.Context, where *ent.PlayerWhereInput) ([]*ent.Player, error) {
	result, err := where.Filter(r.client.Player.Query())
	if err != nil {
		return nil, fmt.Errorf("error querying players")
	}
	return result.All(ctx)
}

// Card returns CardResolver implementation.
func (r *Resolver) Card() CardResolver { return &cardResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Player returns PlayerResolver implementation.
func (r *Resolver) Player() PlayerResolver { return &playerResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type cardResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type playerResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
