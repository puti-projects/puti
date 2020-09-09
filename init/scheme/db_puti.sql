/*
 Navicat Premium Data Transfer

 Source Server         : local-mysql-docker
 Source Server Type    : MySQL
 Source Server Version : 80014
 Source Host           : 127.0.0.1:3306
 Source Schema         : db_puti

 Target Server Type    : MySQL
 Target Server Version : 80014
 File Encoding         : 65001

 Date: 21/03/2019 03:40:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pt_comment
-- ----------------------------
DROP TABLE IF EXISTS `pt_comment`;
CREATE TABLE `pt_comment`  (
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
  `comment_date` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论时间(UTC)',
  `approved` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '是否通过(开启评论审核后，通过后显示)',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论来源agent',
  `created_time` datetime(0) NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_post_ID`(`post_id`) USING BTREE,
  INDEX `comment_parent`(`parent_id`) USING BTREE,
  INDEX `comment_author_email`(`commenter_email`(10)) USING BTREE,
  INDEX `comment_approved_date`(`comment_date`, `approved`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

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
-- Table structure for pt_link
-- ----------------------------
DROP TABLE IF EXISTS `pt_link`;
CREATE TABLE `pt_link`  (
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
-- Table structure for pt_option
-- ----------------------------
DROP TABLE IF EXISTS `pt_option`;
CREATE TABLE `pt_option`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置id',
  `option_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '配置名称',
  `option_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '对应的值',
  `autoload` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否自动加载;默认0不自动加载',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `option_name`(`option_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_option
-- ----------------------------
INSERT INTO `pt_option` VALUES (1, 'blog_name', 'Gogogo', 1);
INSERT INTO `pt_option` VALUES (2, 'blog_description', '一个新的 Puti 站点', 1);
INSERT INTO `pt_option` VALUES (3, 'site_url', 'http://puti.com', 1);
INSERT INTO `pt_option` VALUES (4, 'admin_email', 'example@example.com', 1);
INSERT INTO `pt_option` VALUES (5, 'users_can_register', 'off', 1);
INSERT INTO `pt_option` VALUES (6, 'timezone_string', 'Asia/Shanghai', 1);
INSERT INTO `pt_option` VALUES (7, 'default_category', '1', 0);
INSERT INTO `pt_option` VALUES (8, 'default_link_category', '0', 0);
INSERT INTO `pt_option` VALUES (9, 'show_on_front', 'article', 1);
INSERT INTO `pt_option` VALUES (10, 'show_on_front_page', 'about', 1);
INSERT INTO `pt_option` VALUES (11, 'posts_per_page', '10', 1);
INSERT INTO `pt_option` VALUES (12, 'open_XML', 'on', 1);
INSERT INTO `pt_option` VALUES (13, 'article_comment_status', 'open', 1);
INSERT INTO `pt_option` VALUES (14, 'page_comment_status', 'open', 1);
INSERT INTO `pt_option` VALUES (15, 'comment_need_register', 'no', 1);
INSERT INTO `pt_option` VALUES (16, 'show_comment_page', 'on', 1);
INSERT INTO `pt_option` VALUES (17, 'comment_per_page', '15', 1);
INSERT INTO `pt_option` VALUES (18, 'comment_page_first', 'last', 1);
INSERT INTO `pt_option` VALUES (19, 'comment_page_top', 'new', 1);
INSERT INTO `pt_option` VALUES (20, 'comment_before_show', 'directly', 1);
INSERT INTO `pt_option` VALUES (21, 'show_avatar', 'on', 1);
INSERT INTO `pt_option` VALUES (22, 'image_thumbnail_width', '150', 0);
INSERT INTO `pt_option` VALUES (23, 'image_thumbnail_height', '150', 0);
INSERT INTO `pt_option` VALUES (24, 'image_medium_width', '300', 0);
INSERT INTO `pt_option` VALUES (25, 'image_medium_height', '300', 0);
INSERT INTO `pt_option` VALUES (26, 'image_large_width', '1024', 0);
INSERT INTO `pt_option` VALUES (27, 'image_large_height', '1024', 0);
INSERT INTO `pt_option` VALUES (28, 'site_description', '一个新的 Puti 站点。', 1);
INSERT INTO `pt_option` VALUES (29, 'site_keywords', '独立博客,Puti,PutiProject', 1);
INSERT INTO `pt_option` VALUES (30, 'footer_copyright', '<p> Copyright © 2017 <a target=\"_blank\" href=\"https://github.com/puti-projects\">Puti</a> All Rights Reserved. Powered by <a href=\"https://github.com/puti-projects/puti\" target=\"_blank\" rel=\"nofollow\">Puti</a></p>', 1);
INSERT INTO `pt_option` VALUES (31, 'show_project', '1', 1);
INSERT INTO `pt_option` VALUES (32, 'github_user', '', 0);
INSERT INTO `pt_option` VALUES (33, 'github_show_repo', '', 0);
INSERT INTO `pt_option` VALUES (34, 'site_language', '简体中文', 1);
INSERT INTO `pt_option` VALUES (35, 'current_theme', 'Emma', 1);

-- ----------------------------
-- Table structure for pt_post
-- ----------------------------
DROP TABLE IF EXISTS `pt_post`;
CREATE TABLE `pt_post`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发表人id',
  `post_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'article' COMMENT '类型：article，page',
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `content_markdown` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'markdown格式文章内容',
  `content_html` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'html格式文章内容',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '缩略名（用于url中展示）',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id（如果有）',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'publish' COMMENT '状态:publish,draft,deleted',
  `comment_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '评论状态(是否开启);默认1开启；0关闭',
  `if_top` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否置顶；1置顶',
  `guid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一链接',
  `cover_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '封面图片链接',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT '评论数目',
  `view_count` int(11) NOT NULL DEFAULT 0 COMMENT '浏览量',
  `posted_time` datetime(0) NULL DEFAULT NULL COMMENT '发表时间(UTC)',
  `created_time` datetime(0) NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `post_parent`(`parent_id`) USING BTREE,
  INDEX `post_author`(`user_id`) USING BTREE,
  INDEX `type_status_date`(`id`, `post_type`, `status`) USING BTREE,
  INDEX `post_name`(`slug`(191)) USING BTREE,
  FULLTEXT INDEX `post_title`(`title`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_post
-- ----------------------------
INSERT INTO `pt_post` VALUES (1, 1, 'article', 'Hellow World!', '这是一篇测试文章。\nThis is a test article.', '<p>这是一篇测试文章。<br />\nThis is a test article.</p>\n', '', 0, 'publish', 1, 0, '/article/1.html', '', 0, 1, '2019-03-20 19:33:54', '2019-03-20 19:34:06', '2019-03-20 19:38:57', NULL);
INSERT INTO `pt_post` VALUES (2, 1, 'page', 'About', '这是一个测试页面.\nThis is a test page.', '<p>这是一个测试页面.<br />\nThis is a test page.</p>\n', 'about-me', 0, 'publish', 1, 0, '/about-me', '', 0, 1, '2019-03-20 19:34:26', '2019-03-20 19:34:51', '2019-03-20 19:38:57', NULL);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_post_meta
-- ----------------------------
INSERT INTO `pt_post_meta` VALUES (1, 1, 'description', '');
INSERT INTO `pt_post_meta` VALUES (2, 2, 'description', '');
INSERT INTO `pt_post_meta` VALUES (3, 2, 'page_template', 'default');

-- ----------------------------
-- Table structure for pt_resource
-- ----------------------------
DROP TABLE IF EXISTS `pt_resource`;
CREATE TABLE `pt_resource`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `upload_user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有者id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '资源名称',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '资源说明',
  `guid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一链接',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'picture' COMMENT '资源类型；默认picture',
  `mime_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '资源文件类型',
  `usage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用途；common普通,cover封面',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '资源状态;1正常',
  `created_time` datetime(0) NOT NULL COMMENT '上传时间(UTC)',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `resource_type`(`id`, `type`, `status`) USING BTREE,
  INDEX `resource_name`(`slug`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '资源表' ROW_FORMAT = Dynamic;

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
-- Table structure for pt_subject
-- ----------------------------
DROP TABLE IF EXISTS `pt_subject`;
CREATE TABLE `pt_subject`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '专题 id',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专题名称',
  `slug` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专题缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `cover_image` int(11) NOT NULL DEFAULT 0 COMMENT '封面图;关联resource',
  `is_end` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否末级；1是 0不是',
  `count` int(11) NOT NULL DEFAULT 0 COMMENT '拥有文章数量',
  `last_updated` datetime(0) NULL DEFAULT NULL COMMENT '上次更新(关联文章)',
  `created_time` datetime(0) NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime(0) NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `subject_slug`(`slug`) USING BTREE,
  INDEX `subkect_parent`(`parent_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '专题表' ROW_FORMAT = Dynamic;

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
-- Table structure for pt_term
-- ----------------------------
DROP TABLE IF EXISTS `pt_term`;
CREATE TABLE `pt_term`  (
  `term_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '条件id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '条件名称',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '缩略名',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有的数目',
  PRIMARY KEY (`term_id`) USING BTREE,
  INDEX `slug`(`slug`(191)) USING BTREE,
  INDEX `name`(`name`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_term
-- ----------------------------
INSERT INTO `pt_term` VALUES (1, '未分类', 'uncategorized', '', 1);

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
-- Records of pt_term_relationships
-- ----------------------------
INSERT INTO `pt_term_relationships` VALUES (1, 1, 0);

-- ----------------------------
-- Table structure for pt_term_taxonomy
-- ----------------------------
DROP TABLE IF EXISTS `pt_term_taxonomy`;
CREATE TABLE `pt_term_taxonomy`  (
  `term_taxonomy_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类方式id',
  `term_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'term_id',
  `parent_term_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父term_id',
  `level` int(11) NOT NULL DEFAULT 1 COMMENT '层级',
  `taxonomy` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类方式',
  `term_group` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分组',
  PRIMARY KEY (`term_taxonomy_id`) USING BTREE,
  UNIQUE INDEX `term_id_taxonomy`(`term_id`, `taxonomy`) USING BTREE,
  INDEX `taxonomy`(`taxonomy`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_term_taxonomy
-- ----------------------------
INSERT INTO `pt_term_taxonomy` VALUES (1, 1, 0, 1, 'category', 0);

-- ----------------------------
-- Table structure for pt_user
-- ----------------------------
DROP TABLE IF EXISTS `pt_user`;
CREATE TABLE `pt_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录帐号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '头像',
  `page_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '主页链接',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '状态.1激活2冻结',
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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_user
-- ----------------------------
INSERT INTO `pt_user` VALUES (1, 'admin', '$2a$10$/cbpwIig1p0ahzSgmVU0auVnuBOx6fzpSaOAXc7nw4VpxhOytiU0i', 'Admin', 'example@example.com', '/assets/users/default.jpg', 'https://www.example.com', 1, 'administrator', '2018-07-24 02:51:38', '2019-03-20 19:21:12', NULL);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
