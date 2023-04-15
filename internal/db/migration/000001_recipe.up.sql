CREATE TABLE "recipes" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT NOT NULL,
  "ingredients" TEXT[] NOT NULL,
  "steps" JSONB NOT NULL,
  "total_time" SMALLINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
