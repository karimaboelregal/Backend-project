class Message < ApplicationRecord
  searchkick word_start: [:body, :sender]

  belongs_to :chat

  before_create :assign_message_number

  validates :sender, presence: true
  validates :sender, inclusion: { in: %w[client application] }

  def search_data
    { body: body, sender: sender, chat_id: chat_id }
  end

  private

  def assign_message_number
    last_number = Message.where(chat_id: chat_id).maximum(:number) || 0
    self.number = last_number + 1
  end
end
