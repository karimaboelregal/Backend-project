class ChatsController < BaseController
    def create
      application = Application.find_by!(token: params[:application_token])
      
      # Generate the next chat number using Redis
      chat_number = $redis.get("application:#{application.id}:chat_count").to_i + 1
  
      # Create the chat in MySQL
      chat = Chat.create!(application: application, number: chat_number)
      $redis.set("application:#{application.id}:chat_count", chat_number)

      # Increment a separate Redis counter to track pending updates
      $redis.sadd("pending_chat_updates", application.id)
  
      render json: { chat_number: chat.number }, status: :created
    end
  
    def show
      application = Application.find_by!(token: params[:application_token])
      chat = application.chats.find_by!(number: params[:chat_number])
  
      render json: chat
    end
  end
  