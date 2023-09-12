-- +goose Up
-- +goose StatementBegin
INSERT INTO userschema.user (id, email, password) VALUES
    (gen_random_uuid(), 'test@example.com', '$2a$10$lDc3gVFtY5fk610vJ/ZZUeRWtSTtnNPP4wJt.C4wtQ83tnJQ.uMR6');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE userschema.user;
-- +goose StatementEnd
