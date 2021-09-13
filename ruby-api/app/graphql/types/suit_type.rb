module Types
  class SuitType < Types::BaseEnum
    Card::SUITS.each do |suit|
      value suit
    end
  end
end
