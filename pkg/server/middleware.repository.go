package server

import (
	"github.com/yznts/kyoto/v3/component"
	"github.com/yznts/player/pkg/sync"
)

func NewRepositoryMiddleware(r sync.Repository) func(c component.Component) component.Component {
	return func(c component.Component) component.Component {
		return func(ctx *component.Context) component.State {
			ctx.Set("repository", r)
			state := c(ctx)
			state.SetName(c.GetName())
			return state
		}
	}
}
