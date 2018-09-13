
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE profile_quastion RENAME TO profile_quastion_element;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE profile_quastion_element RENAME TO profile_quastion;