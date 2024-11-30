-- name: AttributeIdFromName :one
SELECT attribute_id FROM Attribute WHERE attribute_name = $1;

-- name: GetAllAttributes :many
SELECT attribute_id, attribute_name, attribute_desc FROM Attribute;

-- name: GetAttributesList :many
SELECT attribute_id, attribute_name FROM Attribute;

-- name: CreateAttribute :exec
INSERT INTO Attribute (attribute_id, attribute_name, attribute_desc) VALUES ($1, $2, $3);

-- name: UpdateAttribute :exec
UPDATE Attribute SET attribute_name = $1, attribute_desc = $2 WHERE attribute_id = $3;