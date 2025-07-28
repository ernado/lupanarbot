-- Create "tries" table
CREATE TABLE "tries" (
  "id" uuid NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "type" character varying NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "try_user_id_type" to table: "tries"
CREATE UNIQUE INDEX "try_user_id_type" ON "tries" ("user_id", "type");
