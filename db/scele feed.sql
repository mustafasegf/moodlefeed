CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "line_id" varchar UNIQUE,
  "scele_id" int UNIQUE,
  "token" varchar(64)
);

CREATE TABLE "courses" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "short_name" varchar,
  "long_name" varchar,
  "resource" json,
  "course_id" int,
  "user_token" varchar(64)
);

CREATE TABLE "token_course" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "token" varchar(64),
  "course_id" int
);

CREATE TABLE "message_type" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "name" varchar(32)
);

CREATE TABLE "user_subscribe" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "user_id" int,
  "type_id" int,
  "course_id" int
);

ALTER TABLE "courses" ADD FOREIGN KEY ("user_token") REFERENCES "users" ("token");

ALTER TABLE "token_course" ADD FOREIGN KEY ("token") REFERENCES "users" ("token");

ALTER TABLE "token_course" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("type_id") REFERENCES "message_type" ("id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

CREATE INDEX "course_type" ON "user_subscribe" ("course_id", "type_id");
