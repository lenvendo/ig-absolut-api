CREATE TABLE "users"
(
    "id"           character varying(128) DEFAULT ''    NOT NULL,
    "is_confirmed" boolean                DEFAULT false NOT NULL,
    "created_at"   timestamp,
    "updated_at"   timestamp              DEFAULT now() NOT NULL,
    CONSTRAINT "constraint_id_unique" UNIQUE ("id")
) WITH (oids = false);