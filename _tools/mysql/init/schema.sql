CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
  `name` varchar(20) NOT NULL COMMENT 'ユーザー名',
  `created_at` DATETIME(6) NOT NULL COMMENT 'ユーザー作成日時',
  `updated_at` DATETIME(6) NOT NULL COMMENT 'ユーザー更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'ユーザー';

CREATE TABLE `balances` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '残高ID',
  `user_id` BIGINT UNSIGNED NOT NULL UNIQUE COMMENT 'ユーザーID',
  `amount` BIGINT UNSIGNED NOT NULL COMMENT '残高',
  `created_at` DATETIME(6) NOT NULL COMMENT '残高作成日時',
  `updated_at` DATETIME(6) NOT NULL COMMENT '残高更新日時',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_id_on_balance` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '残高';

CREATE TABLE `transfer_trans` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '転送トランザクションID',
  `from_user` BIGINT UNSIGNED NOT NULL COMMENT '転送する側のユーザーID',
  `from_balance` BIGINT UNSIGNED NOT NULL COMMENT '転送する側の残高ID',
  `to_user` BIGINT UNSIGNED NOT NULL COMMENT '受け取る側のユーザーID',
  `to_balance` BIGINT UNSIGNED NOT NULL COMMENT '受け取る側の残高ID',
  `amount` BIGINT UNSIGNED NOT NULL COMMENT '転送された残高',
  `processed_at` DATETIME(6) NOT NULL COMMENT '処理日時',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_from_user_on_transfer_trans` FOREIGN KEY (`from_user`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_from_balance_on_transfer_trans` FOREIGN KEY (`from_balance`) REFERENCES `balances` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_to_user_on_transfer_trans` FOREIGN KEY (`to_user`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_to_balance_on_transfer_trans` FOREIGN KEY (`to_balance`) REFERENCES `balances` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '転送トランザクション';

CREATE TABLE `balance_trans` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '残高トランザクションID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーID',
  `balance_id` BIGINT UNSIGNED NOT NULL COMMENT '残高ID',
  `transfer_id` BIGINT UNSIGNED COMMENT '転送トランザクションID',
  `amount` BIGINT NOT NULL COMMENT '残高',
  `processed_at` DATETIME(6) NOT NULL COMMENT '処理日時',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_id_on_balance_trans` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_balance_id_on_balance_trans` FOREIGN KEY (`balance_id`) REFERENCES `balances` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_transfer_id_on_balance_trans` FOREIGN KEY (`transfer_id`) REFERENCES `transfer_trans` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '残高トランザクション';
