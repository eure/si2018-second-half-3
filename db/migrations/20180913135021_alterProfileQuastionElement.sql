
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `profile_quastion_element` ADD `form_type` varchar(255) NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `profile_quastion_element` DROP `form_type`;