package ytdl

type Ytdl struct {
	storage managers.YtdlStorage
}

func New(storage managers.YtdlStorage) *managers.Ytdl {
	return &Ytdl{
		storage: storage,
	}
}

func (y *Ytdl) List(ctx context.Context, sort, filters map[string]database.Filter) {
}