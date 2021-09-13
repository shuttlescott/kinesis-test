class Card < ApplicationRecord

  SUITS = [
    SUIT_HEARTS   = 'Hearts',
    SUIT_DIAMONDS = 'Diamonds',
    SUIT_CLUBS    = 'Clubs',
    SUIT_SPADES   = 'Spades',
  ].freeze

  belongs_to :player

  validates :suit, presence: true, inclusion: SUITS

  validates :value, presence:     true,
                   numericality: { only_integer: true, greater_than_or_equal_to: 1, less_than_or_equal_to: 13 }

end
