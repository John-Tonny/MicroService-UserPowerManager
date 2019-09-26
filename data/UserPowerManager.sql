-- MySQL Script generated by MySQL Workbench
-- Thu Sep 12 18:13:06 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema UserPowerManager
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema UserPowerManager
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `UserPowerManager` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin ;
USE `UserPowerManager` ;

-- -----------------------------------------------------
-- Table `UserPowerManager`.`users_0000`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`users_0000` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`users_0000` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '编号Id',
  `user_name` VARCHAR(200) NULL COMMENT '用户名称',
  `realy_name` VARCHAR(200) NULL COMMENT '真实姓名',
  `password` VARCHAR(100) NOT NULL COMMENT '密码',
  `auth_key` VARCHAR(45) NULL COMMENT 'authkey',
  `email` VARCHAR(100) NULL COMMENT 'email',
  `is_del` TINYINT(1) NULL DEFAULT 0 COMMENT '是否删除：1、true 0、false',
  `note` TEXT NULL COMMENT '加入时间',
  `parent_id` BIGINT UNSIGNED NOT NULL,
  `create_time` DATETIME NOT NULL,
  `last_update_time` DATETIME NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `UserPowerManager`.`users_group`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`users_group` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`users_group` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL COMMENT '用户组名称',
  `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '所属上级Id',
  `sorts` SMALLINT NOT NULL DEFAULT 0 COMMENT '排序',
  `note` VARCHAR(2000) NULL COMMENT '备注',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `last_update_time` DATETIME NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `UserPowerManager`.`casbin_rule`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`casbin_rule` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`casbin_rule` (
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `p_type` VARCHAR(255) NULL COMMENT 'g表示用户组，p表示权限',
  `v0` VARCHAR(255) NULL,
  `v1` VARCHAR(255) NULL,
  `v2` VARCHAR(255) NULL,
  `v3` VARCHAR(255) NULL,
  `v4` VARCHAR(45) NULL,
  `v5` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `UserPowerManager`.`columns`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`columns` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`columns` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NOT NULL COMMENT '栏目名称',
  `URL` VARCHAR(500) NULL COMMENT 'URL',
  `parent_id` BIGINT NOT NULL COMMENT '所属上级Id',
  `sorts` SMALLINT NOT NULL DEFAULT 0 COMMENT '排序',
  `is_show_nav` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否显示在导航',
  `css_icon` VARCHAR(50) NULL COMMENT 'css图标样式',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `last_update_time` DATETIME NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `UserPowerManager`.`history_user_login`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`history_user_login` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`history_user_login` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `geo_remote_addr` VARCHAR(50) NOT NULL COMMENT '用户登录IP',
  `geo_country` VARCHAR(100) NOT NULL COMMENT 'IP所在国家',
  `geo_city` VARCHAR(100) NOT NULL COMMENT 'IP所在城市',
  `device_detector` VARCHAR(1000) NOT NULL COMMENT '设备检测器',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `UserPowerManager`.`users_0001`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `UserPowerManager`.`users_0001` ;

CREATE TABLE IF NOT EXISTS `UserPowerManager`.`users_0001` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '编号Id',
  `user_name` VARCHAR(200) NULL COMMENT '用户名称',
  `realy_name` VARCHAR(200) NULL COMMENT '真实姓名',
  `password` VARCHAR(100) NOT NULL COMMENT '密码',
  `auth_key` VARCHAR(45) NULL COMMENT 'authkey',
  `email` VARCHAR(100) NULL COMMENT 'email',
  `is_del` TINYINT(1) NULL DEFAULT 0 COMMENT '是否删除：1、true 0、false',
  `note` TEXT NULL COMMENT '加入时间',
  `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `create_time` DATETIME NOT NULL,
  `last_update_time` DATETIME NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
