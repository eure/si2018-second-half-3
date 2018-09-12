
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `profile_quastion_content` (
  `profile_quastion_id` integer unsigned NOT NULL COMMENT 'profile_quastion ID',
  `type` varchar(255) NOT NULL COMMENT '質問内容の種類',
  `message` text NOT NULL COMMENT '質問項目の内容'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャット式プロフィール入力のマスターデータの詳細';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `profile_quastion_content`;
