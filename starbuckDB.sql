-- -------------------------------------------------------------
-- TablePlus 6.0.0(550)
--
-- https://tableplus.com/
--
-- Database: starbuck
-- Generation Time: 2567-05-27 12:57:43.3060
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `caffeine`;
CREATE TABLE `caffeine` (
  `caffeine_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`caffeine_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `category_id` int NOT NULL,
  `name` text,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `price` int DEFAULT NULL,
  `detail_id` int DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  `roast_id` int DEFAULT NULL,
  `caffeine` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `detail_id` (`detail_id`),
  KEY `roast_id` (`roast_id`),
  KEY `caffeine` (`caffeine`),
  CONSTRAINT `items_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`),
  CONSTRAINT `items_ibfk_2` FOREIGN KEY (`detail_id`) REFERENCES `items_detail` (`detail_id`),
  CONSTRAINT `items_ibfk_3` FOREIGN KEY (`roast_id`) REFERENCES `roast` (`roast_id`),
  CONSTRAINT `items_ibfk_4` FOREIGN KEY (`caffeine`) REFERENCES `caffeine` (`caffeine_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `items_detail`;
CREATE TABLE `items_detail` (
  `detail_id` int NOT NULL AUTO_INCREMENT,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `acidity` int NOT NULL DEFAULT '0',
  `body` int NOT NULL DEFAULT '0',
  `processing_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `tasting_notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `complementary_flavors` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  PRIMARY KEY (`detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `roast`;
CREATE TABLE `roast` (
  `roast_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`roast_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `caffeine` (`caffeine_id`, `name`) VALUES
(1, 'Decaf'),
(2, 'Regular');

INSERT INTO `categories` (`category_id`, `name`) VALUES
(1, 'Whole Bean'),
(2, 'Starbucks Reserve'),
(3, 'Starbucks VIA'),
(4, 'Starbucks Origami'),
(5, 'Teavana'),
(6, 'Syrup');

INSERT INTO `items` (`id`, `name`, `quantity`, `price`, `detail_id`, `category_id`, `roast_id`, `caffeine`) VALUES
(1, 'Starbucks® Single-Origin Sun-Dried Brazil Carmo De Minas', 90, 750, 1, 1, 2, 1),
(2, 'Starbucks Reserve™ Colombia Narino Granos De Esperanza™', 30, 850, 2, 2, 1, 1),
(3, 'Starbucks® Single-Origin Papua New Guinea Highlands', 40, 750, 3, 3, 2, 2),
(4, 'Starbucks VIA™ Iced Coffee', 50, 150, 4, 4, 3, 2),
(5, 'Starbucks® Veranda Blend™', 60, 545, 5, 2, 3, 2),
(6, 'Breakfast Blend', 70, 545, 1, 1, 1, 1),
(7, 'Starbucks VIA™ Pike Place® Roast', 80, 350, 2, 5, 2, 2),
(8, 'Starbucks Caffè Verona™ 6 Cups', 90, 350, 3, 2, 2, 2),
(9, 'Starbucks Reserve™ Indonesia Gunung Leuser (Available on 13 May 24 onwards)', 100, 1200, 4, 4, 1, 1),
(10, 'Starbucks VIA™ Italian Roast', 110, 350, 5, 5, 3, 1),
(11, 'Starbucks Pike Place® Roast 6 Cups', 120, 350, 1, 2, 2, 2),
(12, 'Vanilla Flavoured Syrup', 130, 180, 2, 1, 2, 1);

INSERT INTO `items_detail` (`detail_id`, `description`, `acidity`, `body`, `processing_method`, `tasting_notes`, `complementary_flavors`, `region`) VALUES
(1, 'Due to the unique geography and climate conditions of the southeastern region of carmo de minas, the most prominent processing method is sun-dried, which intrinsically carries the least amount of impact to the environment by using very little water. This coffee has a subtle acidity and a softer profile that highlights tropical fruits and sweet hazelnuts with a smooth caramelly texture that is best paired with caramel and chocolate.', 4, 4, 'Natural', 'Hazelnut, Dried fruit', 'Caramel, Chocolate', 'Latin America'),
(2, 'Mock data2', 2, 2, 'Mock data2', 'Mock data2', 'Mock data2', 'Mock data2'),
(3, 'Mock data3', 3, 3, 'Mock data3', 'Mock data3', 'Mock data3', 'Mock data3'),
(4, 'Mock data4', 4, 4, 'Mock data4', 'Mock data4', 'Mock data4', 'Mock data4'),
(5, 'Mock data5', 5, 5, 'Mock data5', 'Mock data5', 'Mock data5', 'Mock data5');

INSERT INTO `roast` (`roast_id`, `name`) VALUES
(1, 'Blonde'),
(2, 'Medium'),
(3, 'Dark');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;