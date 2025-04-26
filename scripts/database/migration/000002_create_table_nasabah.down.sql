BEGIN;

-- Drop indexes first
DROP INDEX IF EXISTS idx_rekening_number_nasabah;
DROP INDEX IF EXISTS idx_nik_nasabah;
DROP INDEX IF EXISTS idx_phone_number_nasabah;

-- Drop table nasabah
DROP TABLE IF EXISTS "public"."nasabah";

-- Drop sequence rekehistning_number_seq
DROP SEQUENCE IF EXISTS rekening_number_seq;

COMMIT;
