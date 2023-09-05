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

 Date: 05/09/2023 23:15:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `video_id` bigint NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `is_deleted` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 189 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (182, 40, 7098191152603463680, '测试一下', '2023-09-04 21:06:01', 0);
INSERT INTO `comments` VALUES (183, 40, 7098191152603463680, '测试一下', '2023-09-04 21:06:09', 0);
INSERT INTO `comments` VALUES (184, 40, 7098191152603463680, '测试', '2023-09-04 21:06:12', 0);
INSERT INTO `comments` VALUES (185, 40, 7098191152603463680, '测试', '2023-09-04 21:08:53', 1);
INSERT INTO `comments` VALUES (186, 39, 7098191152603463680, '测试', '2023-09-04 21:09:02', 0);
INSERT INTO `comments` VALUES (187, 50, 7097868890591461376, 'mytest01', '2023-09-04 21:13:06', 0);
INSERT INTO `comments` VALUES (188, 50, 7097868890591461376, 'mytest02', '2023-09-04 21:13:14', 1);

-- ----------------------------
-- Table structure for favoriterelations
-- ----------------------------
DROP TABLE IF EXISTS `favoriterelations`;
CREATE TABLE `favoriterelations`  (
  `id` bigint NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favoriterelations
-- ----------------------------
INSERT INTO `favoriterelations` VALUES (7104449508502667264, '2023-09-05 05:05:57', 0, 7098191152603463680, 40);
INSERT INTO `favoriterelations` VALUES (7104450268267282432, '2023-09-05 05:08:58', 0, 7098191152603463680, 39);
INSERT INTO `favoriterelations` VALUES (7104450591509708800, '2023-09-05 05:10:15', 0, 7098191152603463680, 38);
INSERT INTO `favoriterelations` VALUES (7104452611759144960, '2023-09-05 05:18:17', 0, 7098191152603463680, 50);
INSERT INTO `favoriterelations` VALUES (7104452762095583232, '2023-09-05 05:18:53', 0, 7097868890591461376, 50);
INSERT INTO `favoriterelations` VALUES (7104628340518027264, '2023-09-05 16:56:34', 0, 7098191152603463680, 49);
INSERT INTO `favoriterelations` VALUES (7104663516778332160, '2023-09-05 11:16:21', 0, 7104663433093578752, 49);
INSERT INTO `favoriterelations` VALUES (7104802831944646656, '2023-09-05 20:29:56', 0, 7104663433093578752, 46);

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
  `id` bigint NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `follow_user_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follow
-- ----------------------------
INSERT INTO `follow` VALUES (7104834959315042304, '2023-09-05 22:50:32', 1, 7104663433093578752, 7097830770726666240);
INSERT INTO `follow` VALUES (7104836087125639168, '2023-09-05 22:43:44', 1, 7104663433093578752, 7097847496482750464);
INSERT INTO `follow` VALUES (7104838321628512256, '2023-09-05 23:12:06', 0, 7104663433093578752, 7097862145206910976);
INSERT INTO `follow` VALUES (7104843792712204288, '2023-09-05 23:12:42', 0, 7097830770726666240, 7097847496482750464);
INSERT INTO `follow` VALUES (7104843858810241024, '2023-09-05 23:12:58', 0, 7097847496482750464, 7097830770726666240);

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_user_id` bigint NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `to_user_id` bigint NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of messages
-- ----------------------------
INSERT INTO `messages` VALUES (1, 7098191152603463680, 'Hello!', 7097877640291287040, '2023-08-01 13:42:09', 0);
INSERT INTO `messages` VALUES (2, 7098191152603463680, 'Hi there!', 7097877640291287040, '2023-08-02 13:42:09', 0);
INSERT INTO `messages` VALUES (3, 7098191152603463680, 'How are you?', 7097877640291287040, '2023-08-03 13:42:09', 0);
INSERT INTO `messages` VALUES (4, 7098191152603463680, 'I am good!', 7097877640291287040, '2023-08-04 13:42:09', 0);
INSERT INTO `messages` VALUES (5, 7098191152603463680, 'Good to hear!', 7097877640291287040, '2023-08-05 13:42:09', 0);
INSERT INTO `messages` VALUES (6, 7097877640291287040, 'What are you up to?', 7098191152603463680, '2023-08-06 13:42:09', 0);
INSERT INTO `messages` VALUES (7, 7097877640291287040, 'Just working on a project.', 7098191152603463680, '2023-08-07 13:42:09', 0);
INSERT INTO `messages` VALUES (8, 7097877640291287040, 'Sounds interesting!', 7098191152603463680, '2023-08-08 13:42:09', 0);
INSERT INTO `messages` VALUES (9, 7097877640291287040, 'Thanks!', 7098191152603463680, '2023-08-09 13:42:09', 0);
INSERT INTO `messages` VALUES (10, 7097877640291287040, 'You are welcome!', 7098191152603463680, '2023-08-10 13:42:09', 0);
INSERT INTO `messages` VALUES (11, 7098191152603463680, '你好', 7097877640291287040, '2023-09-03 22:14:39', 0);
INSERT INTO `messages` VALUES (12, 7098191152603463680, '你好1', 7097877640291287040, '2023-09-03 22:16:09', 0);
INSERT INTO `messages` VALUES (13, 7098191152603463680, '你好2', 7097877640291287040, '2023-09-03 22:28:34', 0);
INSERT INTO `messages` VALUES (14, 7098191152603463680, '你好2', 7097877640291287040, '2023-09-03 22:36:40', 0);
INSERT INTO `messages` VALUES (15, 7097868890591461376, '你好1', 7097868890591461376, '2023-09-05 05:15:47', 0);
INSERT INTO `messages` VALUES (16, 7097868890591461376, '你好3', 7098191152603463680, '2023-09-05 05:23:37', 0);
INSERT INTO `messages` VALUES (17, 7097868890591461376, '你好3', 7098191152603463680, '2023-09-05 05:23:39', 0);

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
INSERT INTO `user` VALUES (7097830770726666240, 'user009', 1, 1, 1, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', '活在当下，珍惜每一刻。', '5', 0, 0, 0, '2023-08-18 05:45:27', '$2a$10$6DREfmPqegOaWwTmx0LwfO8JJp3koS.GkCqW.SP60i.u9bCJT07Ya');
INSERT INTO `user` VALUES (7097847496482750464, 'user008', 1, 1, 0, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', '要么找到一条道路，要么创造一条道路。', '0', 0, 0, 0, '2023-08-18 06:51:55', '$2a$10$bLZlQRmTShGGyUrlnbG6wu.RIHg58Y5CMhqNnJUw7xuX7dY1l2InG');
INSERT INTO `user` VALUES (7097862145206910976, 'user007', 0, 1, 0, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', '每一步都是向前的一步。', '0', 0, 0, 0, '2023-08-18 07:50:07', '$2a$10$84AtF6vZFFqfYLGUnb0/k.tWu9wiWMyyQdR..QlEZg8GHp2SxUD4e');
INSERT INTO `user` VALUES (7097863168206045184, 'user005', 0, 0, 0, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', '生活不是等待风暴过去，而是学会在雨中跳舞。', '0', 0, 0, 0, '2023-08-18 07:54:11', '$2a$10$4Mf/iZPjcQ9PkLjuSSnOieQNO4seMSeC/Z7SxAM.iJOY77B9UgDbO');
INSERT INTO `user` VALUES (7097867658552410112, 'user003', 0, 0, 0, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', '成功的秘诀是坚持不懈。', '0', 0, 0, 0, '2023-08-18 08:12:02', '$2a$10$3TEEp112jDvDL3N10pOOlu2F5Pv3md1ubWZF8uwDliXuWwsCr.T5i');
INSERT INTO `user` VALUES (7097868890591461376, 'user001', 0, 0, 0, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', '生活是一本书，不读它就会显得很无趣。', '1', 0, 1, 0, '2023-08-18 08:16:56', '$2a$10$./mA1FZRSH93BcsGnNE68eE1o8z9gzzD2izwFU0uJPGDSeYRyK38C');
INSERT INTO `user` VALUES (7097876051899973632, 'user002', 0, 0, 0, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', '每一天都是一个新的机会，抓住它', '1', 0, 0, 0, '2023-08-18 08:45:23', '$2a$10$2wye4LoP8x8Qk3Ff6Xx8qe0pBWibcKiubxzLKx3lOJt8uHlJihET.');
INSERT INTO `user` VALUES (7097876434839928832, 'user006', 0, 0, 0, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', '相信自己，你能做到！', '1', 0, 0, 0, '2023-08-18 08:46:54', '$2a$10$MrJX13gw8HEJw4k54xHRYup1zOfzd9LCAVM7j.Y8aEF8dQAPhGwiO');
INSERT INTO `user` VALUES (7097877640291287040, 'user004', 0, 0, 0, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', '做自己，因为其他角色已经有人在演了', '0', 0, 0, 0, '2023-08-18 08:51:42', '$2a$10$Ugw4VGXUYz6PuhkjH8U4Cu3khUcSg.vW9Zv/F9XU6ON8wAFDZApka');
INSERT INTO `user` VALUES (7098191152603463680, 'admin', 0, 0, 0, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', '不要等待机会，创造机会。', '0', 0, 5, 0, '2023-08-19 05:37:29', '$2a$10$Ma/oO3FoAGevdJESwlwtcOmvx.G/n8LIwg7wn3vjyNWsTWh3T828m');
INSERT INTO `user` VALUES (7104663433093578752, 'test3', 1, 0, 0, '', '', '', '0', 0, 2, 0, '2023-09-05 11:16:01', '$2a$10$j/EemFxzlrxZVIPmte3kHOleKp3trQ7iVGG2HMK1F1EP9xITbreXK');

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
  `comment_count` int NULL DEFAULT 0,
  `favorite_count` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, '2023-08-18 12:00:00', 0, 7097830770726666240, 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 'https://www.w3schools.com/html/movie.mp4', 'Video Title 1', 0, 0);
INSERT INTO `videos` VALUES (2, '2023-08-18 12:01:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/320/918/427/sky-clouds-sunlight-dark-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_20mb.mp4', 'Video Title 2', 0, 0);
INSERT INTO `videos` VALUES (3, '2023-08-18 12:02:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/683/318/324/wolf-wolves-snow-wolf-landscape-thumb.jpg', 'https://www.w3school.com.cn/example/html5/mov_bbb.mp4', 'Video Title 3', 0, 0);
INSERT INTO `videos` VALUES (4, '2023-08-18 12:03:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/96/70/322/woman-happiness-sunrise-silhouette-thumb.jpg', 'http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4', 'Video Title 4', 0, 0);
INSERT INTO `videos` VALUES (5, '2023-08-18 12:04:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/653/876/844/road-forest-season-autumn-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/16/3be2e4ef4aa21bfe7493064a7415c34d.mp4', 'Video Title 5', 0, 0);
INSERT INTO `videos` VALUES (6, '2023-08-18 12:05:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/174/613/190/mountain-landscape-mountains-landscape-steinweg-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/02319a81c80afed90d9a2b9dc47f85b9.mp4', 'Video Title 6', 0, 0);
INSERT INTO `videos` VALUES (7, '2023-08-18 12:06:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/170/829/152/summerfield-woman-girl-sunset-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 7', 0, 0);
INSERT INTO `videos` VALUES (8, '2023-08-18 12:07:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/299/326/436/write-plan-business-startup-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4', 'Video Title 8', 0, 0);
INSERT INTO `videos` VALUES (9, '2023-08-18 12:08:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/704/633/113/scotland-landscape-scenic-mountains-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218093206z8V1JuPlpe.mp4', 'Video Title 9', 0, 0);
INSERT INTO `videos` VALUES (10, '2023-08-18 12:09:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/97/409/72/youth-active-jump-happy-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/2fca1c77730e54c7b500573c2437003f.mp4', 'Video Title 10', 0, 0);
INSERT INTO `videos` VALUES (11, '2023-08-18 13:00:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/20/987/594/woman-young-rain-pond-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218025702PSiVKDB5ap.mp4', 'Video Title 11', 0, 0);
INSERT INTO `videos` VALUES (12, '2023-08-18 13:01:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/740/837/204/spices-cinnamon-sticks-odor-aroma-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/202002181038474liyNnnSzz.mp4', 'Video Title 12', 0, 0);
INSERT INTO `videos` VALUES (13, '2023-08-18 13:02:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/27/812/1024/landscape-autumn-fog-village-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/02319a81c80afed90d9a2b9dc47f85b9.mp4', 'Video Title 13', 0, 0);
INSERT INTO `videos` VALUES (14, '2023-08-18 13:03:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/243/907/675/california-sunset-dusk-sky-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 14', 0, 0);
INSERT INTO `videos` VALUES (15, '2023-08-18 13:04:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/583/885/292/tianjin-twilight-city-scenery-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 15', 0, 0);
INSERT INTO `videos` VALUES (16, '2023-08-18 13:05:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/794/70/509/home-office-workstation-office-business-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/c292033ef110de9f42d7d539fe0423cf.mp4', 'Video Title 16', 0, 0);
INSERT INTO `videos` VALUES (17, '2023-08-18 13:06:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/908/912/899/pier-wooden-lake-ocean-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/97e3c56e283a10546f22204963086f65.mp4', 'Video Title 17', 0, 0);
INSERT INTO `videos` VALUES (18, '2023-08-18 13:07:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/833/42/545/coffee-cup-coffee-beans-coffee-cup-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/778c5884fa97f460dac8d90493c451de.mp4', 'Video Title 18', 0, 0);
INSERT INTO `videos` VALUES (19, '2023-08-18 13:08:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/57/631/722/barley-field-sunrise-morning-solar-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217021133Eggh6zdlAO.mp4', 'Video Title 19', 0, 0);
INSERT INTO `videos` VALUES (20, '2023-08-18 13:09:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/897/711/802/japanese-cherry-trees-flowers-spring-japanese-flowering-cherry-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/17/4417a27b1a656f4779eaa005ecd1a1a0.mp4', 'Video Title 20', 0, 0);
INSERT INTO `videos` VALUES (21, '2023-08-18 13:39:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/799/203/747/trees-avenue-autumn-away-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217104524H4D6lmByOe.mp4', 'Video Title 21', 0, 0);
INSERT INTO `videos` VALUES (22, '2023-08-18 13:40:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/871/443/724/stones-rocks-pebbles-tranquil-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/17/20200217104524H4D6lmByOe.mp4', 'Video Title 22', 0, 0);
INSERT INTO `videos` VALUES (23, '2023-08-18 13:41:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/195/146/461/roses-bouquet-congratulations-arrangement-thumb.jpg', 'http://stream.iqilu.com/vod_bag_2016//2020/02/16/903BE158056C44fcA9524B118A5BF230/903BE158056C44fcA9524B118A5BF230_H264_mp4_500K.mp4', 'Video Title 23', 0, 0);
INSERT INTO `videos` VALUES (24, '2023-08-18 13:42:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/916/632/1022/office-notes-notepad-entrepreneur-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/16/20200216050645YIMfjPq5Nw.mp4', 'Video Title 24', 0, 0);
INSERT INTO `videos` VALUES (25, '2023-08-18 13:43:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/916/632/1022/office-notes-notepad-entrepreneur-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/16/3be2e4ef4aa21bfe7493064a7415c34d.mp4', 'Video Title 25', 0, 0);
INSERT INTO `videos` VALUES (26, '2023-08-18 13:44:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/361/515/668/background-berries-berry-blackberries-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4', 'Video Title 26', 0, 0);
INSERT INTO `videos` VALUES (27, '2023-08-18 13:45:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/834/732/890/landscape-spring-wood-scenic-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/09/20200209104902N3v5Vpxuvb.mp4', 'Video Title 27', 0, 0);
INSERT INTO `videos` VALUES (28, '2023-08-18 13:46:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/877/494/403/owl-bird-eyes-eagle-owl-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_30mb.mp4', 'Video Title 28', 0, 0);
INSERT INTO `videos` VALUES (29, '2023-08-18 13:47:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/367/415/619/paper-romance-symbol-valentine-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_20mb.mp4', 'Video Title 29', 0, 0);
INSERT INTO `videos` VALUES (30, '2023-08-18 13:48:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/105/931/754/portrait-woman-girl-female-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_10mb.mp4', 'Video Title 30', 0, 0);
INSERT INTO `videos` VALUES (31, '2023-08-18 14:00:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/759/56/854/foggy-mist-forest-trees-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_5mb.mp4', 'Video Title 31', 0, 0);
INSERT INTO `videos` VALUES (32, '2023-08-18 14:01:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/54/648/833/aurora-northen-lights-ice-mountain-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_2mb.mp4', 'Video Title 32', 0, 0);
INSERT INTO `videos` VALUES (33, '2023-08-18 14:02:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/697/286/503/eye-green-eye-close-up-macro-thumb.jpg', 'https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4', 'Video Title 33', 0, 0);
INSERT INTO `videos` VALUES (34, '2023-08-18 14:03:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/971/691/698/mobile-phone-smartphone-3d-manipulation-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_30mb.mp4', 'Video Title 34', 0, 0);
INSERT INTO `videos` VALUES (35, '2023-08-18 14:04:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_30mb.mp4', 'Video Title 35', 0, 0);
INSERT INTO `videos` VALUES (36, '2023-08-18 14:05:00', 0, 7098191152603463680, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_20mb.mp4', 'Video Title 36', 0, 0);
INSERT INTO `videos` VALUES (37, '2023-08-18 14:06:00', 0, 7097877640291287040, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_20mb.mp4', 'Video Title 37', 0, 0);
INSERT INTO `videos` VALUES (38, '2023-08-18 14:07:00', 0, 7097876434839928832, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_10mb.mp4', 'Video Title 38', 0, 1);
INSERT INTO `videos` VALUES (39, '2023-08-18 14:08:00', 0, 7097876051899973632, 'https://i0.hippopx.com/photos/668/495/659/abstract-aqua-background-blue-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_5mb.mp4', 'Video Title 39', 1, 1);
INSERT INTO `videos` VALUES (40, '2023-08-18 14:09:00', 0, 7097868890591461376, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_5mb.mp4', 'Video Title 40', 3, 1);
INSERT INTO `videos` VALUES (41, '2023-08-18 14:39:00', 0, 7097867658552410112, 'https://i0.hippopx.com/photos/910/378/500/iceland-arctic-fox-animal-wildlife-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_2mb.mp4', 'Video Title 41', 0, 0);
INSERT INTO `videos` VALUES (42, '2023-08-18 14:40:00', 0, 7097863168206045184, 'https://i0.hippopx.com/photos/504/693/295/jay-bird-konar-winter-thumb.jpg', 'https://sample-videos.com/video123/mp4/360/big_buck_bunny_360p_1mb.mp4', 'Video Title 42', 0, 0);
INSERT INTO `videos` VALUES (43, '2023-08-18 14:41:00', 0, 7097862145206910976, 'https://i0.hippopx.com/photos/358/761/178/girl-paintings-woman-flight-thumb.jpg', 'http://vjs.zencdn.net/v/oceans.mp4', 'Video Title 43', 0, 0);
INSERT INTO `videos` VALUES (44, '2023-08-18 14:42:00', 0, 7097847496482750464, 'https://i0.hippopx.com/photos/548/90/482/sunrise-phu-quoc-island-ocean-thumb.jpg', 'http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4', 'Video Title 44', 0, 0);
INSERT INTO `videos` VALUES (45, '2023-08-18 14:43:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/532/7/742/rose-beautiful-beauty-bloom-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/2fca1c77730e54c7b500573c2437003f.mp4', 'Video Title 45', 0, 0);
INSERT INTO `videos` VALUES (46, '2023-08-18 14:44:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/617/220/967/beautiful-girl-in-the-park-lying-on-the-leaves-autumn-portrait-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218025702PSiVKDB5ap.mp4', 'Video Title 46', 0, 1);
INSERT INTO `videos` VALUES (47, '2023-08-18 14:45:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/518/365/148/woods-forest-nature-landscape-thumb.jpg', 'https://stream7.iqilu.com/10339/upload_transcode/202002/18/202002181038474liyNnnSzz.mp4', 'Video Title 47', 0, 0);
INSERT INTO `videos` VALUES (48, '2023-08-18 14:46:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/574/398/392/beach-lagoon-sunset-sundown-thumb.jpg', 'https://stream7.iqilu.com/10339/article/202002/18/02319a81c80afed90d9a2b9dc47f85b9.mp4', 'Video Title 48', 0, 0);
INSERT INTO `videos` VALUES (49, '2023-08-18 14:47:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/801/280/882/stained-glass-spiral-circle-pattern-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 49', 0, 2);
INSERT INTO `videos` VALUES (50, '2023-08-18 14:48:00', 0, 7097830770726666240, 'https://i0.hippopx.com/photos/458/225/316/fox-tree-stump-sleeping-thumb.jpg', 'http://stream4.iqilu.com/ksd/video/2020/02/17/c5e02420426d58521a8783e754e9f4e6.mp4', 'Video Title 50', 1, 2);

SET FOREIGN_KEY_CHECKS = 1;
