package merchant_bot

import (
	"context"
	"sync"
	"net/http"
)

type HTTPServer struct {
	*http.Server
	handler http.HandlerFunc
	WebhookPath string
}

func NewHTTPServer(addr, webhookPath string, handler http.HandlerFunc) *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{
			Addr: addr,
		},
		handler: handler,
		WebhookPath: webhookPath,
	}
}

func (s *HTTPServer) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Infow("starting http webhook server", "addr", s.Addr)

	mux := http.NewServeMux()

	mux.HandleFunc(s.WebhookPath, s.handler)

	s.Server.Handler = mux

	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Errorw("http server returned error", "error", err)
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) {
	s.Server.Shutdown(ctx)
}
