CREATE TABLE IF NOT EXISTS "authors" (
  "id"            uuid         PRIMARY KEY DEFAULT gen_random_uuid(),
  "name"          VARCHAR(255) UNIQUE  NOT NULL,
  "registered_at" TIMESTAMP            NOT NULL DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----

DROP TABLE IF EXISTS "authors";
