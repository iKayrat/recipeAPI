-- name: CreateRecipe :one
INSERT INTO recipes (
    name, 
    description, 
    ingredients, 
    steps,
    total_time,
    created_at,
    updated_at
) VALUES (
      $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetRecipeByID :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: GetRecipes :many
SELECT 
    id,
    name,
    description,
    ingredients,
    steps,
    total_time,
    created_at,
    updated_at
FROM recipes ORDER BY id;

-- name: GetRecipesByTime :many
SELECT 
    id,
    name,
    description,
    ingredients,
    steps,
    total_time,
    created_at,
    updated_at
FROM recipes ORDER BY total_time;

-- name: UpdateRecipe :one
UPDATE recipes
SET
	name = $2,
	description = $3,
	ingredients = $4,
	steps = $5,
	total_time = $6,
	updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteRecipeByID :one
DELETE FROM recipes
WHERE id = $1
RETURNING id;

-- name: GetRecipeByIngredients :many
SELECT * FROM recipes 
WHERE ingredients @> ARRAY[$1]
LIMIT 100 OFFSET 0;