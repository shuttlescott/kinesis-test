module Types
  class QueryType < Types::BaseObject
    field :players, resolver: Resolvers::Players
  end
end
