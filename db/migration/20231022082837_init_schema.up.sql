CREATE TABLE "projects" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "port" int,
  "is_gen" int,
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "tb" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "project_id" integer,
  "describe" text,
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "tb_fields" (
  "id" serial PRIMARY KEY,
  "field_name" varchar,
  "show_name" varchar,
  "migration" varchar,
  "model_type" varchar,
  "is_require" int,
  "table_id" integer
);


ALTER TABLE "tb" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");
