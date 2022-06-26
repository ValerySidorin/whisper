package appctx

import "context"

var cc *CoreContext

type CoreContext struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *CoreContext) Ctx() context.Context {
	return c.ctx
}

func (c *CoreContext) Cancel() {
	c.cancel()
}

func Register() *CoreContext {
	if nil == cc {
		cc = new(CoreContext)
		cc.ctx, cc.cancel = context.WithCancel(context.Background())
	}
	return cc
}
