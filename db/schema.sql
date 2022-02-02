CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "scele_id" int UNIQUE NOT NULL,
  "token" varchar(64) UNIQUE NOT NULL
);

CREATE TABLE "client" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "scele_id" int,
  "line_id" varchar UNIQUE,
  "discord_id" varchar UNIQUE
);

CREATE TABLE "courses" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "short_name" varchar,
  "long_name" varchar,
  "resource" text,
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
  "scele_id" int NOT NULL,
  "type_id" int NOT NULL,
  "course_id" int NOT NULL
);

ALTER TABLE "client" ADD FOREIGN KEY ("scele_id") REFERENCES "users" ("scele_id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "courses" ADD FOREIGN KEY ("user_token") REFERENCES "users" ("token") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "token_course" ADD FOREIGN KEY ("token") REFERENCES "users" ("token") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "token_course" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("scele_id") REFERENCES "users" ("scele_id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("type_id") REFERENCES "message_type" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_subscribe" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id") ON DELETE CASCADE ON UPDATE CASCADE;

CREATE INDEX "course_type" ON "user_subscribe" ("course_id", "type_id");
