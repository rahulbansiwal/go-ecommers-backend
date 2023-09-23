CREATE TABLE "sessions" (
  "id" uuid UNIQUE PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "refresh_token" varchar(500) NOT NULL,
  "client_ip" varchar(255) NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expired_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");