CREATE TABLE IF NOT EXISTS "quotes" (
  "id"            VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "text"          TEXT                    NOT NULL,
  "author_id"     uuid,
  "registered_at" TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
  FOREIGN KEY ("author_id") REFERENCES "authors"("id")
);

---- create above / drop below ----

DROP TABLE IF EXISTS "quotes";
