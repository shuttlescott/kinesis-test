// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kinesis-test/ent/card"
	"kinesis-test/ent/player"
	"kinesis-test/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCard   = "Card"
	TypePlayer = "Player"
)

// CardMutation represents an operation that mutates the Card nodes in the graph.
type CardMutation struct {
	config
	op            Op
	typ           string
	id            *int
	create_time   *time.Time
	update_time   *time.Time
	suit          *card.Suit
	value         *int
	addvalue      *int
	clearedFields map[string]struct{}
	player        *int
	clearedplayer bool
	done          bool
	oldValue      func(context.Context) (*Card, error)
	predicates    []predicate.Card
}

var _ ent.Mutation = (*CardMutation)(nil)

// cardOption allows management of the mutation configuration using functional options.
type cardOption func(*CardMutation)

// newCardMutation creates new mutation for the Card entity.
func newCardMutation(c config, op Op, opts ...cardOption) *CardMutation {
	m := &CardMutation{
		config:        c,
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCardID sets the ID field of the mutation.
func withCardID(id int) cardOption {
	return func(m *CardMutation) {
		var (
			err   error
			once  sync.Once
			value *Card
		)
		m.oldValue = func(ctx context.Context) (*Card, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Card.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCard sets the old Card of the mutation.
func withCard(node *Card) cardOption {
	return func(m *CardMutation) {
		m.oldValue = func(context.Context) (*Card, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreateTime sets the "create_time" field.
func (m *CardMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *CardMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *CardMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *CardMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *CardMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *CardMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetSuit sets the "suit" field.
func (m *CardMutation) SetSuit(c card.Suit) {
	m.suit = &c
}

// Suit returns the value of the "suit" field in the mutation.
func (m *CardMutation) Suit() (r card.Suit, exists bool) {
	v := m.suit
	if v == nil {
		return
	}
	return *v, true
}

// OldSuit returns the old "suit" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldSuit(ctx context.Context) (v card.Suit, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSuit is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSuit requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSuit: %w", err)
	}
	return oldValue.Suit, nil
}

// ResetSuit resets all changes to the "suit" field.
func (m *CardMutation) ResetSuit() {
	m.suit = nil
}

// SetValue sets the "value" field.
func (m *CardMutation) SetValue(i int) {
	m.value = &i
	m.addvalue = nil
}

// Value returns the value of the "value" field in the mutation.
func (m *CardMutation) Value() (r int, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old "value" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldValue(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// AddValue adds i to the "value" field.
func (m *CardMutation) AddValue(i int) {
	if m.addvalue != nil {
		*m.addvalue += i
	} else {
		m.addvalue = &i
	}
}

// AddedValue returns the value that was added to the "value" field in this mutation.
func (m *CardMutation) AddedValue() (r int, exists bool) {
	v := m.addvalue
	if v == nil {
		return
	}
	return *v, true
}

// ResetValue resets all changes to the "value" field.
func (m *CardMutation) ResetValue() {
	m.value = nil
	m.addvalue = nil
}

// SetPlayerID sets the "player" edge to the Player entity by id.
func (m *CardMutation) SetPlayerID(id int) {
	m.player = &id
}

// ClearPlayer clears the "player" edge to the Player entity.
func (m *CardMutation) ClearPlayer() {
	m.clearedplayer = true
}

// PlayerCleared reports if the "player" edge to the Player entity was cleared.
func (m *CardMutation) PlayerCleared() bool {
	return m.clearedplayer
}

// PlayerID returns the "player" edge ID in the mutation.
func (m *CardMutation) PlayerID() (id int, exists bool) {
	if m.player != nil {
		return *m.player, true
	}
	return
}

// PlayerIDs returns the "player" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// PlayerID instead. It exists only for internal usage by the builders.
func (m *CardMutation) PlayerIDs() (ids []int) {
	if id := m.player; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPlayer resets all changes to the "player" edge.
func (m *CardMutation) ResetPlayer() {
	m.player = nil
	m.clearedplayer = false
}

// Where appends a list predicates to the CardMutation builder.
func (m *CardMutation) Where(ps ...predicate.Card) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CardMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.create_time != nil {
		fields = append(fields, card.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, card.FieldUpdateTime)
	}
	if m.suit != nil {
		fields = append(fields, card.FieldSuit)
	}
	if m.value != nil {
		fields = append(fields, card.FieldValue)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case card.FieldCreateTime:
		return m.CreateTime()
	case card.FieldUpdateTime:
		return m.UpdateTime()
	case card.FieldSuit:
		return m.Suit()
	case card.FieldValue:
		return m.Value()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CardMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case card.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case card.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case card.FieldSuit:
		return m.OldSuit(ctx)
	case card.FieldValue:
		return m.OldValue(ctx)
	}
	return nil, fmt.Errorf("unknown Card field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case card.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case card.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case card.FieldSuit:
		v, ok := value.(card.Suit)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSuit(v)
		return nil
	case card.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CardMutation) AddedFields() []string {
	var fields []string
	if m.addvalue != nil {
		fields = append(fields, card.FieldValue)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CardMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case card.FieldValue:
		return m.AddedValue()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) AddField(name string, value ent.Value) error {
	switch name {
	case card.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValue(v)
		return nil
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CardMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case card.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case card.FieldSuit:
		m.ResetSuit()
		return nil
	case card.FieldValue:
		m.ResetValue()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.player != nil {
		edges = append(edges, card.EdgePlayer)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CardMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case card.EdgePlayer:
		if id := m.player; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CardMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedplayer {
		edges = append(edges, card.EdgePlayer)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	switch name {
	case card.EdgePlayer:
		return m.clearedplayer
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	switch name {
	case card.EdgePlayer:
		m.ClearPlayer()
		return nil
	}
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	switch name {
	case card.EdgePlayer:
		m.ResetPlayer()
		return nil
	}
	return fmt.Errorf("unknown Card edge %s", name)
}

// PlayerMutation represents an operation that mutates the Player nodes in the graph.
type PlayerMutation struct {
	config
	op            Op
	typ           string
	id            *int
	create_time   *time.Time
	update_time   *time.Time
	age           *int
	addage        *int
	name          *string
	score         *int
	addscore      *int
	clearedFields map[string]struct{}
	cards         map[int]struct{}
	removedcards  map[int]struct{}
	clearedcards  bool
	done          bool
	oldValue      func(context.Context) (*Player, error)
	predicates    []predicate.Player
}

var _ ent.Mutation = (*PlayerMutation)(nil)

// playerOption allows management of the mutation configuration using functional options.
type playerOption func(*PlayerMutation)

// newPlayerMutation creates new mutation for the Player entity.
func newPlayerMutation(c config, op Op, opts ...playerOption) *PlayerMutation {
	m := &PlayerMutation{
		config:        c,
		op:            op,
		typ:           TypePlayer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlayerID sets the ID field of the mutation.
func withPlayerID(id int) playerOption {
	return func(m *PlayerMutation) {
		var (
			err   error
			once  sync.Once
			value *Player
		)
		m.oldValue = func(ctx context.Context) (*Player, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Player.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlayer sets the old Player of the mutation.
func withPlayer(node *Player) playerOption {
	return func(m *PlayerMutation) {
		m.oldValue = func(context.Context) (*Player, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlayerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlayerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PlayerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreateTime sets the "create_time" field.
func (m *PlayerMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *PlayerMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *PlayerMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *PlayerMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *PlayerMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *PlayerMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetAge sets the "age" field.
func (m *PlayerMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *PlayerMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *PlayerMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *PlayerMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *PlayerMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *PlayerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PlayerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PlayerMutation) ResetName() {
	m.name = nil
}

// SetScore sets the "score" field.
func (m *PlayerMutation) SetScore(i int) {
	m.score = &i
	m.addscore = nil
}

// Score returns the value of the "score" field in the mutation.
func (m *PlayerMutation) Score() (r int, exists bool) {
	v := m.score
	if v == nil {
		return
	}
	return *v, true
}

// OldScore returns the old "score" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldScore(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldScore is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldScore requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldScore: %w", err)
	}
	return oldValue.Score, nil
}

// AddScore adds i to the "score" field.
func (m *PlayerMutation) AddScore(i int) {
	if m.addscore != nil {
		*m.addscore += i
	} else {
		m.addscore = &i
	}
}

// AddedScore returns the value that was added to the "score" field in this mutation.
func (m *PlayerMutation) AddedScore() (r int, exists bool) {
	v := m.addscore
	if v == nil {
		return
	}
	return *v, true
}

// ResetScore resets all changes to the "score" field.
func (m *PlayerMutation) ResetScore() {
	m.score = nil
	m.addscore = nil
}

// AddCardIDs adds the "cards" edge to the Card entity by ids.
func (m *PlayerMutation) AddCardIDs(ids ...int) {
	if m.cards == nil {
		m.cards = make(map[int]struct{})
	}
	for i := range ids {
		m.cards[ids[i]] = struct{}{}
	}
}

// ClearCards clears the "cards" edge to the Card entity.
func (m *PlayerMutation) ClearCards() {
	m.clearedcards = true
}

// CardsCleared reports if the "cards" edge to the Card entity was cleared.
func (m *PlayerMutation) CardsCleared() bool {
	return m.clearedcards
}

// RemoveCardIDs removes the "cards" edge to the Card entity by IDs.
func (m *PlayerMutation) RemoveCardIDs(ids ...int) {
	if m.removedcards == nil {
		m.removedcards = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.cards, ids[i])
		m.removedcards[ids[i]] = struct{}{}
	}
}

// RemovedCards returns the removed IDs of the "cards" edge to the Card entity.
func (m *PlayerMutation) RemovedCardsIDs() (ids []int) {
	for id := range m.removedcards {
		ids = append(ids, id)
	}
	return
}

// CardsIDs returns the "cards" edge IDs in the mutation.
func (m *PlayerMutation) CardsIDs() (ids []int) {
	for id := range m.cards {
		ids = append(ids, id)
	}
	return
}

// ResetCards resets all changes to the "cards" edge.
func (m *PlayerMutation) ResetCards() {
	m.cards = nil
	m.clearedcards = false
	m.removedcards = nil
}

// Where appends a list predicates to the PlayerMutation builder.
func (m *PlayerMutation) Where(ps ...predicate.Player) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PlayerMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Player).
func (m *PlayerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PlayerMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.create_time != nil {
		fields = append(fields, player.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, player.FieldUpdateTime)
	}
	if m.age != nil {
		fields = append(fields, player.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, player.FieldName)
	}
	if m.score != nil {
		fields = append(fields, player.FieldScore)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PlayerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case player.FieldCreateTime:
		return m.CreateTime()
	case player.FieldUpdateTime:
		return m.UpdateTime()
	case player.FieldAge:
		return m.Age()
	case player.FieldName:
		return m.Name()
	case player.FieldScore:
		return m.Score()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PlayerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case player.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case player.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case player.FieldAge:
		return m.OldAge(ctx)
	case player.FieldName:
		return m.OldName(ctx)
	case player.FieldScore:
		return m.OldScore(ctx)
	}
	return nil, fmt.Errorf("unknown Player field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case player.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case player.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case player.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case player.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case player.FieldScore:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetScore(v)
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PlayerMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, player.FieldAge)
	}
	if m.addscore != nil {
		fields = append(fields, player.FieldScore)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PlayerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case player.FieldAge:
		return m.AddedAge()
	case player.FieldScore:
		return m.AddedScore()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case player.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	case player.FieldScore:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddScore(v)
		return nil
	}
	return fmt.Errorf("unknown Player numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PlayerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PlayerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlayerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Player nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PlayerMutation) ResetField(name string) error {
	switch name {
	case player.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case player.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case player.FieldAge:
		m.ResetAge()
		return nil
	case player.FieldName:
		m.ResetName()
		return nil
	case player.FieldScore:
		m.ResetScore()
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PlayerMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cards != nil {
		edges = append(edges, player.EdgeCards)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PlayerMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case player.EdgeCards:
		ids := make([]ent.Value, 0, len(m.cards))
		for id := range m.cards {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PlayerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcards != nil {
		edges = append(edges, player.EdgeCards)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PlayerMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case player.EdgeCards:
		ids := make([]ent.Value, 0, len(m.removedcards))
		for id := range m.removedcards {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PlayerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcards {
		edges = append(edges, player.EdgeCards)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PlayerMutation) EdgeCleared(name string) bool {
	switch name {
	case player.EdgeCards:
		return m.clearedcards
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PlayerMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Player unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PlayerMutation) ResetEdge(name string) error {
	switch name {
	case player.EdgeCards:
		m.ResetCards()
		return nil
	}
	return fmt.Errorf("unknown Player edge %s", name)
}