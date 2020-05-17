/*
Navicat MySQL Data Transfer

Source Database       : blog
Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

*/
use blog;
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned DEFAULT '0' COMMENT '文章ID',
  `content` text,
  `created_at` DateTime  DEFAULT NULL,
  `created_on` int(10) unsigned DEFAULT '0',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论管理';

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_name` varchar(100) DEFAULT '' COMMENT '文章名称',
  `content` text,
  `created_on` int(10) unsigned DEFAULT '0',
  `created_at` DateTime  DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '作者',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `user_type` int(1) DEFAULT '0' COMMENT '用户类型',
  `face_picture` varchar(100) DEFAULT '' COMMENT '人脸照',
  `icon` varchar(500) DEFAULT '' COMMENT '头像',
  `address` varchar(50) DEFAULT '' COMMENT '地区',
  `security_question` varchar(100) DEFAULT '' COMMENT '密保',
  `duration` varchar(100) DEFAULT '' COMMENT '年限',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `blog_favorite`;
CREATE TABLE `blog_favorite` (
  `id` int(100) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(100) DEFAULT '0' COMMENT '医生id',
  `article_id` int(100) DEFAULT '0' COMMENT '病人id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `blog_follow`;
CREATE TABLE `blog_follow` (
  `id` int(100) unsigned NOT NULL AUTO_INCREMENT,
  `doctor_id` int(100) DEFAULT '0' COMMENT '医生id',
  `patient_id` int(100) DEFAULT '0' COMMENT '病人id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;