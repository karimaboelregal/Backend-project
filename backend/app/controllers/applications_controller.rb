class ApplicationsController < BaseController
  def create
    application = Application.create!(application_params.merge(token: SecureRandom.hex(10)))
    render json: { token: application.token, name: application.name }, status: :created
  end

  def show
    application = Application.find_by!(token: params[:token])
    render json: application
  end

  private

  def application_params
    params.require(:application).permit(:name)
  end
end
