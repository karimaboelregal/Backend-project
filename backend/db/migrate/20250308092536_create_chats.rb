class CreateChats < ActiveRecord::Migration[8.0]
  def change
    create_table :chats do |t|
      t.references :application, null: false, foreign_key: true
      t.integer :number
      t.integer :messages_count

      t.timestamps
    end
    add_index :chats, :number
  end
end
