// Code generated by sqlc. DO NOT EDIT.
// source: ride.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createRide = `-- name: CreateRide :one
INSERT INTO rides(
    created_at,
    idempotency_key_id,
    origin_lat,
    origin_lon,
    target_lat,
    target_lon,
    stripe_charge_id,
    user_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, idempotency_key_id, origin_lat, origin_lon, target_lat, target_lon, stripe_charge_id, user_id
`

type CreateRideParams struct {
	CreatedAt        time.Time
	IdempotencyKeyID sql.NullInt64
	OriginLat        float64
	OriginLon        float64
	TargetLat        float64
	TargetLon        float64
	StripeChargeID   sql.NullString
	UserID           int64
}

func (q *Queries) CreateRide(ctx context.Context, arg CreateRideParams) (Ride, error) {
	row := q.db.QueryRowContext(ctx, createRide,
		arg.CreatedAt,
		arg.IdempotencyKeyID,
		arg.OriginLat,
		arg.OriginLon,
		arg.TargetLat,
		arg.TargetLon,
		arg.StripeChargeID,
		arg.UserID,
	)
	var i Ride
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.IdempotencyKeyID,
		&i.OriginLat,
		&i.OriginLon,
		&i.TargetLat,
		&i.TargetLon,
		&i.StripeChargeID,
		&i.UserID,
	)
	return i, err
}

const getRideByID = `-- name: GetRideByID :one
SELECT id, created_at, idempotency_key_id, origin_lat, origin_lon, target_lat, target_lon, stripe_charge_id, user_id FROM rides
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRideByID(ctx context.Context, id int64) (Ride, error) {
	row := q.db.QueryRowContext(ctx, getRideByID, id)
	var i Ride
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.IdempotencyKeyID,
		&i.OriginLat,
		&i.OriginLon,
		&i.TargetLat,
		&i.TargetLon,
		&i.StripeChargeID,
		&i.UserID,
	)
	return i, err
}

const getRideByIdempotencyKeyID = `-- name: GetRideByIdempotencyKeyID :one
SELECT id, created_at, idempotency_key_id, origin_lat, origin_lon, target_lat, target_lon, stripe_charge_id, user_id FROM rides
WHERE idempotency_key_id = $1 LIMIT 1
`

func (q *Queries) GetRideByIdempotencyKeyID(ctx context.Context, idempotencyKeyID sql.NullInt64) (Ride, error) {
	row := q.db.QueryRowContext(ctx, getRideByIdempotencyKeyID, idempotencyKeyID)
	var i Ride
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.IdempotencyKeyID,
		&i.OriginLat,
		&i.OriginLon,
		&i.TargetLat,
		&i.TargetLon,
		&i.StripeChargeID,
		&i.UserID,
	)
	return i, err
}

const updateRide = `-- name: UpdateRide :one
UPDATE rides SET
    stripe_charge_id=$2
WHERE id = $1
RETURNING id, created_at, idempotency_key_id, origin_lat, origin_lon, target_lat, target_lon, stripe_charge_id, user_id
`

type UpdateRideParams struct {
	ID             int64
	StripeChargeID sql.NullString
}

func (q *Queries) UpdateRide(ctx context.Context, arg UpdateRideParams) (Ride, error) {
	row := q.db.QueryRowContext(ctx, updateRide, arg.ID, arg.StripeChargeID)
	var i Ride
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.IdempotencyKeyID,
		&i.OriginLat,
		&i.OriginLon,
		&i.TargetLat,
		&i.TargetLon,
		&i.StripeChargeID,
		&i.UserID,
	)
	return i, err
}
