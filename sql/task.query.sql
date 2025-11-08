-- name: GetTask :one
SELECT * FROM tasks
WHERE task_id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY search_location;

-- name: CreateTask :one
INSERT INTO tasks(task_id, search_location,location_id,delay_in_seconds,task_state,task_type,search_keyword,distance_radius,created_at,updated_at)
 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING *;

-- name: UpdateTask :exec
UPDATE tasks SET search_location = $1,location_id = $2,delay_in_seconds = $3,task_state = $4,search_keyword = $5,distance_radius = $6,updated_at = $7 WHERE task_id = $8
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE task_id = $1;