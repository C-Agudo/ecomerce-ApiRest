CREATE TABLE IF NOT EXISTS `product` (
    
    `id`             int(11) not null auto_increment,
    `name`           varchar(150),
    `price`           varchar(10),
    PRIMARY KEY (`id`)
)engine = InnoDB
  DEFAULT charset = utf8;