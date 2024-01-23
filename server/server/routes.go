package server

import (
	"fmt"
	c "github.com/gx/youtubeDownloader/constants"
)

func (s *Server) SetupRoutes() {
	s.Router.GET(fmt.Sprintf("%s", s.Config.ConnectionRoute), s.UpgradeConnection)
	s.Router.GET(fmt.Sprintf("%s/:%s", s.Config.ConnectionRoute, c.SESSION_PARAMETER), s.UpgradeConnection)
}
