-- +goose Up
-- +goose StatementBegin
CREATE TABLE note
(
    user_id BIGINT NOT NULL,
    note text NOT NULL,
    note_time TIMESTAMP NULL,
    PRIMARY KEY (user_id, note)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE note;
-- +goose StatementEnd
