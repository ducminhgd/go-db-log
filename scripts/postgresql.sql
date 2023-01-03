DROP TABLE IF EXISTS public."event_log";
CREATE TABLE public."event_log" (
    "id" VARCHAR(36) NOT NULL,
    "event_type" VARCHAR(100) NOT NULL,
    "object_type" VARCHAR(100) NOT NULL,
    "object_id" VARCHAR(100) NOT NULL,
    "actor_type" VARCHAR(100) NOT NULL,
    "actor_id" VARCHAR(100) NOT NULL,
    "data" TEXT NULL,
    "result" SMALLINT NOT NULL DEFAULT 1,
    "timestamps" TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);