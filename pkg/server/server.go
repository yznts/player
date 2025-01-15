package server

import (
	"net/http"

	"github.com/yznts/kyoto/v3/rendering"
	"github.com/yznts/player/pkg/sync"
)

type Options struct {
	TemplateGlob string
	Repository   sync.Repository
}

func NewServer(o Options) *http.ServeMux {
	// Configure rendering
	rendering.TEMPLATE_GLOB = o.TemplateGlob

	// Configure repository middleware
	rmd := NewRepositoryMiddleware(o.Repository)

	// Configure mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", rendering.Handler(rmd(PlayerPage)))

	// Return
	return mux
}
