CREATE TABLE "products" (
  "id" varchar PRIMARY KEY NOT NULL,
  "category_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "price" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "customers" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "shopping_carts" (
  "id" varchar PRIMARY KEY NOT NULL,
  "customer_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transaction_reports" (
  "id" varchar PRIMARY KEY NOT NULL,
  "cart_id" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "shopping_carts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "shopping_carts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "transaction_reports" ADD FOREIGN KEY ("cart_id") REFERENCES "shopping_carts" ("id");