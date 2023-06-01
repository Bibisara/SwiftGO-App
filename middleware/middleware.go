package middleware

import (
	"myapp/data"

	"github.com/bibisara/swiftgo"
)

type Middleware struct {
	App    *swiftgo.SwiftGO
	Models data.Models
}
