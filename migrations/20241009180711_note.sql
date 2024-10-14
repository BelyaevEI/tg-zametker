-- +goose Up
-- +goose StatementBegin
CREATE TABLE note
(
    user_id BIGINT PRIMARY KEY NOT NULL,
    note text NOT NULL PRIMARY KEY,
    note_time TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE note;
-- +goose StatementEnd
