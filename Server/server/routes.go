package server

import c "github.com/gx/youtubeDownloader/constants"

type Route struct {
	Path       string `json:"path"`
	Method     string `json:"method"`
	Controller func() `json:"-"`
	RateLimit  int32  `json:"rateLimit"`
}

func (s *Server) SetupRoutes() {
	r := &Resource{
		Ws:        nil,
		SessionID: "",
	}

	s.Router.GET(c.ROUTE_CONNECT, r.UpgradeConnection)
}
