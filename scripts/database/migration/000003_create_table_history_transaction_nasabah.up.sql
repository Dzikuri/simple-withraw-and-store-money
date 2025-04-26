BEGIN;

CREATE TABLE IF NOT EXISTS "public"."history_transaction_nasabah" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "nasabah_id" uuid NOT NULL,
    "transaction_type" varchar(50) NOT NULL,
    "amount" bigint NOT NULL,
    "description" text NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz NULL,
    
    CONSTRAINT fk_nasabah FOREIGN KEY (nasabah_id) REFERENCES "nasabah" (id)
);

CREATE INDEX idx_nasabah_id_history_transaction ON history_transaction_nasabah (nasabah_id);

CREATE INDEX idx_transaction_type_history_transaction ON history_transaction_nasabah (transaction_type);

COMMIT;
