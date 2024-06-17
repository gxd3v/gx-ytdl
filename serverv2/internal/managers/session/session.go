package session

type Session struct {
	storage managers.SessionStorage
}

func New(storage managers.SessionStorage) *managers.Session {
	return &Session{
		storage: storage,
	}
}

func (y *Session) List(ctx context.Context, sort, filters map[string]database.Filter) {
}