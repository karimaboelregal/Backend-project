Rails.application.config.to_prepare do
    Thread.new do
      begin
        Message.reindex
        puts "✅ Searchkick indexing completed!"
      rescue => e
        puts "❌ Error during Searchkick reindex: #{e.message}"
      end
    end
  end
  