package repository

type implAuthor struct {}

func authorRepo() implAuthor {
  impl := implAuthor{}

  return impl
}

func (a *implAuthor) Create(name string) (string, error) {
  query := `INSERT INTRO "authors" ("name") VALUES ($1) RETURNING "id"`

  var authorId string

  err := pool.QueryRow(
    ctx,
    query,
    name,
  ).Scan(&authorId)

  if err != nil {
    return "", parsePgError(err)
  }

  return authorId, nil
}
