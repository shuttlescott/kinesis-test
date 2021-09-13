class CreateCards < ActiveRecord::Migration[6.1]
  def change
    create_table :cards do |t|
      t.string :suit, null: false
      t.integer :value, null: false
      t.references :player

      t.timestamps
    end
  end
end
