/*
Navicat MySQL Data Transfer

Source Server         : 临时阿里云
Source Server Version : 50727
Source Host           : 112.74.172.81:3306
Source Database       : mytestdb

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2019-10-27 12:33:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for example_table
-- ----------------------------
DROP TABLE IF EXISTS `example_table`;
CREATE TABLE `example_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of example_table
-- ----------------------------
INSERT INTO `example_table` VALUES ('1', 'Mike', '1');
INSERT INTO `example_table` VALUES ('2', 'Jone', '1');
INSERT INTO `example_table` VALUES ('3', 'Chan', '1');
INSERT INTO `example_table` VALUES ('4', 'Chan1572150508', '1');
INSERT INTO `example_table` VALUES ('5', 'Chan1572150536', '1');
