-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-09-22T20:23:35.506Z

CREATE TABLE "users" (
  "username" varchar(255) UNIQUE PRIMARY KEY NOT NULL,
  "hashed_password" varchar(500) NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "mobile_number" bigint DEFAULT null,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "is_email_verified" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(500) UNIQUE NOT NULL,
  "price" numeric NOT NULL,
  "created_by" varchar(255) NOT NULL,
  "discount" int NOT NULL DEFAULT 0,
  "category" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "item_images" (
  "id" SERIAL PRIMARY KEY,
  "item_id" int NOT NULL,
  "image_url" varchar(500) NOT NULL
);

CREATE TABLE "addresses" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "country_code" varchar(3) NOT NULL,
  "city" varchar(10) NOT NULL,
  "street" varchar(100) NOT NULL,
  "landmark" varchar(100) NOT NULL,
  "mobile_number" bigint NOT NULL
);

CREATE TABLE "cart" (
  "username" varchar(255) NOT NULL,
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "total_value" int NOT NULL DEFAULT 0
);

CREATE TABLE "cart_items" (
  "cart_id" int NOT NULL,
  "item_id" int NOT NULL,
  "quantity" int DEFAULT 1
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "items" ("id", "name");

CREATE INDEX ON "cart_items" ("cart_id");

CREATE UNIQUE INDEX ON "cart_items" ("cart_id", "item_id");

ALTER TABLE "items" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("username");

ALTER TABLE "item_images" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "cart" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");
