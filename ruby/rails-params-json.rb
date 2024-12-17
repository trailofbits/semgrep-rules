class ProductsController < ApplicationController
  def create
    # ruleid: rails-params-json
    id1 = params[:_json][:id]

    # ruleid: rails-params-json
    id2 = params["_json"]["id"]

    # ruleid: rails-params-json
    id3 = params['_json']['id']

    # ok: rails-params-json
    id4 = params[:something][:id]

    # ruleid: rails-params-json
    id5 = params.fetch(:_json)

    # ruleid: rails-params-json
    id6 = params.fetch(:_json, {})

    # ruleid: rails-params-json
    product_params = params.require(:_json).map do |product|
      product.permit(:name, :price)
    end
  end
end
