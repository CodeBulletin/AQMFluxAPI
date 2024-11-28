// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: attribute.sql

package repo

import (
	"context"
)

const attributeIdFromName = `-- name: AttributeIdFromName :one
SELECT attribute_id FROM Attribute WHERE attribute_name = $1
`

func (q *Queries) AttributeIdFromName(ctx context.Context, attributeName string) (int32, error) {
	row := q.queryRow(ctx, q.attributeIdFromNameStmt, attributeIdFromName, attributeName)
	var attribute_id int32
	err := row.Scan(&attribute_id)
	return attribute_id, err
}

const createAttribute = `-- name: CreateAttribute :exec
INSERT INTO Attribute (attribute_id, attribute_name, attribute_desc) VALUES ($1, $2, $3)
`

type CreateAttributeParams struct {
	AttributeID   int32  `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	AttributeDesc string `json:"attribute_desc"`
}

func (q *Queries) CreateAttribute(ctx context.Context, arg CreateAttributeParams) error {
	_, err := q.exec(ctx, q.createAttributeStmt, createAttribute, arg.AttributeID, arg.AttributeName, arg.AttributeDesc)
	return err
}

const getAllAttributes = `-- name: GetAllAttributes :many
SELECT attribute_id, attribute_name, attribute_desc FROM Attribute
`

func (q *Queries) GetAllAttributes(ctx context.Context) ([]Attribute, error) {
	rows, err := q.query(ctx, q.getAllAttributesStmt, getAllAttributes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Attribute{}
	for rows.Next() {
		var i Attribute
		if err := rows.Scan(&i.AttributeID, &i.AttributeName, &i.AttributeDesc); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAttribute = `-- name: UpdateAttribute :exec
UPDATE Attribute SET attribute_name = $1, attribute_desc = $2 WHERE attribute_id = $3
`

type UpdateAttributeParams struct {
	AttributeName string `json:"attribute_name"`
	AttributeDesc string `json:"attribute_desc"`
	AttributeID   int32  `json:"attribute_id"`
}

func (q *Queries) UpdateAttribute(ctx context.Context, arg UpdateAttributeParams) error {
	_, err := q.exec(ctx, q.updateAttributeStmt, updateAttribute, arg.AttributeName, arg.AttributeDesc, arg.AttributeID)
	return err
}