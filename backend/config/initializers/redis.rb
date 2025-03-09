require 'redis'

redis_url = ENV.fetch("REDIS_URL") { "redis://backend-ror-golang-redis-1:6379/1" } 
$redis = Redis.new(url: redis_url)