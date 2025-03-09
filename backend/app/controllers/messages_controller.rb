class MessagesController < BaseController
    def create
        allowed_senders = ["client", "application"]
        unless allowed_senders.include?(params[:sender])
          return render json: { error: "Invalid sender. Must be 'client' or 'application'." }, status: :unprocessable_entity
        end
        chat = Chat.joins(:application).find_by!(
          applications: { token: params[:application_token] },
          number: params[:chat_chat_number]
        )

        ProcessMessageJob.perform_later(chat.id, params[:body], params[:sender])
      
        render json: { status: "Message queued", chat_id: chat.id }, status: :accepted
      end
      
    def search
        chat = Chat.joins(:application).find_by!(
            applications: { token: params[:application_token] },
            number: params[:chat_chat_number]
        )
    
        query = params[:query]

        results = Message.search(params[:query], where: { chat_id: chat.id })
        render json: results.map { |m| { message_number: m.number, body: m.body } }

    end
end
  