class UpdateChatMessagesCountJob < ApplicationJob
  queue_as :default

  REDIS_KEY = "chats:pending_message_updates"

  def perform
    chat_ids = $redis.smembers(REDIS_KEY) 
    return if chat_ids.empty?

    chat_ids.each do |chat_id|
      redis_message_count = $redis.get("chat:#{chat_id}:message_count").to_i
      chat = Chat.find_by(id: chat_id)

      next unless chat && redis_message_count > 0 && redis_message_count != chat.messages_count

      chat.update_columns(messages_count: redis_message_count) 

      $redis.srem(REDIS_KEY, chat_id) 
    end
  end
end