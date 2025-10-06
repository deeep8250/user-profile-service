Migration

First we installed migration tool globally and set the bin path into enviroment variable
   ( go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest )

For applying the migration
  First we need to create a folder inside the db called migration and there we save our migration files
  Then we need to use this
    Use migrate.exe to generate a new migration file:
    
    paste `migrate create -ext sql -dir db/migrations create_users_table`

 meaning :
    -ext sql → file extension

    -dir db/migrations → folder for your migration files

    create_users_table → descriptive name for this migration    