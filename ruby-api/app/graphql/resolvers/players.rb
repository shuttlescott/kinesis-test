module Resolvers
  class Players < Resolvers::BaseResolver
    type [Types::PlayerType], null: false

    def resolve
      Player.all
    end
  end
end
