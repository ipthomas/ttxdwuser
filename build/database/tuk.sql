-- MySQL dump 10.13  Distrib 8.0.33, for macos13.3 (x86_64)
--
-- Host: tuk.coil1nnpqdlr.eu-west-1.rds.amazonaws.com    Database: tuk
-- ------------------------------------------------------
-- Server version	8.0.33

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
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '';

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `creationtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `eventtype` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `docname` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `classcode` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `confcode` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `formatcode` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `facilitycode` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `practicecode` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `speciality` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `expression` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `authors` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `xdsdocentryuid` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `repositoryuid` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `nhs` varchar(10) NOT NULL DEFAULT 'None',
  `user` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `org` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `role` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `topic` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `pathway` varchar(256) NOT NULL DEFAULT 'Not Provided',
  `comments` varchar(4000) NOT NULL DEFAULT 'None',
  `version` int NOT NULL DEFAULT '0',
  `taskid` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `nhsid` (`nhs`),
  KEY `user` (`user`),
  KEY `pathway` (`pathway`),
  KEY `expression` (`expression`),
  KEY `xdsdocentryuId` (`xdsdocentryuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2509 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `idmaps`
--

DROP TABLE IF EXISTS `idmaps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `idmaps` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `lid` varchar(256) DEFAULT NULL,
  `mid` varchar(256) DEFAULT NULL,
  `user` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `statics`
--

DROP TABLE IF EXISTS `statics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `statics` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `content` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=331 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subscriptions`
--

DROP TABLE IF EXISTS `subscriptions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subscriptions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `brokerref` varchar(256) DEFAULT '',
  `pathway` varchar(256) DEFAULT '',
  `topic` varchar(256) DEFAULT '',
  `expression` varchar(256) DEFAULT '',
  `email` varchar(100) DEFAULT '',
  `nhsid` varchar(100) DEFAULT '',
  `user` varchar(100) DEFAULT '',
  `org` varchar(100) DEFAULT '',
  `role` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `brokerref` (`brokerref`),
  KEY `pathway` (`pathway`)
) ENGINE=InnoDB AUTO_INCREMENT=1369 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `templates`
--

DROP TABLE IF EXISTS `templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `templates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `template` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `user` varchar(100) DEFAULT 'system',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34151 DEFAULT CHARSET=utf8mb3 COMMENT='newTable';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `workflows`
--

DROP TABLE IF EXISTS `workflows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `workflows` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `pathway` varchar(128) NOT NULL,
  `nhsid` varchar(10) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `xdw_key` varchar(256) NOT NULL DEFAULT '',
  `xdw_uid` varchar(255) NOT NULL DEFAULT '',
  `xdw_doc` longtext NOT NULL,
  `xdw_def` longtext NOT NULL,
  `version` int NOT NULL DEFAULT '0',
  `published` tinyint(1) NOT NULL DEFAULT '0',
  `status` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3009 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `xdws`
--

DROP TABLE IF EXISTS `xdws`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `xdws` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `isxdsmeta` tinyint(1) NOT NULL DEFAULT '1',
  `xdw` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=422 DEFAULT CHARSET=utf8mb3 COMMENT='newTable';
/*!40101 SET character_set_client = @saved_cs_client */;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-28 17:09:19
