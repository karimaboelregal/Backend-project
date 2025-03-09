class Chat < ApplicationRecord
  belongs_to :application
  has_many :messages, dependent: :destroy

  before_create :assign_chat_number

  private

  def assign_chat_number
    last_number = Chat.where(application_id: application_id).maximum(:number) || 0
    self.number = last_number + 1
  end
end
