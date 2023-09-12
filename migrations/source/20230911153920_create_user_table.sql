-- +goose Up
-- +goose StatementBegin
CREATE TABLE userschema.user (
    id uuid,
    email character varying(255),
    password character varying(255),
    PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE userschema.user;
-- +goose StatementEnd
