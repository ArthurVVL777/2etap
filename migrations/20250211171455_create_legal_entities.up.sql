-- Удаляем таблицу, если она уже существует
DROP TABLE IF EXISTS "public"."legalentities" CASCADE;

-- Создаём таблицу
CREATE TABLE "public"."legalentities" (
    "uuid" UUID DEFAULT gen_random_uuid(),
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "deleted_at" TIMESTAMPTZ,
    PRIMARY KEY ("uuid")
);

-- Индексы для оптимизации запросов
CREATE INDEX "legalentities_created_at" ON "public"."legalentities" ("created_at" DESC);
CREATE INDEX "legalentities_updated_at" ON "public"."legalentities" ("updated_at" DESC);

-- Удаляем старую функцию, если она уже есть (аналогично `federations`)
DROP FUNCTION IF EXISTS update_timestamp() CASCADE;

-- Создаём новую функцию обновления поля updated_at
CREATE FUNCTION update_timestamp() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создаём триггер для автоматического обновления updated_at
CREATE TRIGGER update_legalentities_updated_at
BEFORE UPDATE ON "public"."legalentities"
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();