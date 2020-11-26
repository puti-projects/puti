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
-- Table structure for table `pt_knowledge_item`
--

DROP TABLE IF EXISTS `pt_knowledge_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_knowledge_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `knowledge_id` int NOT NULL,
  `symbol` bigint NOT NULL COMMENT '唯一标识',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '发表人id',
  `title` varchar(512) NOT NULL COMMENT '标题',
  `content_version` bigint NOT NULL DEFAULT '0' COMMENT '指向一个当前版本；对应content表的version；默认0；笔记类型的为0，因为没有多版本',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级id',
  `level` int NOT NULL DEFAULT '0' COMMENT '目录级别',
  `index` int NOT NULL DEFAULT '0' COMMENT '排序值',
  `comment_count` int NOT NULL DEFAULT '0' COMMENT '评论数目',
  `view_count` int NOT NULL DEFAULT '0' COMMENT '浏览量',
  `last_published` datetime DEFAULT NULL COMMENT '上次发布内容时间',
  `created_time` datetime NOT NULL COMMENT '创建时间(UTC)',
  `updated_time` datetime NOT NULL COMMENT '更新时间(UTC)',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间(UTC)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`),
  KEY `knowledge_id` (`knowledge_id`),
  KEY `index` (`parent_id`,`level`,`index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_knowledge_item`
--

LOCK TABLES `pt_knowledge_item` WRITE;
/*!40000 ALTER TABLE `pt_knowledge_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_knowledge_item` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-26 21:51:06
