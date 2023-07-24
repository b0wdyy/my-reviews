package initializers

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func init() {
	InitStorage()
}

func InitSession() {
	Store = session.New(session.Config{
		Storage:   RedisStorage,
		KeyLookup: "cookie:my_reviews_session",
	})
}
