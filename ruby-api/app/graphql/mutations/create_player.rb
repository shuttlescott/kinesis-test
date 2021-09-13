module Mutations
  class CreatePlayer < Mutations::BaseMutation
    argument :player, Types::PlayerInput, required: true

    field :player, Types::PlayerType, null: true

    def resolve(player:)
      record = Player.find_or_create_by!(**player) || nil

      { player: record}
    end
  end
end
