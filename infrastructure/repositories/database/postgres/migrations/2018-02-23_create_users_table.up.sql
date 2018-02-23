 
CREATE TABLE "users" (
  "id" serial NOT NULL,
  "name" character varying NOT NULL,
  "email" character varying NOT NULL,
  "password" character varying NOT NULL,
  "is_disabled" boolean NOT NULL,
  "email_verified" boolean NOT NULL,
  "email_token" character varying NOT NULL,
  "zip_code" integer NOT NULL,
  "failed_logins" smallint NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);