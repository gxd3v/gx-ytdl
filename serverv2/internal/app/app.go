package app

type App struct {
	YtdlManager managers.Ytdl
}

func New(ytdlManager managers.Ytdl) *App {
	return &App{
		YtdlManager: ytdlManager
	}
}