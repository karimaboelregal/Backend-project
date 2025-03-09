class SyncChatCountsJob < ApplicationJob
  queue_as :default

  def perform
    redis_key = "pending_chat_updates"
    application_ids = $redis.smembers(redis_key) 

    application_ids.each do |app_id|
      app = Application.find_by(id: app_id)
      next unless app 

      # Update MySQL chat count
      new_count = app.chats.count
      app.update!(chats_count: new_count)

      $redis.srem(redis_key, app_id)
    end
  end
end