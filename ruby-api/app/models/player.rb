class Player < ApplicationRecord

  has_many :cards, dependent: :destroy

  validates :age, presence:     true,
                 numericality: { only_integer: true, greater_than: 0 }

  validates :name, presence: true

  validates :score, presence:     true,
                   numericality: { only_integer: true, greater_than_or_equal_to: 0 }

end
