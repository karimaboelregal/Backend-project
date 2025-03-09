class ProcessMessageJob < ApplicationJob
  queue_as :default

  def perform(chat_id, body, sender)
    chat = Chat.find(chat_id)
    message_number = $redis.get("chat:#{chat.id}:message_count").to_i + 1
      

    message = Message.create!(
      chat: chat,
      number: message_number,
      body: body,
      sender: sender
    )
    $redis.sadd("chats:pending_message_updates", chat.id)
    $redis.set("chat:#{chat.id}:message_count", message_number)

    # Log success
    Rails.logger.info("Message #{message.number} created successfully in chat #{chat_id}")
  rescue => e
    Rails.logger.error("Failed to process message: #{e.message}")
  end
end