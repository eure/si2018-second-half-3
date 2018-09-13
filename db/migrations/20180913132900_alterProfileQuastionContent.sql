
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `profile_quastion_content` CHANGE `profile_quastion_id` `profile_quastion_element_id` integer unsigned;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `profile_quastion_content` CHANGE `profile_quastion_element_id` `profile_quastion_id` integer unsigned;