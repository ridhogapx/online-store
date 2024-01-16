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
    category_name
) VALUES (
    $1
) RETURNING *;

-- name: CreateCart :one
INSERT into shopping_carts (
    cart_id,
    customer_id, 
    product_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: FindCart :many
SELECT shopping_carts.cart_id, customers.customer_name, products.product_name FROM shopping_carts
    INNER JOIN customers ON shopping_carts.customer_id=customers.customer_id
    INNER JOIN products on shopping_carts.product_id=products.product_id
    WHERE customers.customer_name=$1;