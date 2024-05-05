package downloader

import (
	"fmt"
	"github.com/gin-gonic/gin"
	c "github.com/gx/youtubeDownloader/constants"
)

func (s *Server) SetupRoutes(router *gin.Engine) {
	router.GET(fmt.Sprintf("%s", s.Config.ConnectionRoute), s.UpgradeConnection)
	router.GET(fmt.Sprintf("%s/:%s", s.Config.ConnectionRoute, c.SessionParameter), s.UpgradeConnection)
}
