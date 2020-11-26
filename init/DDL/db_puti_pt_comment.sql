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
-- Table structure for table `pt_comment`
--

DROP TABLE IF EXISTS `pt_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父评论id',
  `post_id` int unsigned NOT NULL DEFAULT '0' COMMENT '评论的文章或页面id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `if_visitor` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否游客;1是,0不是;默认游客',
  `commenter_user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '评论者id;是游客时为0;默认为0',
  `commenter_name` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论者名称',
  `commenter_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '评论者email',
  `commenter_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '评论者链接',
  `commenter_ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '评论者ip',
  `comment_date` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论时间(UTC)',
  `approved` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '1' COMMENT '是否通过(开启评论审核后，通过后显示)',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '评论来源agent',
  `created_time` datetime NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `comment_post_ID` (`post_id`) USING BTREE,
  KEY `comment_parent` (`parent_id`) USING BTREE,
  KEY `comment_author_email` (`commenter_email`(10)) USING BTREE,
  KEY `comment_approved_date` (`comment_date`,`approved`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_comment`
--

LOCK TABLES `pt_comment` WRITE;
/*!40000 ALTER TABLE `pt_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_comment` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-26 21:51:11
