package server

func (s *Server) SetupRoutes() {
	s.Router.GET(s.Config.ConnectionRoute, s.UpgradeConnection)
}
