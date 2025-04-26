BEGIN;

DROP INDEX IF EXISTS idx_transaction_type_history_transaction;
DROP INDEX IF EXISTS idx_nasabah_id_history_transaction;

DROP TABLE IF EXISTS "public"."history_transaction_nasabah";

COMMIT;
