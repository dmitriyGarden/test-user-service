-- +goose Up
-- +goose StatementBegin
CREATE schema userschema;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP schema userschema;
-- +goose StatementEnd
