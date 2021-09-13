module Types
  class CardType < Types::BaseObject
    field :id,         ID,                              null: false
    field :suit,       Types::SuitType,                 null: false
    field :value,      Integer,                         null: false
    field :created_at, GraphQL::Types::ISO8601DateTime, null: false
    field :updated_at, GraphQL::Types::ISO8601DateTime, null: false
  end
end
