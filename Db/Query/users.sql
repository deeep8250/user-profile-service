-- name: GetUserByIdNew :one
select * from users where id=$1;

-- name: CreateUser :one
insert into users (email,name,password)
values($1,$2,$3)
   RETURNING id, email, name, password, created_at, updated_at;