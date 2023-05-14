CREATE TABLE document (
    "id" uuid NOT NULL PRIMARY KEY,
    "text" TEXT,
    "created_at" TIMESTAMP,
    "last_read_at" TIMESTAMP
);