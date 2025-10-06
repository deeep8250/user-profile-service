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





# 🧩 Migration Summary (Beginner Version)

## 🚀 What is a Migration?

A **migration** is a set of SQL files that describe **changes to your database schema** (tables, columns, indexes).  
Think of it like **version control for your database** — similar to how Git tracks code changes.

Each migration includes:

- **`.up.sql`** → applies changes (e.g., create table, add column)
- **`.down.sql`** → rolls back changes (e.g., drop table, remove column)

---

## 🎯 Why Do We Need Migrations?

- 🕒 Track database schema changes over time  
- 🔁 Reproduce your database easily on other machines  
- 🧱 Allow **safe rollbacks** if a change fails  
- 👥 Keep all developers' databases **in sync**  
- 🚀 Automate schema updates during deployments  

> Without migrations, small projects can work, but for bigger apps, migrations save you from huge headaches.

---

## ⚙️ How `migrate` Works

1. `migrate` keeps a special table in your database called **`schema_migrations`**.
2. When you run `migrate up`, it:
   - Reads all migration files
   - Checks which ones are **not yet applied**
   - Runs only those pending migrations
3. When you run `migrate down`, it:
   - Uses the corresponding `.down.sql` to **undo** the last migration

---

## 💻 Common Migration Commands

| Command | Description |
|----------|--------------|
| `migrate -path "DB/Migration" -database "<DB_URL>" up` | Apply all pending migrations |
| `migrate -path "DB/Migration" -database "<DB_URL>" down 1` | Rollback the last migration |
| `migrate -path "DB/Migration" -database "<DB_URL>" version` | Show the current migration version |
| `migrate -path "DB/Migration" -database "<DB_URL>" force <version>` | Force set migration version (advanced use) |

---

## 🗂️ Migration File Structure

Each migration file should have a version number prefix:

```
DB/Migration/
├─ 0001_create_users.up.sql
├─ 0001_create_users.down.sql
├─ 0002_add_bio_column.up.sql
└─ 0002_add_bio_column.down.sql
```

- `.up.sql` → SQL commands to **apply** the change  
- `.down.sql` → SQL commands to **rollback** the change  

> Each new schema change = new migration file.  
> You **never edit old migrations** once they’re applied.

---

## 🔐 Using `.env` with `migrate`

### Example `.env`
```env
DB_URL=postgres://postgres:deep@localhost:5432/New?sslmode=disable
```

### Load it in PowerShell
```powershell
$env:DB_URL = (Get-Content .env | Select-String 'DB_URL' | ForEach-Object { $_ -replace 'DB_URL=', '' })
```

### Run Migrations
```powershell
migrate -path "DB/Migration" -database $env:DB_URL up
```

### Rollback Last Migration
```powershell
migrate -path "DB/Migration" -database $env:DB_URL down 1
```

> This avoids hardcoding your database URL — it’s cleaner and safer.

---

## 📜 Common Outputs and What They Mean

| Output | Meaning |
|---------|----------|
| `Applying migration 0001_create_users.up.sql` | Migration applied successfully |
| `no change` | All migrations are already applied — database is up to date |
| `Error: connection refused` | Database not running or wrong connection details |
| `Failed to open source` | Wrong folder path for migrations |

---

## 🧠 Important Notes

- ✅ Always **keep migration files** — they are part of your project’s history.  
- ⚠️ **Never edit old migration files** — always create new ones for new changes.  
- 📄 `schema_migrations` is created **automatically** by `migrate` and tracks applied versions.  
- 🧩 You can run migrations from PowerShell, Docker, or in CI/CD pipelines.

---

## 🏁 TL;DR — Quick Recap

1. Migration = version-controlled database changes  
2. `.up.sql` → apply | `.down.sql` → rollback  
3. `migrate` uses `schema_migrations` to track versions  
4. Always create a new migration for schema changes  
5. Use Postgres URL directly or load it from `.env`  
6. `no change` = everything is already done ✅  

---

> 💡 **Pro Tip:** Keep a `run_migrations.ps1` file in your project to load `.env` and apply migrations automatically.  
> Example:
> ```powershell
> $env:DB_URL = (Get-Content .env | Select-String 'DB_URL' | ForEach-Object { $_ -replace 'DB_URL=', '' })
> migrate -path "DB/Migration" -database $env:DB_URL up
> ```

---

🧰 **Now you understand migrations like a pro beginner!**
