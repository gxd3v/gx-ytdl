package managers

type Session interface {
	List(ctx context.Context, sort, filters map[string]database.Filter)
}

type SessionStorage storage.Generic[models.Session]