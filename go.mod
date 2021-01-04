module github.com/josephshih13/short-url

go 1.15

replace github.com/josephshih13/short-url/redis => ../redis

require (
	github.com/go-redis/redis/v8 v8.4.4
	github.com/labstack/echo/v4 v4.1.17
)
