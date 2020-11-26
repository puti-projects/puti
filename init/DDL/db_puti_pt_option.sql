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
-- Table structure for table `pt_option`
--

DROP TABLE IF EXISTS `pt_option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_option` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '配置id',
  `option_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置名称',
  `option_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '对应的值',
  `autoload` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否自动加载;默认0不自动加载',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `option_name` (`option_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_option`
--

LOCK TABLES `pt_option` WRITE;
/*!40000 ALTER TABLE `pt_option` DISABLE KEYS */;
INSERT INTO `pt_option` VALUES (1,'blog_name','gogogo',1),(2,'blog_description','一个新的 Puti 站点',1),(3,'site_url','http://puti.com',1),(4,'admin_email','example@example.com',1),(5,'users_can_register','off',1),(6,'timezone_string','Asia/Shanghai',1),(7,'default_category','1',0),(8,'default_link_category','0',0),(9,'show_on_front','article',1),(10,'show_on_front_page','about',1),(11,'posts_per_page','10',1),(12,'open_XML','on',1),(13,'article_comment_status','open',1),(14,'page_comment_status','open',1),(15,'comment_need_register','no',1),(16,'show_comment_page','on',1),(17,'comment_per_page','15',1),(18,'comment_page_first','last',1),(19,'comment_page_top','new',1),(20,'comment_before_show','directly',1),(21,'show_avatar','on',1),(22,'image_thumbnail_width','150',0),(23,'image_thumbnail_height','150',0),(24,'image_medium_width','300',0),(25,'image_medium_height','300',0),(26,'image_large_width','1024',0),(27,'image_large_height','1024',0),(28,'site_description','一个新的 Puti 站点。',1),(29,'site_keywords','独立博客,Puti,PutiProject',1),(30,'footer_copyright','<p> Copyright © 2017 <a target=\"_blank\" href=\"https://github.com/puti-projects\">Puti</a> All Rights Reserved. Powered by <a href=\"https://github.com/puti-projects/puti\" target=\"_blank\" rel=\"nofollow\">Puti</a></p>',1),(31,'show_project','1',1),(32,'github_user','',0),(33,'github_show_repo','',0),(34,'site_language','简体中文',1),(35,'current_theme','Lin',1);
/*!40000 ALTER TABLE `pt_option` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-26 21:50:56
