package server

import (
	"github.com/yznts/kyoto/v3/component"
	"github.com/yznts/kyoto/v3/rendering"
	"github.com/yznts/player/pkg/sync"
	"github.com/yznts/zen/v3/conv"
	"github.com/yznts/zen/v3/httpx"
)

type PlayerPageState struct {
	component.Disposable
	rendering.Template

	ID  string
	Src string
	Sec int
}

func PlayerPage(ctx *component.Context) component.State {
	// Create state
	state := &PlayerPageState{}

	// Get repository
	repo := ctx.Get("sync").(sync.Repository)
	// Set player parameters if provided
	query := ctx.Request.URL.Query()
	if src := query.Get("src"); src != "" {
		repo.SetSrc(state.ID, src)
	}
	if sec := query.Get("sec"); sec != "" {
		repo.SetSec(state.ID, conv.Int(sec))
	}
	// Pass player ID with parameters
	state.ID = httpx.Path(ctx.Request.URL.Path).Get(0)
	state.Src, _ = repo.GetSrc(state.ID)
	state.Sec, _ = repo.GetSec(state.ID)
	// Return state
	return state
}
