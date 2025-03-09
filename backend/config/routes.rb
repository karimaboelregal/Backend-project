require 'sidekiq/web'

Rails.application.routes.draw do
  mount Sidekiq::Web => "/sidekiq"

  resources :applications, only: [:create, :show], param: :token do
    resources :chats, only: [:create, :show], param: :chat_number do
      resources :messages, only: [:create]
      post "messages/search", to: "messages#search"
    end
  end
end