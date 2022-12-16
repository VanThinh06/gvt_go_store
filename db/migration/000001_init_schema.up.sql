CREATE TABLE "category" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL,
  "name" varchar NOT NULL,
  "national" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY (id)
);
 
CREATE TABLE "product" (
    "id" UUID DEFAULT uuid_generate_v4() NOT NULL,
    "id_category" UUID NOT NULL,
    "name" varchar NOT NULL,
    "price" BIGINT,
    "image" VARCHAR,
    "list_image" VARCHAR [],
    "description" varchar,
    "sold" integer,
    "status" INTEGER,
    "sale" integer,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "update_at" timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id) 
);

CREATE TYPE type AS ENUM (
  'admin',
  'user'
);

CREATE TABLE "user" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL,
  "name" varchar,
  "address" varchar,
  "phone" varchar,
  "email" varchar UNIQUE NOT NULL,
  "type_user" type NOT NULL,
  "password" varchar NOT NULL,
  "payment" varchar,
  "payment_info" varchar,
  "payment_number" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY (id)
);

CREATE TABLE "transaction" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL,
  "id_user" UUID NOT NULL,
  "status" int,
  "amount" BIGINT,
  "message" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY (id)
);

CREATE TABLE "order" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL,
  "id_transaction" UUID NOT NULL,
  "id_product" UUID NOT NULL,
  "amount" BIGINT,
  "data" varchar,
  "status" int,
  "qty" integer,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY (id)
);

ALTER TABLE "product" ADD FOREIGN KEY ("id_category") REFERENCES "category" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("id_user") REFERENCES "user" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("id_transaction") REFERENCES "transaction" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("id_product") REFERENCES "product" ("id");
