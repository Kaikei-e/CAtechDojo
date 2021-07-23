Create Database CAtech;

Use CAtech;


Create table CAtech.users(id int not null primary key AUTO_INCREMENT, name varchar(20), token varchar(200));

CREATE TABLE `CAtech`.`gacha_table` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `chara_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE `CAtech`.`userPossession` (
  `uniqueID` CHAR(100) NOT NULL,
  `chara_ID` INT NOT NULL,
  `chara_name` VARCHAR(45) NOT NULL,
  `userID` INT NOT NULL,
  PRIMARY KEY (`uniqueID`));

  CREATE TABLE `CAtech`.`gacha_rate` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `chara_id` INT NOT NULL,
  `chara_rarity` INT NOT NULL DEFAULT 0,
  `weight` INT NOT NULL,
  PRIMARY KEY (`id`));
