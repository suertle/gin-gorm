package middleware

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func Sentry() []gin.HandlerFunc {
	dsn := os.Getenv("SENTRY_DSN")

	if dsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              dsn,
			Debug:            false,
			EnableTracing:    true,
			TracesSampleRate: 1.0,
			Environment:      os.Getenv("APP_ENV"),
		}); err != nil {
			panic(err)
		}
		defer sentry.Flush(2 * time.Second)

		return []gin.HandlerFunc{
			sentrygin.New(sentrygin.Options{
				Repanic: true,
			}),
		}
	}

	return []gin.HandlerFunc{}
}
