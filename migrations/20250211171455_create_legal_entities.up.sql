-- Удаляем таблицу, если она уже существует
DROP TABLE IF EXISTS "public"."legal_entities" CASCADE;

-- Создаём таблицу
CREATE TABLE "public"."legal_entities" (
    "uuid" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "deleted_at" TIMESTAMPTZ
);

-- Индексы для оптимизации запросов
CREATE INDEX "legal_entities_created_at_idx" ON "public"."legal_entities" ("created_at" DESC);
CREATE INDEX "legal_entities_updated_at_idx" ON "public"."legal_entities" ("updated_at" DESC);

-- Удаляем старую функцию, если она уже есть (на всякий случай)
DROP FUNCTION IF EXISTS update_timestamp() CASCADE;

-- Создаём новую функцию обновления поля updated_at
CREATE FUNCTION update_timestamp() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создаём триггер для автоматического обновления updated_at
CREATE TRIGGER update_legal_entities_updated_at
BEFORE UPDATE ON "public"."legal_entities"
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
