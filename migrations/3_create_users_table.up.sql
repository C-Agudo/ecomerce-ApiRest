CREATE TABLE IF NOT EXISTS `users` (
    
    `id`             int(11) not null auto_increment,
    `name`           varchar(150),
    `password`       varchar(11),
    `role`       varchar(11),
    PRIMARY KEY (`id`)
)engine = InnoDB
  DEFAULT charset = utf8;