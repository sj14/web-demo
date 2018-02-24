CREATE TABLE "posts" (
  "id" serial NOT NULL,
  "user_id" integer NOT NULL,
  "text" character varying(280) NOT NULL,
  "updated_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL
);