-- wrk script to benchmark https://localhost:5001/latest-posts endpoint
-- Usage: wrk -t10 -c50 -d10s -s bench.lua http://localhost:5001

request = function()
    path = "/latest-posts"
    return wrk.format("GET", path)
end
