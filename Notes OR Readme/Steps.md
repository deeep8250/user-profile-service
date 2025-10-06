# User Profiles with PostgreSQL

Build a **Users & Profiles API** with PostgreSQL.  
Learn schema design, migrations, transactions, sqlc/GORM, pagination, filtering, sorting, and soft deletes.  
Includes **Docker Compose** with Postgres + Adminer.  
A practical project to master real-world database skills.

---

## 🚀 What You’ll Implement

### 1. Schema + SQL Migrations
- Create tables for **users** and **profiles**.
- Write SQL migration files so schema changes are version-controlled.

### 2. Switch to Postgres + sqlc
- Move from SQLite/other DB to **Postgres**.
- Use **sqlc** to generate type-safe Go code for queries.

### 3. Repository / Service Layers with Transactions
- Build repo layer functions (`createUser`, `updateProfile`, etc.).
- Wrap multiple DB operations inside a transaction (e.g., user + profile creation).

### 4. Pagination, Filtering & Sorting
- Implement **cursor-based pagination** (instead of offset for efficiency).
- Add filters (e.g., search by email, status).
- Add sorting (e.g., newest first, alphabetical).

### 5. Soft Deletes + Unique Constraints
- Mark rows as deleted instead of physical deletion.
- Apply unique constraints (e.g., unique email per user).

### 6. Contract Tests for Listing Endpoints
- Write tests to confirm pagination, sorting, and filtering behave correctly.

### 7. Docker Milestone
- Use **docker-compose** to spin up:
  - **Postgres** (the database)
  - **Adminer** (simple web UI to inspect DB)

---

## 🛠️ Tools & Tech
- **Go**
- **PostgreSQL**
- **sqlc** / **GORM**
- **Docker & Docker Compose**
- **Adminer** (DB UI)
- **Testing Frameworks** (Go tests)

---

## 🎯 Learning Outcomes
By the end of this project, you will:
- Confidently design and evolve database schemas.
- Work with SQL migrations for version control.
- Implement repo/service patterns with transactions.
- Master pagination, filtering, and sorting in queries.
- Use soft deletes and unique constraints for integrity.
- Run Postgres in Docker with an Adminer dashboard.
