/*
 Navicat Premium Data Transfer

 Source Server         : local-mysql-docker
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : 127.0.0.1:3306
 Source Schema         : db_puti

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 20/08/2018 20:13:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pt_comment_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_comment_meta`;
CREATE TABLE `pt_comment_meta`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `comment_id` int(11) UNSIGNED NOT NULL DEFAULT 0,
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `meta_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_id`(`comment_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_comments
-- ----------------------------
DROP TABLE IF EXISTS `pt_comments`;
CREATE TABLE `pt_comments`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论id',
  `post_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论的文章或页面id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `if_visitor` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否游客;1是,0不是;默认游客',
  `commenter_user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论者id;是游客时为0;默认为0',
  `commenter_name` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论者名称',
  `commenter_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论者email',
  `commenter_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论者链接',
  `commenter_ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论者ip',
  `approved` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '是否通过(开启评论审核后，通过后显示)',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论来源agent',
  `created_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论时间(UTC)',
  `updated_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_post_ID`(`post_id`) USING BTREE,
  INDEX `comment_date_gmt`(`updated_time`) USING BTREE,
  INDEX `comment_parent`(`parent_id`) USING BTREE,
  INDEX `comment_author_email`(`commenter_email`(10)) USING BTREE,
  INDEX `comment_approved_date_gmt`(`updated_time`, `approved`) USING BTREE,
  INDEX `comment_deleted`(`deleted_time`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_links
-- ----------------------------
DROP TABLE IF EXISTS `pt_links`;
CREATE TABLE `pt_links`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '链接id',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '链接url',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '链接名称',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '链接图像地址',
  `target` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '目标(如_blank)',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '链接描述',
  `visible` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Y' COMMENT '是否可见',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 1 COMMENT '所属用户',
  `rating` int(11) NOT NULL DEFAULT 0 COMMENT '评分',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间',
  `notes` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '备注',
  `rss` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'rss地址',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `link_visible`(`visible`) USING BTREE,
  INDEX `link_owner_user`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_options
-- ----------------------------
DROP TABLE IF EXISTS `pt_options`;
CREATE TABLE `pt_options`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置id',
  `option_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '配置名称',
  `option_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '对应的值',
  `autoload` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否自动加载;默认0不自动加载',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `option_name`(`option_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_post_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_post_meta`;
CREATE TABLE `pt_post_meta`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `post_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'post_id',
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设置的key',
  `meta_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '设置的value',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `post_id`(`post_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_posts
-- ----------------------------
DROP TABLE IF EXISTS `pt_posts`;
CREATE TABLE `pt_posts`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发表人id',
  `post_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'article' COMMENT '类型：article，page',
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `content_markdown` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'markdown格式文章内容',
  `content_html` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'html格式文章内容',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '缩略名（用于url中展示）',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id（如果有）',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'publish' COMMENT '状态',
  `comment_status` tinyint(20) NOT NULL DEFAULT 1 COMMENT '评论状态(是否开启);默认1开启；0关闭',
  `guid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一链接',
  `cover_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '封面图片链接',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT '评论数目',
  `view_count` int(11) NOT NULL DEFAULT 0 COMMENT '浏览量',
  `posted_time` datetime(0) NOT NULL COMMENT '发表时间（UTC）',
  `created_time` datetime(0) NOT NULL COMMENT '创建时间（UTC）',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间（UTC）',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间（UTC）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `post_parent`(`parent_id`) USING BTREE,
  INDEX `post_author`(`user_id`) USING BTREE,
  INDEX `type_status_date`(`id`, `post_type`, `status`, `posted_time`) USING BTREE,
  INDEX `post_name`(`slug`(191)) USING BTREE,
  FULLTEXT INDEX `post_title`(`title`),
  INDEX `post_deleted`(`deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_resource_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_resource_meta`;
CREATE TABLE `pt_resource_meta`  (
  `meta_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `resource_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '资源id',
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设置的key',
  `meta_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '设置的value',
  PRIMARY KEY (`meta_id`) USING BTREE,
  INDEX `resource_id`(`resource_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_resources
-- ----------------------------
DROP TABLE IF EXISTS `pt_resources`;
CREATE TABLE `pt_resources`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `upload_user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有者id',
  `post_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '归属的post_id\r\n',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '资源名称',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '资源说明',
  `guid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一链接',
  `upload_date` datetime(0) NOT NULL COMMENT '上传时间',
  `upload_date_gmt` datetime(0) NOT NULL COMMENT '上传时GMT时间',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'picture' COMMENT '资源类型；默认picture',
  `mime_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '资源文件类型',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '资源状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `resource_parent`(`post_id`) USING BTREE,
  INDEX `resource_type`(`id`, `upload_date`, `type`, `status`) USING BTREE,
  INDEX `resource_name`(`slug`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 235 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '资源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_subject_relationships
-- ----------------------------
DROP TABLE IF EXISTS `pt_subject_relationships`;
CREATE TABLE `pt_subject_relationships`  (
  `object_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '附属于专题的项目id（一般是文章）',
  `subject_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '专题id',
  `order_num` int(11) NOT NULL DEFAULT 0 COMMENT '排序值',
  PRIMARY KEY (`object_id`, `subject_id`) USING BTREE,
  INDEX `subject_id`(`subject_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '专题关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_subjects
-- ----------------------------
DROP TABLE IF EXISTS `pt_subjects`;
CREATE TABLE `pt_subjects`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '专题 id',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专题名称',
  `slug` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专题缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `cover_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '封面图',
  `is_end` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否末级；1是 0不是',
  `count` int(11) NOT NULL DEFAULT 0 COMMENT '拥有文章数量',
  `last_updated` datetime(0) NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '上次更新',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `subject_slug`(`slug`(191)) USING BTREE,
  INDEX `subkect_parent`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '专题表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_term_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_term_meta`;
CREATE TABLE `pt_term_meta`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `term_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类条目id',
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '属性名称',
  `meta_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '属性值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `term_id`(`term_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_term_relationships
-- ----------------------------
DROP TABLE IF EXISTS `pt_term_relationships`;
CREATE TABLE `pt_term_relationships`  (
  `object_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '归属分类的对象id',
  `term_taxonomy_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类id',
  `term_order` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`object_id`, `term_taxonomy_id`) USING BTREE,
  INDEX `term_taxonomy_id`(`term_taxonomy_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_term_taxonomy
-- ----------------------------
DROP TABLE IF EXISTS `pt_term_taxonomy`;
CREATE TABLE `pt_term_taxonomy`  (
  `term_taxonomy_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类方式id',
  `term_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'term_id',
  `parent_term_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父term_id',
  `taxonomy` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类方式',
  `term_group` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分组',
  PRIMARY KEY (`term_taxonomy_id`) USING BTREE,
  UNIQUE INDEX `term_id_taxonomy`(`term_id`, `taxonomy`) USING BTREE,
  INDEX `taxonomy`(`taxonomy`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 74 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_terms
-- ----------------------------
DROP TABLE IF EXISTS `pt_terms`;
CREATE TABLE `pt_terms`  (
  `term_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '条件id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '条件名称',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有的数目',
  PRIMARY KEY (`term_id`) USING BTREE,
  INDEX `slug`(`slug`(191)) USING BTREE,
  INDEX `name`(`name`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 74 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_user_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_user_meta`;
CREATE TABLE `pt_user_meta`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设置的key',
  `meta_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '设置的value',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_users
-- ----------------------------
DROP TABLE IF EXISTS `pt_users`;
CREATE TABLE `pt_users`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录帐号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '头像',
  `page_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '主页链接',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '状态.0激活1冻结',
  `role` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'subscriber' COMMENT '用户角色',
  `created_time` datetime(0) NOT NULL COMMENT '注册时间(UTC)',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`account`) USING BTREE,
  UNIQUE INDEX `user_email_2`(`email`) USING BTREE,
  INDEX `user_login_key`(`account`) USING BTREE,
  INDEX `user_nicename`(`nickname`) USING BTREE,
  INDEX `user_email`(`email`) USING BTREE,
  INDEX `user_delete`(`deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
