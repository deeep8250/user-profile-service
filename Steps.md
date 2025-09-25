What You’ll Implement

Schema + SQL migrations

Create tables for users and profiles.

Write SQL migration files so schema changes are version-controlled.

Switch to Postgres + sqlc

Move from SQLite/other DB to Postgres.

Use sqlc to generate type-safe Go code for queries.

Repository / Service Layers with Transactions

Build repo layer functions (createUser, updateProfile, etc.).

Wrap multiple DB operations inside a transaction (e.g., user + profile creation).

Pagination, Filtering & Sorting

Implement cursor-based pagination (instead of offset for efficiency).

Add filters (e.g., search by email, status).

Add sorting (e.g., newest first, alphabetical).

Soft Deletes + Unique Constraints

Mark rows as deleted instead of physical deletion.

Apply unique constraints (e.g., unique email per user).

Contract Tests for Listing Endpoints

Write tests to confirm that pagination, sorting, filtering all behave correctly.

Docker Milestone

Use docker-compose to spin up:

Postgres (the DB)

Adminer (simple web UI to check DB).