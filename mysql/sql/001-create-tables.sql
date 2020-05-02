---- create ----
create table IF not exists `places`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(100) NOT NULL,
 `detail`           VARCHAR(10000) NOT NULL,
 `address`          VARCHAR(100) NOT NULL,
 `lat`              INT NOT NULL,
 `lng`              INT NOT NULL,
 `img`	            MEDIUMBLOB NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `users`
(
 `id`               INT(20) AUTO_INCREMENT,
 `first_name`       VARCHAR(50) NOT NULL,
 `last_name`        VARCHAR(50) NOT NULL,
 `email`            VARCHAR(50) NOT NULL,
 `password`         VARCHAR(50) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;