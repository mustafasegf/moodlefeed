CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "line_id" varchar UNIQUE NOT NULL,
  "scele_id" int UNIQUE NOT NULL,
  "token" varchar(64) UNIQUE NOT NULL
);

CREATE TABLE "courses" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "short_name" varchar,
  "long_name" varchar,
  "resource" jsonb,
  "course_id" int UNIQUE NOT NULL,
  "user_token" varchar(64) NOT NULL
);

CREATE TABLE "token_course" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "token" varchar(64) NOT NULL,
  "course_id" int NOT NULL
);

CREATE TABLE "message_type" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "name" varchar(32) NOT NULL
);

CREATE TABLE "user_subscribe" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "user_id" int NOT NULL,
  "type_id" int NOT NULL,
  "course_id" int NOT NULL
);

ALTER TABLE "courses" ADD FOREIGN KEY ("user_token") REFERENCES "users" ("token");

ALTER TABLE "token_course" ADD FOREIGN KEY ("token") REFERENCES "users" ("token");

ALTER TABLE "token_course" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("scele_id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("type_id") REFERENCES "message_type" ("id");

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

CREATE INDEX "course_type" ON "user_subscribe" ("course_id", "type_id");
