# Thread.new do
#     loop do
#       begin
#         UpdateChatMessagesCountJob.perform_later
#         SyncChatCountsJob.perform_later
#     rescue => e
#         Rails.logger.error("jobs failed: #{e.message}")
#       end
#       sleep 1.hour
#     end
# end
  