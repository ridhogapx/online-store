CREATE TABLE "products" (
  "product_id" varchar PRIMARY KEY NOT NULL,
  "category_id" varchar NOT NULL,
  "product_name" varchar NOT NULL,
  "price" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "category_id" BIGSERIAL PRIMARY KEY NOT NULL,
  "category_name" varchar NOT NULL
);

CREATE TABLE "customers" (
  "customer_id" varchar PRIMARY KEY NOT NULL,
  "customer_name" varchar NOT NULL,
  "customer_address" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "shopping_carts" (
  "cart_id" varchar PRIMARY KEY NOT NULL,
  "customer_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transaction_reports" (
  "transaction_id" varchar PRIMARY KEY NOT NULL,
  "cart_id" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");

ALTER TABLE "shopping_carts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "shopping_carts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "transaction_reports" ADD FOREIGN KEY ("cart_id") REFERENCES "shopping_carts" ("cart_id");