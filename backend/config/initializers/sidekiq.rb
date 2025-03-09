require 'sidekiq/cron/job'

Sidekiq.configure_server do |config|
    config.redis = { url: ENV.fetch("REDIS_URL", "redis://backend-ror-golang-redis-1:6379/1") }
end
  
Sidekiq.configure_client do |config|
    config.redis = { url: ENV.fetch("REDIS_URL", "redis://backend-ror-golang-redis-1:6379/1") }
end
  

Sidekiq::Cron::Job.create(
  name: 'SyncChatCountsJob - every hour',
  cron: '0 * * * *', # This runs every hour on the hour
  class: 'SyncChatCountsJob'
)

# Schedule UpdateChatMessagesCountJob every hour
Sidekiq::Cron::Job.create(
  name: 'UpdateChatMessagesCountJob - every hour',
  cron: '0 * * * *', # This runs every hour on the hour
  class: 'UpdateChatMessagesCountJob'
)