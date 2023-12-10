package pgsqlTracker

import Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"

func CreateTableAvatar(db *Postgres.Db) error {
	_, err := db.Exec(
		Postgres.Table(
			"avatar",
			Postgres.Columns{
				Postgres.Column("id").
					Type(Postgres.Uuid).
					Property(Postgres.PrimaryKey).
					Default("gen_random_uuid()"),
				Postgres.Column("url").
					Type(Postgres.Text).
					Property(Postgres.NotNull),
			}).
			RenderSql())
	return err
}
