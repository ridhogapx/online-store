-- name: CreateCustomer :one
INSERT INTO customers (
    customer_id, 
    customer_name,
    customer_address,
    email,
    password
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: FindCustomerByEmail :one
SELECT customer_id, customer_name, email, password FROM customers
WHERE email=$1;