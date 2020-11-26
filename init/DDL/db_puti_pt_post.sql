-- MySQL dump 10.13  Distrib 8.0.22, for macos10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: db_puti
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `pt_post`
--

DROP TABLE IF EXISTS `pt_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_post` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '发表人id',
  `post_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'article' COMMENT '类型：article，page',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
  `content_markdown` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'markdown格式文章内容',
  `content_html` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'html格式文章内容',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '缩略名（用于url中展示）',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父id（如果有）',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'publish' COMMENT '状态:publish,draft,deleted',
  `comment_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '评论状态(是否开启);默认1开启；0关闭',
  `if_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶；1置顶',
  `guid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '唯一链接',
  `cover_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面图片链接',
  `comment_count` int NOT NULL DEFAULT '0' COMMENT '评论数目',
  `view_count` int NOT NULL DEFAULT '0' COMMENT '浏览量',
  `posted_time` datetime DEFAULT NULL COMMENT '发表时间(UTC)',
  `created_time` datetime NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `post_parent` (`parent_id`) USING BTREE,
  KEY `post_author` (`user_id`) USING BTREE,
  KEY `type_status_date` (`id`,`post_type`,`status`) USING BTREE,
  KEY `post_name` (`slug`(191)) USING BTREE,
  FULLTEXT KEY `post_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_post`
--

LOCK TABLES `pt_post` WRITE;
/*!40000 ALTER TABLE `pt_post` DISABLE KEYS */;
INSERT INTO `pt_post` VALUES (1,1,'article','Hellow World!','这是一篇测试文章。\nThis is a test article.','<p>这是一篇测试文章。<br />\nThis is a test article.</p>\n','',0,'publish',1,0,'/article/1.html','',0,0,'2019-03-20 19:33:54','2019-03-20 19:34:06','2020-11-21 16:19:55',NULL),(2,1,'page','About','这是一个测试页面. This is a test page.','<p>这是一个测试页面.<br /> This is a test page.</p> ','about-me',0,'publish',1,0,'/about-me','',0,0,'2019-03-20 19:34:26','2019-03-20 19:34:51','2020-11-21 19:01:26',NULL);
/*!40000 ALTER TABLE `pt_post` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-26 21:51:07
