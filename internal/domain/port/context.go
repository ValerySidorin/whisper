package port

import "context"

type ContextProvider interface {
	Ctx() context.Context
}
