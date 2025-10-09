# SQLC - Beginner's Guide

## 1️⃣ What is SQLC?

**SQLC** is a tool that **generates Go code from your SQL queries**.  
Instead of manually writing Go code to execute SQL, scan rows, and handle types, `sqlc` does it automatically.  

- You write **raw SQL** in `.sql` files.
- `sqlc` reads your **schema** and **queries**.
- It generates **type-safe Go structs and functions**.

---

## 2️⃣ Why use SQLC?

| Problem Without SQLC | Solution With SQLC |
|--------------------|------------------|
| Manually scanning rows in Go | Auto-generated functions return Go structs |
| Risk of runtime errors due to wrong types | Compile-time type-safe code |
| Mixing SQL strings with Go logic | Queries stay in `.sql` files, Go stays clean |

---

## 3️⃣ Folder Structure

```
DB/
├── Migration/   # Migration files (.up.sql, .down.sql)
├── Schema/      # Table definitions for sqlc
└── Query/       # SQL queries for sqlc
internal/
└── db/          # Auto-generated Go code
```

---

## 4️⃣ How to Write SQLC Queries

### Example: Single Row Query
```sql
-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
```

- `$1` is a **parameter placeholder** → dynamic at runtime.
- `:one` → expects **exactly one row**.

---

### Example: Multiple Row Query
```sql
-- name: GetAllUsers :many
SELECT * FROM users;
```

- `:many` → returns **slice of structs**.

---

### Example: Insert / Update / Delete
```sql
-- name: CreateUser :exec
INSERT INTO users (name, email) VALUES ($1, $2);
```

- `:exec` → no rows returned, only affected rows

---

## 5️⃣ How to Generate Go Code

1. Write your schema in `DB/Schema/` (CREATE TABLE statements).  
2. Write your queries in `DB/Query/`.  
3. Run:
```powershell
sqlc generate
```
4. Go code will be generated in `internal/db/`.

- Structs for tables → `models.go`
- Functions for queries → `*.sql.go`

---

## 6️⃣ How to Use Generated Code

```go
ctx := context.Background()
queries := db.New(dbConnection)

// Fetch single user by ID
user, err := queries.GetUserByID(ctx, 5)

// Fetch all users
users, err := queries.GetAllUsers(ctx)

// Create a new user
err := queries.CreateUser(ctx, "Deep", "deep@example.com")
```

- Use **placeholders `$1`, `$2`** in SQL → passed as arguments in Go functions.
- Each SQL query generates **one function**.  
- Structs are automatically mapped from table schema.

---

## 7️⃣ Quick Tips
- Name each query uniquely using `-- name: FunctionName :one|:many|:exec`.  
- Reuse table structs where possible; custom queries may generate separate structs.  
- Organize queries by table for readability.  
- Always run `sqlc generate` after adding new queries.

---

**Memorization Tip:** Think of SQLC as a **translator** between SQL and Go — your queries stay in SQL, Go functions do the work automatically.