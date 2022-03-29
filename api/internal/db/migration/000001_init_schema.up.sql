CREATE TABLE "event" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text,
  "start_date" date,
  "end_date" date,
  "prefix" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ticket_category" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "prefix" varchar NOT NULL,
  "qty" bigint NOT NULL,
  "price" bigint NOT NULL,
  "start_date" date,
  "end_date" date,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "area" varchar,
  "event_id" uuid NOT NULL
);

CREATE TABLE "ticket" (
  "id" uuid PRIMARY KEY,
  "serial_number" varchar NOT NULL,
  "purchase_date" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "ticket_category_id" uuid NOT NULL
);

CREATE TABLE "customer" (
  "id" uuid PRIMARY KEY,
  "full_name" varchar,
  "email" varchar,
  "password" varchar,
  "phone_number" varchar,
  "confirmation_code" text,
  "confirmation_time" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "customer_order" (
  "id" uuid PRIMARY KEY,
  "order_time" timestamptz NOT NULL,
  "time_paid" timestamptz,
  "total_price" bigint NOT NULL,
  "discount" int,
  "final_price" bigint NOT NULL,
  "customer_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_ticket" (
  "qty" int DEFAULT 1,
  "ticket_id" uuid NOT NULL,
  "customer_order_id" uuid NOT NULL
);

CREATE TABLE "customer_payment" (
  "id" uuid PRIMARY KEY,
  "status" varchar NOT NULL,
  "success_at" timestamptz,
  "failed_reason" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "customer_id" uuid,
  "customer_order_id" uuid NOT NULL,
  "payment_option_id" int NOT NULL
);

CREATE TABLE "payment_option" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

ALTER TABLE "ticket_category" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("ticket_category_id") REFERENCES "ticket_category" ("id");

ALTER TABLE "customer_order" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");

ALTER TABLE "order_ticket" ADD FOREIGN KEY ("ticket_id") REFERENCES "ticket" ("id");

ALTER TABLE "order_ticket" ADD FOREIGN KEY ("customer_order_id") REFERENCES "customer_order" ("id");

ALTER TABLE "customer_payment" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");

ALTER TABLE "customer_payment" ADD FOREIGN KEY ("customer_order_id") REFERENCES "customer_order" ("id");

ALTER TABLE "customer_payment" ADD FOREIGN KEY ("payment_option_id") REFERENCES "payment_option" ("id");

CREATE INDEX ON "event" ("name");

CREATE INDEX ON "ticket_category" ("name");

CREATE INDEX ON "ticket_category" ("event_id");

CREATE UNIQUE INDEX ON "ticket" ("serial_number");

CREATE INDEX ON "ticket" ("ticket_category_id");

CREATE UNIQUE INDEX ON "customer" ("email");

CREATE UNIQUE INDEX ON "customer" ("phone_number");

CREATE INDEX ON "customer" ("full_name");

CREATE INDEX ON "order_ticket" ("ticket_id");

CREATE INDEX ON "order_ticket" ("customer_order_id");

CREATE INDEX ON "customer_payment" ("status");

CREATE INDEX ON "customer_payment" ("customer_id");

CREATE INDEX ON "customer_payment" ("payment_option_id");

COMMENT ON COLUMN "event"."created_at" IS 'When event created';

COMMENT ON COLUMN "event"."updated_at" IS 'When event created';

COMMENT ON COLUMN "ticket_category"."created_at" IS 'When ticket category created';

COMMENT ON COLUMN "ticket_category"."updated_at" IS 'When ticket category created';

COMMENT ON COLUMN "ticket"."created_at" IS 'When ticket created';

COMMENT ON COLUMN "ticket"."updated_at" IS 'When ticket created';

COMMENT ON COLUMN "customer"."created_at" IS 'When customer created';

COMMENT ON COLUMN "customer"."updated_at" IS 'When customer created';

COMMENT ON COLUMN "customer_order"."created_at" IS 'When order created';

COMMENT ON COLUMN "customer_order"."updated_at" IS 'When order created';

COMMENT ON COLUMN "customer_payment"."created_at" IS 'When payment created';

COMMENT ON COLUMN "customer_payment"."updated_at" IS 'When payment created';
