-- name: CreateTransactionReport :one
INSERT INTO transaction_reports (
    transaction_id,
    customer_id,
    total_price
) VALUES (
    $1, $2, $3
) RETURNING *;