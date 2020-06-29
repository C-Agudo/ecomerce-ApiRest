CREATE TABLE IF NOT EXISTS `purchases` (
    
    `id`             int(11) not null auto_increment,
    `shop`           varchar(150),
    `products`       varchar(255),
    PRIMARY KEY (`id`)
)engine = InnoDB
  DEFAULT charset = utf8;