
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `profile_quastion` (
  `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `priority` integer NOT NULL COMMENT '優先順位',
  `name` varchar(255) NOT NULL COMMENT '質問項目の名前',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャット式プロフィール入力のマスターデータ';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `profile_quastion`;
