-- name: GetAllGames :many
SELECT name, description, genre, image_url
FROM GAMES;