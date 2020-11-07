/*
 Navicat Premium Data Transfer

 Source Server         : 915游戏本地数据
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 07/11/2020 14:10:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `category_id` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类id',
  `column_id` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '栏目id',
  `users_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id 发表者',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述-总结',
  `label` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签',
  `title_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题图片',
  `is_show` tinyint(2) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1公开, 2加密',
  `create_date` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类别名称',
  `column_id` tinyint(3) NULL DEFAULT NULL COMMENT '栏目id',
  `top` tinyint(3) NULL DEFAULT NULL COMMENT '排序',
  `is_show` tinyint(3) NULL DEFAULT NULL COMMENT '是否展示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '类别' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_column
-- ----------------------------
DROP TABLE IF EXISTS `blog_column`;
CREATE TABLE `blog_column`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id\r\n主键id',
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '栏目名称',
  `top` tinyint(3) NULL DEFAULT NULL COMMENT '排序',
  `is_show` tinyint(3) NULL DEFAULT NULL COMMENT '是否展示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '栏目表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for game_reg_201809
-- ----------------------------
DROP TABLE IF EXISTS `game_reg_201809`;
CREATE TABLE `game_reg_201809`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `game_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏id',
  `agent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '渠道id',
  `site_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '广告位id',
  `ip` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册ip（登陆）',
  `imei` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `reg_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间（登陆）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name_2`(`user_name`, `game_id`) USING BTREE,
  INDEX `game_id`(`game_id`) USING BTREE,
  INDEX `agent_id`(`agent_id`) USING BTREE,
  INDEX `reg_time`(`reg_time`) USING BTREE,
  INDEX `reg_ip`(`ip`) USING BTREE,
  INDEX `site_id`(`site_id`) USING BTREE,
  INDEX `imei`(`imei`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7992954 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '按游戏注册' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `migration` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for password_resets
-- ----------------------------
DROP TABLE IF EXISTS `password_resets`;
CREATE TABLE `password_resets`  (
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  INDEX `password_resets_email_index`(`email`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_login_diff
-- ----------------------------
DROP TABLE IF EXISTS `user_login_diff`;
CREATE TABLE `user_login_diff`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `reg_time` int(11) NULL DEFAULT NULL,
  `login_time` int(11) NULL DEFAULT NULL,
  `activity` int(11) NULL DEFAULT NULL,
  `login_num` int(11) NULL DEFAULT NULL,
  `time_diff` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_name`(`user_name`) USING BTREE,
  INDEX `reg_time`(`reg_time`) USING BTREE,
  INDEX `login_time`(`login_time`) USING BTREE,
  INDEX `activity`(`activity`) USING BTREE,
  INDEX `login_num`(`login_num`) USING BTREE,
  INDEX `time_diff`(`time_diff`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5777939 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_login_info
-- ----------------------------
DROP TABLE IF EXISTS `user_login_info`;
CREATE TABLE `user_login_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `game_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆游戏',
  `server_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆区服',
  `agent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户注册渠道id',
  `site_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户注册广告位id',
  `cplaceid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户注册子广告位',
  `adid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户注册创意id',
  `turn` tinyint(4) NOT NULL DEFAULT 0 COMMENT '用户注册创意轮数',
  `login_agent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆渠道id（流失召回渠道id）',
  `login_site_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆广告位id（流失召回广告位id）',
  `imei` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登陆设备号',
  `model` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机机型',
  `devicebrand` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机厂商',
  `systemversion` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机系统版本',
  `mnos` tinyint(4) NOT NULL DEFAULT 0 COMMENT '网络类型 1电信，2联通，3移动，4wifi	',
  `sdk_version_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'SDK 版本号',
  `app_version_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'app(游戏) 版本号',
  `ip` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆ip',
  `login_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆时间',
  `reg_game_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '按游戏注册时间',
  `reg_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_name`(`user_name`) USING BTREE,
  INDEX `game_id`(`game_id`) USING BTREE,
  INDEX `reg_agent_id`(`agent_id`) USING BTREE,
  INDEX `login_agent_id`(`login_agent_id`) USING BTREE,
  INDEX `login_time`(`login_time`) USING BTREE,
  INDEX `reg_game_time`(`reg_game_time`) USING BTREE,
  INDEX `reg_time`(`reg_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11556337 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户登陆日志（新）' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` timestamp(0) NULL DEFAULT NULL,
  `api_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `users_email_unique`(`password`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
