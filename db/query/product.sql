-- name: CreateProduct :one 
INSERT INTO products (
    id,
    category_id,
    name,
    price,
)

-- name: FindProductByCategory :many
SELECT products.id, categories.name, products.name, products.price, products.created_at FROM products INNER JOIN categories ON products.category_id=categories.id