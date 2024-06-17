package downloader

var _ downloader = (*Downloader)(nil)

func NewDownloader() *Downloader {
	return new(Downloader)
}
