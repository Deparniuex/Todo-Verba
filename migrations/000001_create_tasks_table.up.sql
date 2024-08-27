CREATE TABLE IF NOT EXISTS tasks (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    due_date timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);