module Types
  class PlayerType < Types::BaseObject
    field :id,         ID,                              null: false
    field :age,        Integer,                         null: false
    field :cards,      [Types::CardType],               null: false
    field :name,       String,                          null: false
    field :score,      Integer,                         null: false
    field :created_at, GraphQL::Types::ISO8601DateTime, null: false
    field :updated_at, GraphQL::Types::ISO8601DateTime, null: false
  end
end
