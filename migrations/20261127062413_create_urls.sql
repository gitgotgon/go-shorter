-- +goose Up
CREATE TABLE urls (
	-- Using BIGINT instead of UUID for now, it's sufficient for a single PostgreSQL database, it's compact, easy for indexing.
	-- `id` is NOT visible to users.
	id BIGINT GENERATED ALWAYS AS IDENTITY,
	short_code VARCHAR(16) NOT NULL UNIQUE,
	-- Enforce resonable URL length in application layer.
	original_url TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE urls;
