CREATE TABLE "token"
(
    "user_id" character varying(128) DEFAULT '' NOT NULL,
    "token"   character varying(128) DEFAULT '' NOT NULL,
    CONSTRAINT "constraint_id_unique" UNIQUE ("user_id", "token")
) WITH (oids = false);