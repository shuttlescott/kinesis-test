class CreatePlayers < ActiveRecord::Migration[6.1]
  def change
    create_table :players do |t|
      t.integer :age, null: false
      t.string :name, null: false
      t.integer :score, default: 0, null: false

      t.timestamps
    end
  end
end
