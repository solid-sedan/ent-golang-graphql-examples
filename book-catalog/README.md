

### transactioner handles mutations in transaction
[https://entgo.io/docs/graphql#transactional-mutations](https://entgo.io/docs/graphql#transactional-mutations)

## migrations

[https://entgo.io/docs/versioned-migrations](https://entgo.io/docs/versioned-migrations)

### create migrations
```
atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://mariadb/latest/test"
```

### apply migrations
```
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "maria://root@localhost:3306/ent_book_catalog"
```

### check migration status
```
atlas migrate status \
  --dir "file://ent/migrate/migrations" \
  --url "maria://root@localhost:3306/ent_book_catalog"
```