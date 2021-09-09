// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"creative-production/ent/card"
	"creative-production/ent/player"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerCreate is the builder for creating a Player entity.
type PlayerCreate struct {
	config
	mutation *PlayerMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PlayerCreate) SetCreateTime(t time.Time) *PlayerCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableCreateTime(t *time.Time) *PlayerCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PlayerCreate) SetUpdateTime(t time.Time) *PlayerCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableUpdateTime(t *time.Time) *PlayerCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetAge sets the "age" field.
func (pc *PlayerCreate) SetAge(i int) *PlayerCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetName sets the "name" field.
func (pc *PlayerCreate) SetName(s string) *PlayerCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableName(s *string) *PlayerCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetScore sets the "score" field.
func (pc *PlayerCreate) SetScore(i int) *PlayerCreate {
	pc.mutation.SetScore(i)
	return pc
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableScore(i *int) *PlayerCreate {
	if i != nil {
		pc.SetScore(*i)
	}
	return pc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (pc *PlayerCreate) AddCardIDs(ids ...int) *PlayerCreate {
	pc.mutation.AddCardIDs(ids...)
	return pc
}

// AddCards adds the "cards" edges to the Card entity.
func (pc *PlayerCreate) AddCards(c ...*Card) *PlayerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCardIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pc *PlayerCreate) Mutation() *PlayerMutation {
	return pc.mutation
}

// Save creates the Player in the database.
func (pc *PlayerCreate) Save(ctx context.Context) (*Player, error) {
	var (
		err  error
		node *Player
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlayerCreate) SaveX(ctx context.Context) *Player {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlayerCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlayerCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlayerCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := player.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := player.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := player.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Score(); !ok {
		v := player.DefaultScore
		pc.mutation.SetScore(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlayerCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "age"`)}
	}
	if v, ok := pc.mutation.Age(); ok {
		if err := player.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "age": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := pc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "score"`)}
	}
	if v, ok := pc.mutation.Score(); ok {
		if err := player.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "score": %w`, err)}
		}
	}
	return nil
}

func (pc *PlayerCreate) sqlSave(ctx context.Context) (*Player, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PlayerCreate) createSpec() (*Player, *sqlgraph.CreateSpec) {
	var (
		_node = &Player{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: player.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: player.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: player.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: player.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Score(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldScore,
		})
		_node.Score = value
	}
	if nodes := pc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.CardsTable,
			Columns: []string{player.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerCreateBulk is the builder for creating many Player entities in bulk.
type PlayerCreateBulk struct {
	config
	builders []*PlayerCreate
}

// Save creates the Player entities in the database.
func (pcb *PlayerCreateBulk) Save(ctx context.Context) ([]*Player, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Player, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlayerCreateBulk) SaveX(ctx context.Context) []*Player {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlayerCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlayerCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}