-- MySQL dump 10.13  Distrib 5.7.26, for macos10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: TUserDB
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(32) DEFAULT '' COMMENT '英文名',
  `realname` char(11) DEFAULT '' COMMENT '姓名',
  `email` char(64) DEFAULT '' COMMENT '邮箱',
  `mobile` char(11) DEFAULT '' COMMENT '手机',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `verify` char(32) DEFAULT '' COMMENT '密钥',
  `entry_date` bigint(20) DEFAULT '0' COMMENT '入职时间',
  `op_role` tinyint(11) DEFAULT '0' COMMENT '后台用户',
  `admin_role` tinyint(4) DEFAULT '0' COMMENT '管理员',
  `email_flag` tinyint(11) DEFAULT '0' COMMENT '邮件验证状态',
  `mobile_flag` tinyint(11) DEFAULT '0' COMMENT '手机验证状态',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`,`realname`,`email`,`mobile`,`status`,`op_role`,`mobile_flag`,`email_flag`),
  KEY `update_at` (`updated_at`),
  KEY `deleted_at` (`deleted_at`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'liushuojia','刘','liushuojia@qq.com','13725588389',1,'001',1,1,1,1,1,'2019-09-29 17:16:57',NULL,NULL),(2,'apple','陈洁萍','apple@qq.com','13427502964',1,'002',0,0,1,0,0,'2019-09-29 17:16:57',NULL,NULL),(4,'liushuojia','刘硕嘉','liushuojia@qq.com','13725588389',1,'001',1,1,1,1,1,'2019-09-29 17:16:57',NULL,'2019-09-30 02:20:58'),(7,'刘硕嘉001','','49650719@qq.com','13725588389',1,'',0,0,0,0,0,'2019-09-30 05:35:16','2019-09-30 05:35:16',NULL),(8,'liushuojia','刘硕嘉','49650719@qq.com','13725588389',1,'',0,0,0,0,0,'2019-09-30 05:35:54','2019-09-30 05:35:54',NULL),(9,'liushuojia','刘硕嘉','49650719@qq.com','13725588389',1,'',0,0,0,0,0,'2019-09-30 05:36:30','2019-09-30 05:36:30',NULL),(10,'liushuojia','刘硕嘉','49650719@qq.com','13725588389',1,'75e91ed7b5ac6650444b93ec19ef2761',0,0,0,0,0,'2019-09-30 05:39:36','2019-09-30 05:39:36','2019-09-30 05:52:59'),(12,'liushuojia12222','刘硕嘉2','49650719@qq.com3','137255811',1,'D617697DE621B2C53EAFCD341A4705C3',6,7,8,9,10,'2019-09-30 06:02:07','2019-10-03 08:50:59',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userExt`
--

DROP TABLE IF EXISTS `userExt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `userExt` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(11) NOT NULL COMMENT '用户id',
  `desc` text COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userExt`
--

LOCK TABLES `userExt` WRITE;
/*!40000 ALTER TABLE `userExt` DISABLE KEYS */;
/*!40000 ALTER TABLE `userExt` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-03 17:35:53
