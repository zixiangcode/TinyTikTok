/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : tinytiktok

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 03/09/2023 23:45:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `video_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `is_deleted` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_comments_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 153 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '6', 4, 'Comment 1', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (2, '7', 15, 'Comment 2', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (3, '10', 1, 'Comment 3', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (4, '5', 20, 'Comment 4', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (5, '10', 15, 'Comment 5', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (6, '2', 6, 'Comment 6', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (7, '8', 1, 'Comment 7', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (8, '3', 21, 'Comment 8', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (9, '4', 30, 'Comment 9', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (10, '10', 18, 'Comment 10', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (11, '2', 30, 'Comment 11', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (12, '6', 17, 'Comment 12', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (13, '7', 26, 'Comment 13', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (14, '2', 25, 'Comment 14', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (15, '10', 9, 'Comment 15', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (16, '10', 29, 'Comment 16', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (17, '8', 21, 'Comment 17', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (18, '7', 17, 'Comment 18', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (19, '3', 13, 'Comment 19', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (20, '3', 24, 'Comment 20', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (21, '7', 4, 'Comment 21', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (22, '7', 28, 'Comment 22', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (23, '8', 17, 'Comment 23', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (24, '3', 3, 'Comment 24', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (25, '3', 24, 'Comment 25', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (26, '8', 27, 'Comment 26', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (27, '6', 29, 'Comment 27', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (28, '8', 4, 'Comment 28', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (29, '7', 18, 'Comment 29', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (30, '7', 9, 'Comment 30', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (31, '2', 7, 'Comment 31', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (32, '2', 11, 'Comment 32', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (33, '5', 17, 'Comment 33', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (34, '10', 14, 'Comment 34', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (35, '1', 12, 'Comment 35', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (36, '8', 6, 'Comment 36', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (37, '3', 10, 'Comment 37', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (38, '4', 13, 'Comment 38', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (39, '5', 9, 'Comment 39', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (40, '1', 15, 'Comment 40', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (41, '3', 20, 'Comment 41', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (42, '9', 4, 'Comment 42', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (43, '5', 1, 'Comment 43', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (44, '4', 5, 'Comment 44', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (45, '6', 17, 'Comment 45', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (46, '9', 10, 'Comment 46', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (47, '3', 5, 'Comment 47', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (48, '8', 28, 'Comment 48', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (49, '2', 26, 'Comment 49', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (50, '6', 7, 'Comment 50', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (51, '9', 10, 'Comment 51', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (52, '3', 12, 'Comment 52', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (53, '5', 1, 'Comment 53', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (54, '7', 30, 'Comment 54', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (55, '4', 15, 'Comment 55', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (56, '5', 9, 'Comment 56', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (57, '7', 6, 'Comment 57', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (58, '1', 25, 'Comment 58', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (59, '7', 30, 'Comment 59', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (60, '6', 9, 'Comment 60', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (61, '9', 21, 'Comment 61', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (62, '10', 3, 'Comment 62', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (63, '5', 21, 'Comment 63', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (64, '5', 12, 'Comment 64', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (65, '3', 27, 'Comment 65', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (66, '3', 11, 'Comment 66', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (67, '2', 29, 'Comment 67', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (68, '3', 22, 'Comment 68', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (69, '3', 7, 'Comment 69', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (70, '2', 18, 'Comment 70', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (71, '2', 28, 'Comment 71', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (72, '1', 4, 'Comment 72', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (73, '3', 4, 'Comment 73', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (74, '3', 1, 'Comment 74', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (75, '10', 5, 'Comment 75', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (76, '5', 15, 'Comment 76', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (77, '7', 2, 'Comment 77', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (78, '4', 9, 'Comment 78', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (79, '7', 29, 'Comment 79', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (80, '8', 11, 'Comment 80', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (81, '5', 29, 'Comment 81', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (82, '3', 10, 'Comment 82', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (83, '3', 13, 'Comment 83', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (84, '8', 27, 'Comment 84', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (85, '2', 3, 'Comment 85', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (86, '8', 14, 'Comment 86', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (87, '8', 19, 'Comment 87', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (88, '4', 28, 'Comment 88', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (89, '3', 3, 'Comment 89', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (90, '6', 2, 'Comment 90', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (91, '8', 14, 'Comment 91', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (92, '1', 8, 'Comment 92', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (93, '2', 14, 'Comment 93', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (94, '8', 7, 'Comment 94', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (95, '6', 29, 'Comment 95', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (96, '2', 29, 'Comment 96', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (97, '2', 28, 'Comment 97', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (98, '6', 24, 'Comment 98', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (99, '4', 19, 'Comment 99', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (128, '1', 1, 'mytest01', '2023-08-17 00:10:33', 0);
INSERT INTO `comments` VALUES (130, '1', 7097830770726666240, 'mytest03', '2023-08-17 15:00:38', 0);
INSERT INTO `comments` VALUES (131, '1', 7097830770726666240, 'mytest04', '2023-08-17 15:01:07', 1);
INSERT INTO `comments` VALUES (132, '1', 7097830770726666240, 'mytest05', '2023-08-17 15:04:33', 0);
INSERT INTO `comments` VALUES (133, '1', 7097830770726666240, 'mytest05', '2023-08-17 15:17:30', 1);
INSERT INTO `comments` VALUES (134, '1', 7097830770726666240, 'mytest06', '2023-08-17 15:43:32', 1);
INSERT INTO `comments` VALUES (135, '1', 7097830770726666240, 'mytest07', '2023-08-17 15:51:15', 1);
INSERT INTO `comments` VALUES (136, '1', 7097847496482750464, 'mytest08', '2023-08-17 15:53:04', 1);
INSERT INTO `comments` VALUES (137, '1', 7097847496482750464, 'mytest08', '2023-08-17 17:00:50', 0);
INSERT INTO `comments` VALUES (138, '1', 7097847496482750464, 'mytest 09', '2023-08-17 17:00:59', 0);
INSERT INTO `comments` VALUES (139, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:16', 0);
INSERT INTO `comments` VALUES (141, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:35', 1);
INSERT INTO `comments` VALUES (142, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:35', 1);
INSERT INTO `comments` VALUES (143, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:35', 1);
INSERT INTO `comments` VALUES (144, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:35', 1);
INSERT INTO `comments` VALUES (145, '1', 7097830770726666240, 'test09', '2023-08-17 17:09:45', 1);
INSERT INTO `comments` VALUES (146, '1', 7097830770726666240, 'test10', '2023-08-17 17:10:14', 0);
INSERT INTO `comments` VALUES (147, '1', 7097867658552410112, 'test0011', '2023-08-17 17:14:13', 0);
INSERT INTO `comments` VALUES (148, '1', 7097877640291287040, 'test007', '2023-08-17 17:51:56', 1);
INSERT INTO `comments` VALUES (149, '1', 7097876434839928832, 'test0033', '2023-08-17 17:54:42', 1);
INSERT INTO `comments` VALUES (150, '1', 7097847496482750464, 'mytest 0900', '2023-08-18 14:36:14', 0);
INSERT INTO `comments` VALUES (151, '1', 7098191152603463680, 'wodevideo', '2023-08-18 14:38:26', 1);
INSERT INTO `comments` VALUES (152, '1', 7100873833145434112, '1', '2023-08-26 00:17:45', 0);

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
  `user_id` bigint NOT NULL,
  `follow_user_id` bigint NOT NULL,
  `id` bigint NOT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `is_deleted` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follow
-- ----------------------------
INSERT INTO `follow` VALUES (7100873833145434112, 20, 7103735333728026624, '2023-09-02 21:48:05', 0);
INSERT INTO `follow` VALUES (7103759082837770240, 14, 7103762772566999040, '2023-09-02 23:37:07', 0);
INSERT INTO `follow` VALUES (7103759082837770240, 13, 7103762786366259200, '2023-09-02 23:37:10', 0);
INSERT INTO `follow` VALUES (7103759082837770240, 12, 7103764440469733376, '2023-09-02 23:43:44', 0);
INSERT INTO `follow` VALUES (7103759082837770240, 19, 7104123406395637760, '2023-09-03 23:30:08', 1);
INSERT INTO `follow` VALUES (7103759082837770240, 19, 7104125064269791232, '2023-09-03 23:36:44', 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `follow_count` bigint NULL DEFAULT NULL,
  `follower_count` bigint NULL DEFAULT NULL,
  `is_follow` tinyint(1) NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `signature` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `total_favorited` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `work_count` int NULL DEFAULT NULL,
  `favorite_count` int NULL DEFAULT NULL,
  `is_deleted` int NOT NULL DEFAULT 0,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '123456',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'User1', 10, 21, 0, 'avatar_url', 'background_url', 'User1 signature', '100', 5, 50, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (2, 'User2', 15, 25, 0, 'avatar_url2', 'background_url2', 'User2 signature', '200', 8, 60, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (3, 'User3', 8, 30, 0, 'avatar_url3', 'background_url3', 'User3 signature', '50', 3, 25, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (4, 'User4', 12, 18, 0, 'avatar_url4', 'background_url4', 'User4 signature', '80', 7, 40, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (5, 'User5', 25, 35, 0, 'avatar_url5', 'background_url5', 'User5 signature', '150', 10, 80, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (6, 'User6', 20, 40, 0, 'avatar_url6', 'background_url6', 'User6 signature', '120', 9, 65, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (7, 'User7', 30, 50, 0, 'avatar_url7', 'background_url7', 'User7 signature', '200', 12, 90, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (8, 'User8', 18, 22, 0, 'avatar_url8', 'background_url8', 'User8 signature', '75', 6, 30, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (9, 'User9', 40, 60, 0, 'avatar_url9', 'background_url9', 'User9 signature', '250', 15, 100, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (10, 'User10', 15, 28, 0, 'avatar_url10', 'background_url10', 'User10 signature', '90', 5, 45, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (11, 'User11', 22, 42, 0, 'avatar_url11', 'background_url11', 'User11 signature', '130', 8, 70, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (12, 'User12', 32, 56, 1, 'avatar_url12', 'background_url12', 'User12 signature', '180', 11, 85, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (13, 'User13', 50, 71, 1, 'avatar_url13', 'background_url13', 'User13 signature', '300', 20, 120, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (14, 'User14', 16, 21, 1, 'avatar_url14', 'background_url14', 'User14 signature', '70', 4, 25, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (15, 'User15', 45, 65, 0, 'avatar_url15', 'background_url15', 'User15 signature', '220', 14, 110, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (16, 'User16', 17, 30, 0, 'avatar_url16', 'background_url16', 'User16 signature', '95', 6, 35, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (17, 'User17', 28, 48, 0, 'avatar_url17', 'background_url17', 'User17 signature', '140', 9, 75, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (18, 'User18', 35, 58, 0, 'avatar_url18', 'background_url18', 'User18 signature', '190', 10, 95, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (19, 'User19', 42, 62, 0, 'avatar_url19', 'background_url19', 'User19 signature', '210', 12, 105, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (20, 'User20', 19, 27, 0, 'avatar_url20', 'background_url20', 'User20 signature', '85', 5, 40, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (21, 'User21', 52, 75, 0, 'avatar_url21', 'background_url21', 'User21 signature', '280', 18, 115, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (22, 'User22', 20, 32, 0, 'avatar_url22', 'background_url22', 'User22 signature', '100', 7, 50, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (23, 'User23', 38, 52, 0, 'avatar_url23', 'background_url23', 'User23 signature', '170', 10, 85, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (24, 'User24', 55, 78, 0, 'avatar_url24', 'background_url24', 'User24 signature', '320', 22, 135, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (25, 'User25', 24, 38, 0, 'avatar_url25', 'background_url25', 'User25 signature', '110', 8, 55, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (26, 'User26', 33, 55, 0, 'avatar_url26', 'background_url26', 'User26 signature', '175', 9, 90, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (27, 'User27', 40, 60, 0, 'avatar_url27', 'background_url27', 'User27 signature', '230', 13, 110, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (28, 'User28', 21, 29, 0, 'avatar_url28', 'background_url28', 'User28 signature', '95', 6, 45, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (29, 'User29', 48, 68, 0, 'avatar_url29', 'background_url29', 'User29 signature', '260', 17, 100, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (30, 'User30', 25, 35, 0, 'avatar_url30', 'background_url30', 'User30 signature', '120', 8, 60, 0, '2023-08-17 06:15:45', '123456');
INSERT INTO `user` VALUES (7097830770726666240, 'User31', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-17 21:45:27', '$2a$10$6DREfmPqegOaWwTmx0LwfO8JJp3koS.GkCqW.SP60i.u9bCJT07Ya');
INSERT INTO `user` VALUES (7097847496482750464, 'User32', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-17 22:51:55', '$2a$10$bLZlQRmTShGGyUrlnbG6wu.RIHg58Y5CMhqNnJUw7xuX7dY1l2InG');
INSERT INTO `user` VALUES (7097862145206910976, 'User34', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-17 23:50:07', '$2a$10$84AtF6vZFFqfYLGUnb0/k.tWu9wiWMyyQdR..QlEZg8GHp2SxUD4e');
INSERT INTO `user` VALUES (7097863168206045184, 'User33', 20, 21, 0, '', '', '', '55', 55, 44, 0, '2023-08-17 23:54:11', '$2a$10$4Mf/iZPjcQ9PkLjuSSnOieQNO4seMSeC/Z7SxAM.iJOY77B9UgDbO');
INSERT INTO `user` VALUES (7097867658552410112, 'user001', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 00:12:02', '$2a$10$3TEEp112jDvDL3N10pOOlu2F5Pv3md1ubWZF8uwDliXuWwsCr.T5i');
INSERT INTO `user` VALUES (7097868890591461376, 'User35', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 00:16:56', '$2a$10$./mA1FZRSH93BcsGnNE68eE1o8z9gzzD2izwFU0uJPGDSeYRyK38C');
INSERT INTO `user` VALUES (7097876051899973632, 'user002', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 00:45:23', '$2a$10$2wye4LoP8x8Qk3Ff6Xx8qe0pBWibcKiubxzLKx3lOJt8uHlJihET.');
INSERT INTO `user` VALUES (7097876434839928832, 'user003', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 00:46:54', '$2a$10$MrJX13gw8HEJw4k54xHRYup1zOfzd9LCAVM7j.Y8aEF8dQAPhGwiO');
INSERT INTO `user` VALUES (7097877640291287040, 'user004', 0, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 00:51:42', '$2a$10$Ugw4VGXUYz6PuhkjH8U4Cu3khUcSg.vW9Zv/F9XU6ON8wAFDZApka');
INSERT INTO `user` VALUES (7098191152603463680, 'user005', 0, 1, 0, '', '', '', '0', 0, 0, 0, '2023-08-18 21:37:29', '$2a$10$Ma/oO3FoAGevdJESwlwtcOmvx.G/n8LIwg7wn3vjyNWsTWh3T828m');
INSERT INTO `user` VALUES (7099924053208596480, 'wangjinyin', 1, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-23 09:23:24', '$2a$10$N9M.AKhyghw965PjKf3ITeggU0e70L/ZF6S1p1sBmC997l33iKTAW');
INSERT INTO `user` VALUES (7100873833145434112, 'test2', 3, 0, 0, '', '', '', '0', 0, 0, 0, '2023-08-26 00:17:30', '$2a$10$nI42S3TSVHHmk.pnY6zvjuTQ6E9urRRfQCYq0waYnW2hV2CAggr.q');
INSERT INTO `user` VALUES (7103759082837770240, 'test3', 3, 0, 0, '', '', '', '0', 0, 0, 0, '2023-09-02 23:22:27', '$2a$10$Bpqb2MU8gIzK.JB0fikQPOYCy.Uflmv6cgE618EHx7iZL5B9bSZq.');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `author_id` bigint NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, '2023-08-18 12:00:00', 0, 7097830770726666240, 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 'https://www.w3schools.com/html/movie.mp4', 'Video Title 1');
INSERT INTO `videos` VALUES (2, '2023-08-18 12:01:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/320/918/427/sky-clouds-sunlight-dark-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_20mb.mp4', 'Video Title 2');
INSERT INTO `videos` VALUES (3, '2023-08-18 12:02:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/683/318/324/wolf-wolves-snow-wolf-landscape-thumb.jpg', 'https://www.w3school.com.cn/example/html5/mov_bbb.mp4', 'Video Title 3');
INSERT INTO `videos` VALUES (4, '2023-08-18 12:03:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/96/70/322/woman-happiness-sunrise-silhouette-thumb.jpg', 'http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4', 'Video Title 4');
INSERT INTO `videos` VALUES (5, '2023-08-18 12:04:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/653/876/844/road-forest-season-autumn-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/16/3be2e4ef4aa21bfe7493064a7415c34d.mp4', 'Video Title 5');
INSERT INTO `videos` VALUES (6, '2023-08-18 12:05:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/174/613/190/mountain-landscape-mountains-landscape-steinweg-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/02319a81c80afed90d9a2b9dc47f85b9.mp4', 'Video Title 6');
INSERT INTO `videos` VALUES (7, '2023-08-18 12:06:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/170/829/152/summerfield-woman-girl-sunset-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 7');
INSERT INTO `videos` VALUES (8, '2023-08-18 12:07:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/299/326/436/write-plan-business-startup-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4', 'Video Title 8');
INSERT INTO `videos` VALUES (9, '2023-08-18 12:08:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/704/633/113/scotland-landscape-scenic-mountains-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218093206z8V1JuPlpe.mp4', 'Video Title 9');
INSERT INTO `videos` VALUES (10, '2023-08-18 12:09:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/97/409/72/youth-active-jump-happy-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/2fca1c77730e54c7b500573c2437003f.mp4', 'Video Title 10');
INSERT INTO `videos` VALUES (11, '2023-08-18 13:00:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/20/987/594/woman-young-rain-pond-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218025702PSiVKDB5ap.mp4', 'Video Title 11');
INSERT INTO `videos` VALUES (12, '2023-08-18 13:01:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/740/837/204/spices-cinnamon-sticks-odor-aroma-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/202002181038474liyNnnSzz.mp4', 'Video Title 12');
INSERT INTO `videos` VALUES (13, '2023-08-18 13:02:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/27/812/1024/landscape-autumn-fog-village-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/02319a81c80afed90d9a2b9dc47f85b9.mp4', 'Video Title 13');
INSERT INTO `videos` VALUES (14, '2023-08-18 13:03:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/243/907/675/california-sunset-dusk-sky-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 14');
INSERT INTO `videos` VALUES (15, '2023-08-18 13:04:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/583/885/292/tianjin-twilight-city-scenery-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 15');
INSERT INTO `videos` VALUES (16, '2023-08-18 13:05:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/794/70/509/home-office-workstation-office-business-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/c292033ef110de9f42d7d539fe0423cf.mp4', 'Video Title 16');
INSERT INTO `videos` VALUES (17, '2023-08-18 13:06:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/908/912/899/pier-wooden-lake-ocean-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/97e3c56e283a10546f22204963086f65.mp4', 'Video Title 17');
INSERT INTO `videos` VALUES (18, '2023-08-18 13:07:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/833/42/545/coffee-cup-coffee-beans-coffee-cup-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/778c5884fa97f460dac8d90493c451de.mp4', 'Video Title 18');
INSERT INTO `videos` VALUES (19, '2023-08-18 13:08:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/57/631/722/barley-field-sunrise-morning-solar-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217021133Eggh6zdlAO.mp4', 'Video Title 19');
INSERT INTO `videos` VALUES (20, '2023-08-18 13:09:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/897/711/802/japanese-cherry-trees-flowers-spring-japanese-flowering-cherry-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/4417a27b1a656f4779eaa005ecd1a1a0.mp4', 'Video Title 20');
INSERT INTO `videos` VALUES (21, '2023-08-18 13:39:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/799/203/747/trees-avenue-autumn-away-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217104524H4D6lmByOe.mp4', 'Video Title 21');
INSERT INTO `videos` VALUES (22, '2023-08-18 13:40:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/871/443/724/stones-rocks-pebbles-tranquil-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217104524H4D6lmByOe.mp4', 'Video Title 22');
INSERT INTO `videos` VALUES (23, '2023-08-18 13:41:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/195/146/461/roses-bouquet-congratulations-arrangement-thumb.jpg', 'http://stream.iqilu.com/vod_bag_2016//2020/02/16/903BE158056C44fcA9524B118A5BF230/903BE158056C44fcA9524B118A5BF230_H264_mp4_500K.mp4', 'Video Title 23');
INSERT INTO `videos` VALUES (24, '2023-08-18 13:42:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/916/632/1022/office-notes-notepad-entrepreneur-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/16/20200216050645YIMfjPq5Nw.mp4', 'Video Title 24');
INSERT INTO `videos` VALUES (25, '2023-08-18 13:43:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/916/632/1022/office-notes-notepad-entrepreneur-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/16/3be2e4ef4aa21bfe7493064a7415c34d.mp4', 'Video Title 25');
INSERT INTO `videos` VALUES (26, '2023-08-18 13:44:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/361/515/668/background-berries-berry-blackberries-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4', 'Video Title 26');
INSERT INTO `videos` VALUES (27, '2023-08-18 13:45:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/834/732/890/landscape-spring-wood-scenic-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/09/20200209104902N3v5Vpxuvb.mp4', 'Video Title 27');
INSERT INTO `videos` VALUES (28, '2023-08-18 13:46:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/877/494/403/owl-bird-eyes-eagle-owl-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_30mb.mp4', 'Video Title 28');
INSERT INTO `videos` VALUES (29, '2023-08-18 13:47:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/367/415/619/paper-romance-symbol-valentine-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_20mb.mp4', 'Video Title 29');
INSERT INTO `videos` VALUES (30, '2023-08-18 13:48:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/105/931/754/portrait-woman-girl-female-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_10mb.mp4', 'Video Title 30');
INSERT INTO `videos` VALUES (31, '2023-08-18 14:00:00', 0, 1, 'https://i0.hippopx.com/photos/759/56/854/foggy-mist-forest-trees-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_5mb.mp4', 'Video Title 31');
INSERT INTO `videos` VALUES (32, '2023-08-18 14:01:00', 0, 2, 'https://i0.hippopx.com/photos/54/648/833/aurora-northen-lights-ice-mountain-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_2mb.mp4', 'Video Title 32');
INSERT INTO `videos` VALUES (33, '2023-08-18 14:02:00', 0, 3, 'https://i0.hippopx.com/photos/697/286/503/eye-green-eye-close-up-macro-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4', 'Video Title 33');
INSERT INTO `videos` VALUES (34, '2023-08-18 14:03:00', 0, 4, 'https://i0.hippopx.com/photos/971/691/698/mobile-phone-smartphone-3d-manipulation-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_30mb.mp4', 'Video Title 34');
INSERT INTO `videos` VALUES (35, '2023-08-18 14:04:00', 0, 5, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_30mb.mp4', 'Video Title 35');
INSERT INTO `videos` VALUES (36, '2023-08-18 14:05:00', 0, 6, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_20mb.mp4', 'Video Title 36');
INSERT INTO `videos` VALUES (37, '2023-08-18 14:06:00', 0, 7, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_20mb.mp4', 'Video Title 37');
INSERT INTO `videos` VALUES (38, '2023-08-18 14:07:00', 0, 8, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_10mb.mp4', 'Video Title 38');
INSERT INTO `videos` VALUES (39, '2023-08-18 14:08:00', 0, 9, 'https://i0.hippopx.com/photos/668/495/659/abstract-aqua-background-blue-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_5mb.mp4', 'Video Title 39');
INSERT INTO `videos` VALUES (40, '2023-08-18 14:09:00', 0, 10, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_5mb.mp4', 'Video Title 40');
INSERT INTO `videos` VALUES (41, '2023-08-18 14:39:00', 0, 11, 'https://i0.hippopx.com/photos/910/378/500/iceland-arctic-fox-animal-wildlife-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_2mb.mp4', 'Video Title 41');
INSERT INTO `videos` VALUES (42, '2023-08-18 14:40:00', 0, 12, 'https://i0.hippopx.com/photos/504/693/295/jay-bird-konar-winter-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_1mb.mp4', 'Video Title 42');
INSERT INTO `videos` VALUES (43, '2023-08-18 14:41:00', 0, 13, 'https://i0.hippopx.com/photos/358/761/178/girl-paintings-woman-flight-thumb.jpg', 'http://vjs.zencdn.net/v/oceans.mp4', 'Video Title 43');
INSERT INTO `videos` VALUES (44, '2023-08-18 14:42:00', 0, 14, 'https://i0.hippopx.com/photos/548/90/482/sunrise-phu-quoc-island-ocean-thumb.jpg', 'http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4', 'Video Title 44');
INSERT INTO `videos` VALUES (45, '2023-08-18 14:43:00', 0, 15, 'https://i0.hippopx.com/photos/532/7/742/rose-beautiful-beauty-bloom-thumb.jpg', 'https://media.w3.org/2010/05/sintel/trailer.mp4', 'Video Title 45');
INSERT INTO `videos` VALUES (46, '2023-08-18 14:44:00', 0, 16, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217104524H4D6lmByOe.mp4', 'Video Title 46');
INSERT INTO `videos` VALUES (47, '2023-08-18 14:45:00', 0, 17, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'http://stream.iqilu.com/vod_bag_2016//2020/02/16/903BE158056C44fcA9524B118A5BF230/903BE158056C44fcA9524B118A5BF230_H264_mp4_500K.mp4', 'Video Title 47');
INSERT INTO `videos` VALUES (48, '2023-08-18 14:46:00', 0, 18, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_30mb.mp4', 'Video Title 48');
INSERT INTO `videos` VALUES (49, '2023-08-18 14:47:00', 0, 19, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_20mb.mp4', 'Video Title 49');
INSERT INTO `videos` VALUES (50, '2023-08-18 14:48:00', 0, 20, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'https://www.w3schools.com/html/movie.mp4', 'Video Title 50');

SET FOREIGN_KEY_CHECKS = 1;
