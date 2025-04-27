BEGIN;

CREATE SEQUENCE IF NOT EXISTS rekening_number_seq START 10000000000;

CREATE TABLE IF NOT EXISTS "public"."nasabah" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "rekening_number" bigint NOT NULL DEFAULT nextval('rekening_number_seq'),
    "name" varchar(255) NOT NULL,
    "nik" varchar(255) NOT NULL,
    "phone_number" varchar(21) NOT NULL,
    "total_money" bigint NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rekening_number_nasabah ON "public"."nasabah" ("rekening_number");
CREATE INDEX IF NOT EXISTS idx_nik_nasabah ON "public"."nasabah" ("nik");
CREATE INDEX IF NOT EXISTS idx_phone_number_nasabah ON "public"."nasabah" ("phone_number");

COMMIT;
