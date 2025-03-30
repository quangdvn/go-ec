-- name: GetUserByEmailSQLC :one
SELECT usr_email, usr_id FROM `pre_go_crm_user_c`
WHERE usr_email = ? LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `pre_go_crm_user_c`
SET usr_status = $2,
    usr_updated_at = $3
WHERE usr_id = $1;

-- -- name: CreateUserSQLC :one
-- insert into users (   name, email, password ) VALUES (   $1, $2, $3 ) RETURNING *;

-- -- name: UpdateUserSQLC :one
-- UPDATE users
-- SET name = $2, email = $3, password = $4
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteUserSQLC :exec
-- DELETE FROM users
-- WHERE id = $1;