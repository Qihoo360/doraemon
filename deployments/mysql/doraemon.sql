CREATE DATABASE IF NOT EXISTS doraemon;

-- ----------------------------
-- Table structure for alert
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`alert`;
CREATE TABLE `doraemon`.`alert`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rule_id` bigint(20) NOT NULL,
  `labels` varchar(4095) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `value` double NOT NULL DEFAULT 0,
  `count` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `summary` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `hostname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `confirmed_by` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `fired_at` datetime(0) NOT NULL,
  `confirmed_at` datetime(0) DEFAULT NULL,
  `confirmed_before` datetime(0) DEFAULT NULL,
  `resolved_at` datetime(0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ruleid_labels_firedat`(`rule_id`, `labels`(255), `fired_at`) USING BTREE,
  INDEX `alert_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 163632 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`config`;
CREATE TABLE `doraemon`.`config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `service_id` bigint(20) NOT NULL DEFAULT 0,
  `idc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `proto` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `auto` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `port` int(11) NOT NULL DEFAULT 0,
  `metric` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`group`;
CREATE TABLE `doraemon`.`group`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `user` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`host`;
CREATE TABLE `doraemon`.`host`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) NOT NULL DEFAULT 0,
  `hostname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `mid`(`mid`, `hostname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for maintain
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`maintain`;
CREATE TABLE `doraemon`.`maintain`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `flag` tinyint(1) NOT NULL DEFAULT 0,
  `time_start` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `time_end` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `month` int(11) NOT NULL DEFAULT 0,
  `day_start` tinyint(4) NOT NULL DEFAULT 0,
  `day_end` tinyint(4) NOT NULL DEFAULT 0,
  `valid` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `maintain_valid_day_start_day_end_flag_time_start_time_end`(`valid`, `day_start`, `day_end`, `flag`, `time_start`, `time_end`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for manage
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`manage`;
CREATE TABLE `doraemon`.`manage`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `servicename` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `servicename`(`servicename`) USING BTREE,
  INDEX `manage_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for plan
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`plan`;
CREATE TABLE `doraemon`.`plan`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rule_labels` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 401 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for plan_receiver
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`plan_receiver`;
CREATE TABLE `doraemon`.`plan_receiver` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `plan_id` bigint(20) NOT NULL,
  `start_time` varchar(31) NOT NULL DEFAULT '',
  `end_time` varchar(31) NOT NULL DEFAULT '',
  `start` int(11) NOT NULL DEFAULT '0',
  `period` int(11) NOT NULL DEFAULT '0',
  `expression` varchar(1023) NOT NULL DEFAULT '',
  `reverse_polish_notation` varchar(1023) NOT NULL DEFAULT '',
  `user` varchar(1023) NOT NULL DEFAULT '',
  `group` varchar(1023) NOT NULL DEFAULT '',
  `duty_group` varchar(255) NOT NULL DEFAULT '',
  `odin_group` varchar(255) NOT NULL DEFAULT '',
  `method` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `plan_receiver_plan_id` (`plan_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for prom
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`prom`;
CREATE TABLE `doraemon`.`prom`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for rule
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`rule`;
CREATE TABLE `doraemon`.`rule`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `expr` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `op` varchar(31) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `value` varchar(31) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `for` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `summary` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `prom_id` bigint(20) NOT NULL,
  `plan_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 870 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for inhibit_logs
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`inhibit_logs`;
CREATE TABLE `doraemon`.`inhibit_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `alert_id` bigint(20) NOT NULL DEFAULT '0',
  `summary` varchar(1023) NOT NULL DEFAULT '',
  `labels` varchar(255) NOT NULL DEFAULT '',
  `source_expression` varchar(1023) NOT NULL DEFAULT '',
  `target_expression` varchar(1023) NOT NULL DEFAULT '',
  `sources` varchar(127) NOT NULL DEFAULT '',
  `relate_labels` varchar(255) NOT NULL DEFAULT '',
  `trigger_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for inhibits
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`inhibits`;
CREATE TABLE `doraemon`.`inhibits` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `source_expression` varchar(1023) NOT NULL DEFAULT '',
  `source_reverse_polish_notation` varchar(1023) NOT NULL DEFAULT '',
  `target_expression` varchar(1023) NOT NULL DEFAULT '',
  `target_reverse_polish_notation` varchar(1023) NOT NULL DEFAULT '',
  `labels` varchar(1023) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for logs
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`logs`;
CREATE TABLE `doraemon`.`logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `target` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) NOT NULL DEFAULT '',
  `user` varchar(255) NOT NULL DEFAULT '',
  `param` varchar(1023) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `doraemon`.`users`;
CREATE TABLE `doraemon`.`users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `password` varchar(1023) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

INSERT INTO `doraemon`.`users`(`id`, `name`, `password`) VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e');

