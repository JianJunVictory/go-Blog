/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50560
 Source Host           : localhost:3306
 Source Schema         : jun

 Target Server Type    : MySQL
 Target Server Version : 50560
 File Encoding         : 65001

 Date: 19/09/2018 15:44:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `articleId` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `userId` int(11) NOT NULL,
  `articleTitle` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `articleContent` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `articleViews` tinytext CHARACTER SET utf8 COLLATE utf8_general_ci,
  `articleCommentCount` int(11) DEFAULT NULL,
  `articleDate` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `articleLikeCount` int(9) DEFAULT NULL,
  PRIMARY KEY (`articleId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
