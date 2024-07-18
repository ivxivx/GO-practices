CREATE TABLE transaction
(
    id              UUID PRIMARY KEY,
    internal_data    JSONB NOT NULL,
    external_data    JSONB NULL
);
