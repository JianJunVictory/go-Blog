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

 Date: 19/09/2018 15:46:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `status` tinyint(255) DEFAULT 0 COMMENT '0:未激活，1：激活',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'kui123@dappworks.cn', '12345678', 0);
INSERT INTO `user` VALUES (2, 'jianjun@dappworks.cn', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 1);
INSERT INTO `user` VALUES (3, '12345@qq.com', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 0);
INSERT INTO `user` VALUES (4, 'kui@dappworks.cn', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 0);
INSERT INTO `user` VALUES (5, '123@dappworks.cn', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 0);
INSERT INTO `user` VALUES (6, '123@dappwork', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 0);
INSERT INTO `user` VALUES (8, '2579757405@qq.com', '7c6272c160883d92dd2a4f069bfbba7b530d1327cddf3f4a8f54dc77f6fc2c2a', 1);

SET FOREIGN_KEY_CHECKS = 1;
