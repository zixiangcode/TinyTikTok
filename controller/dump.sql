-- MySQL dump 10.13  Distrib 8.0.34, for Linux (x86_64)
--
-- Host: localhost    Database: TikTok
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `video_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `content` longtext,
  `create_date` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_comments_user` (`user_id`),
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,6,4,'Comment 1','2023-08-08 00:10:33'),(2,7,15,'Comment 2','2023-08-08 00:10:33'),(3,10,1,'Comment 3','2023-08-08 00:10:33'),(4,5,20,'Comment 4','2023-08-08 00:10:33'),(5,10,15,'Comment 5','2023-08-08 00:10:33'),(6,2,6,'Comment 6','2023-08-08 00:10:33'),(7,8,1,'Comment 7','2023-08-08 00:10:33'),(8,3,21,'Comment 8','2023-08-08 00:10:33'),(9,4,30,'Comment 9','2023-08-08 00:10:33'),(10,10,18,'Comment 10','2023-08-08 00:10:33'),(11,2,30,'Comment 11','2023-08-08 00:10:33'),(12,6,17,'Comment 12','2023-08-08 00:10:33'),(13,7,26,'Comment 13','2023-08-08 00:10:33'),(14,2,25,'Comment 14','2023-08-08 00:10:33'),(15,10,9,'Comment 15','2023-08-08 00:10:33'),(16,10,29,'Comment 16','2023-08-08 00:10:33'),(17,8,21,'Comment 17','2023-08-08 00:10:33'),(18,7,17,'Comment 18','2023-08-08 00:10:33'),(19,3,13,'Comment 19','2023-08-08 00:10:33'),(20,3,24,'Comment 20','2023-08-08 00:10:33'),(21,7,4,'Comment 21','2023-08-08 00:10:33'),(22,7,28,'Comment 22','2023-08-08 00:10:33'),(23,8,17,'Comment 23','2023-08-08 00:10:33'),(24,3,3,'Comment 24','2023-08-08 00:10:33'),(25,3,24,'Comment 25','2023-08-08 00:10:33'),(26,8,27,'Comment 26','2023-08-08 00:10:33'),(27,6,29,'Comment 27','2023-08-08 00:10:33'),(28,8,4,'Comment 28','2023-08-08 00:10:33'),(29,7,18,'Comment 29','2023-08-08 00:10:33'),(30,7,9,'Comment 30','2023-08-08 00:10:33'),(31,2,7,'Comment 31','2023-08-08 00:10:33'),(32,2,11,'Comment 32','2023-08-08 00:10:33'),(33,5,17,'Comment 33','2023-08-08 00:10:33'),(34,10,14,'Comment 34','2023-08-08 00:10:33'),(35,1,12,'Comment 35','2023-08-08 00:10:33'),(36,8,6,'Comment 36','2023-08-08 00:10:33'),(37,3,10,'Comment 37','2023-08-08 00:10:33'),(38,4,13,'Comment 38','2023-08-08 00:10:33'),(39,5,9,'Comment 39','2023-08-08 00:10:33'),(40,1,15,'Comment 40','2023-08-08 00:10:33'),(41,3,20,'Comment 41','2023-08-08 00:10:33'),(42,9,4,'Comment 42','2023-08-08 00:10:33'),(43,5,1,'Comment 43','2023-08-08 00:10:33'),(44,4,5,'Comment 44','2023-08-08 00:10:33'),(45,6,17,'Comment 45','2023-08-08 00:10:33'),(46,9,10,'Comment 46','2023-08-08 00:10:33'),(47,3,5,'Comment 47','2023-08-08 00:10:33'),(48,8,28,'Comment 48','2023-08-08 00:10:33'),(49,2,26,'Comment 49','2023-08-08 00:10:33'),(50,6,7,'Comment 50','2023-08-08 00:10:33'),(51,9,10,'Comment 51','2023-08-08 00:10:33'),(52,3,12,'Comment 52','2023-08-08 00:10:33'),(53,5,1,'Comment 53','2023-08-08 00:10:33'),(54,7,30,'Comment 54','2023-08-08 00:10:33'),(55,4,15,'Comment 55','2023-08-08 00:10:33'),(56,5,9,'Comment 56','2023-08-08 00:10:33'),(57,7,6,'Comment 57','2023-08-08 00:10:33'),(58,1,25,'Comment 58','2023-08-08 00:10:33'),(59,7,30,'Comment 59','2023-08-08 00:10:33'),(60,6,9,'Comment 60','2023-08-08 00:10:33'),(61,9,21,'Comment 61','2023-08-08 00:10:33'),(62,10,3,'Comment 62','2023-08-08 00:10:33'),(63,5,21,'Comment 63','2023-08-08 00:10:33'),(64,5,12,'Comment 64','2023-08-08 00:10:33'),(65,3,27,'Comment 65','2023-08-08 00:10:33'),(66,3,11,'Comment 66','2023-08-08 00:10:33'),(67,2,29,'Comment 67','2023-08-08 00:10:33'),(68,3,22,'Comment 68','2023-08-08 00:10:33'),(69,3,7,'Comment 69','2023-08-08 00:10:33'),(70,2,18,'Comment 70','2023-08-08 00:10:33'),(71,2,28,'Comment 71','2023-08-08 00:10:33'),(72,1,4,'Comment 72','2023-08-08 00:10:33'),(73,3,4,'Comment 73','2023-08-08 00:10:33'),(74,3,1,'Comment 74','2023-08-08 00:10:33'),(75,10,5,'Comment 75','2023-08-08 00:10:33'),(76,5,15,'Comment 76','2023-08-08 00:10:33'),(77,7,2,'Comment 77','2023-08-08 00:10:33'),(78,4,9,'Comment 78','2023-08-08 00:10:33'),(79,7,29,'Comment 79','2023-08-08 00:10:33'),(80,8,11,'Comment 80','2023-08-08 00:10:33'),(81,5,29,'Comment 81','2023-08-08 00:10:33'),(82,3,10,'Comment 82','2023-08-08 00:10:33'),(83,3,13,'Comment 83','2023-08-08 00:10:33'),(84,8,27,'Comment 84','2023-08-08 00:10:33'),(85,2,3,'Comment 85','2023-08-08 00:10:33'),(86,8,14,'Comment 86','2023-08-08 00:10:33'),(87,8,19,'Comment 87','2023-08-08 00:10:33'),(88,4,28,'Comment 88','2023-08-08 00:10:33'),(89,3,3,'Comment 89','2023-08-08 00:10:33'),(90,6,2,'Comment 90','2023-08-08 00:10:33'),(91,8,14,'Comment 91','2023-08-08 00:10:33'),(92,1,8,'Comment 92','2023-08-08 00:10:33'),(93,2,14,'Comment 93','2023-08-08 00:10:33'),(94,8,7,'Comment 94','2023-08-08 00:10:33'),(95,6,29,'Comment 95','2023-08-08 00:10:33'),(96,2,29,'Comment 96','2023-08-08 00:10:33'),(97,2,28,'Comment 97','2023-08-08 00:10:33'),(98,6,24,'Comment 98','2023-08-08 00:10:33'),(99,4,19,'Comment 99','2023-08-08 00:10:33'),(128,1,1,'mytest01\n','08-08');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL,
  `name` longtext,
  `follow_count` bigint DEFAULT NULL,
  `follower_count` bigint DEFAULT NULL,
  `is_follow` tinyint(1) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `background_image` varchar(255) DEFAULT NULL,
  `signature` text,
  `total_favorited` varchar(255) DEFAULT NULL,
  `work_count` int DEFAULT NULL,
  `favorite_count` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'User1',10,20,1,'avatar_url','background_url','User1 signature','100',5,50),(2,'User2',15,25,1,'avatar_url2','background_url2','User2 signature','200',8,60),(3,'User3',8,30,0,'avatar_url3','background_url3','User3 signature','50',3,25),(4,'User4',12,18,1,'avatar_url4','background_url4','User4 signature','80',7,40),(5,'User5',25,35,0,'avatar_url5','background_url5','User5 signature','150',10,80),(6,'User6',20,40,1,'avatar_url6','background_url6','User6 signature','120',9,65),(7,'User7',30,50,1,'avatar_url7','background_url7','User7 signature','200',12,90),(8,'User8',18,22,0,'avatar_url8','background_url8','User8 signature','75',6,30),(9,'User9',40,60,1,'avatar_url9','background_url9','User9 signature','250',15,100),(10,'User10',15,28,1,'avatar_url10','background_url10','User10 signature','90',5,45),(11,'User11',22,42,0,'avatar_url11','background_url11','User11 signature','130',8,70),(12,'User12',32,55,1,'avatar_url12','background_url12','User12 signature','180',11,85),(13,'User13',50,70,1,'avatar_url13','background_url13','User13 signature','300',20,120),(14,'User14',16,20,0,'avatar_url14','background_url14','User14 signature','70',4,25),(15,'User15',45,65,1,'avatar_url15','background_url15','User15 signature','220',14,110),(16,'User16',17,30,1,'avatar_url16','background_url16','User16 signature','95',6,35),(17,'User17',28,48,0,'avatar_url17','background_url17','User17 signature','140',9,75),(18,'User18',35,58,1,'avatar_url18','background_url18','User18 signature','190',10,95),(19,'User19',42,62,1,'avatar_url19','background_url19','User19 signature','210',12,105),(20,'User20',19,25,0,'avatar_url20','background_url20','User20 signature','85',5,40),(21,'User21',52,75,1,'avatar_url21','background_url21','User21 signature','280',18,115),(22,'User22',20,32,1,'avatar_url22','background_url22','User22 signature','100',7,50),(23,'User23',38,52,0,'avatar_url23','background_url23','User23 signature','170',10,85),(24,'User24',55,78,1,'avatar_url24','background_url24','User24 signature','320',22,135),(25,'User25',24,38,1,'avatar_url25','background_url25','User25 signature','110',8,55),(26,'User26',33,55,0,'avatar_url26','background_url26','User26 signature','175',9,90),(27,'User27',40,60,1,'avatar_url27','background_url27','User27 signature','230',13,110),(28,'User28',21,29,1,'avatar_url28','background_url28','User28 signature','95',6,45),(29,'User29',48,68,0,'avatar_url29','background_url29','User29 signature','260',17,100),(30,'User30',25,35,1,'avatar_url30','background_url30','User30 signature','120',8,60);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-08  1:35:16
