class CreateApplications < ActiveRecord::Migration[8.0]
  def change
    create_table :applications do |t|
      t.string :token
      t.string :name
      t.integer :chats_count

      t.timestamps
    end
    add_index :applications, :token
  end
end
