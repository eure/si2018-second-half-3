
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `profile_quastion_choice` (
  `profile_quastion_element_id` integer unsigned NOT NULL COMMENT 'profile_quastion_element ID',
  `content` text NOT NULL COMMENT '選択肢の内容'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャット式プロフィール入力の選択肢';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `profile_quastion_choice`;