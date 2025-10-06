-- name: GetUserByIdNew :one
select * from users where id=$1;