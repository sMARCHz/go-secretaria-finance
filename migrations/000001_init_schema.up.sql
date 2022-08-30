CREATE TYPE "currency" AS ENUM (
  'USD',
  'THB'
);

CREATE TYPE "transaction_type" AS ENUM (
  'WITHDRAW',
  'DEPOSIT',
  'TRANSFER'
);

CREATE TABLE "accounts" (
  "account_id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" decimal NOT NULL,
  "currency" currency NOT NULL DEFAULT 'THB',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "entry_id" bigserial PRIMARY KEY,
  "account_id" int NOT NULL,
  "category_id" int NOT NULL,
  "amount" decimal NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "category_id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "name_abbr" varchar NOT NULL,
  "transaction_type" transaction_type NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "transfer_id" bigserial PRIMARY KEY,
  "from_account_id" int NOT NULL,
  "to_account_id" int NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "entries" ("category_id");

CREATE INDEX ON "entries" ("account_id", "category_id");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "categories" ("name_abbr");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("account_id");
