
SET SQL_REQUIRE_PRIMARY_KEY = OFF;

CREATE DATABASE IF NOT EXISTS `blog`;
USE `blog`;

CREATE TABLE IF NOT EXISTS `user_info` (
  `user_info_id` int NOT NULL AUTO_INCREMENT COMMENT '회원 고유키',
  `profile_img_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '이미지 경로',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '이름',
  `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'S' COMMENT '성별',
  `social` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'N' COMMENT '가입 방식',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '이메일',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '비밀번호',
  `country` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '국가명',
  `lang` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '언어',
  `ip_addr` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '접속 IP',
  `login_date` timestamp DEFAULT NULL COMMENT '로그인 시간',
  `logout_date` timestamp DEFAULT NULL COMMENT '로그아웃 시간',
  `insert_date` timestamp DEFAULT NULL COMMENT '가입 날짜',
  `update_date` timestamp DEFAULT NULL COMMENT '수정 날짜',
  `block` tinyint(1) DEFAULT '0' COMMENT '차단여부',
  PRIMARY KEY (`user_info_id`),
  UNIQUE KEY `user_info_UN` (`email`)
);

