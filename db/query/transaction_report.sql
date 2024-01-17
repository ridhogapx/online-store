-- name: CreateTransactionReport :one
INSERT INTO transaction_reports (
    transaction_id,
    cart_id,
    product_id,
    total_price
) VALUES (
    $1, $2, $3, $4
) RETURNING *;