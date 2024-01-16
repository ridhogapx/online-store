-- name: CreateProduct :one 
INSERT INTO products (
    product_id,
    category_id,
    product_name,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: FindProductByCategory :many
SELECT products.product_id, categories.category_name, products.product_name, products.price, products.created_at FROM products INNER JOIN categories ON products.category_id=categories.category_id WHERE products.category_id=$1;

-- name: CreateCategory :one
INSERT INTO categories (
    category_id,
    category_name
) VALUES (
    $1, $2
) RETURNING *;