-- Удаляем таблицу, если уже существует
DROP TABLE IF EXISTS "public"."bank_accounts" CASCADE;

-- Создаём таблицу банковских счетов
CREATE TABLE "public"."bank_accounts" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "legal_entity_id" UUID NOT NULL,
    "bik" VARCHAR(20) NOT NULL,
    "bank_name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "corr_account" VARCHAR(50) NOT NULL,
    "account_number" VARCHAR(50) NOT NULL,
    "currency" VARCHAR(10) NOT NULL,
    "comment" TEXT NOT NULL,
    FOREIGN KEY (legal_entity_id) REFERENCES legal_entities(uuid) ON DELETE CASCADE
);
