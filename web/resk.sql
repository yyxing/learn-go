/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : resk

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 22/12/2020 11:46:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '账户自增id',
  `account_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '账户编号 账户唯一标识',
  `account_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '账户名称',
  `account_type` tinyint NOT NULL COMMENT '账户类型 1货币账户 积分账户',
  `currency_code` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '货币类型编码',
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户编号 账户所属用户',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户名称',
  `balance` decimal(30, 6) NOT NULL COMMENT '账户可用余额',
  `status` tinyint NOT NULL COMMENT '账户状态 是否被冻结 1正常 -1冻结',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account_no_idx`(`account_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for account_log
-- ----------------------------
DROP TABLE IF EXISTS `account_log`;
CREATE TABLE `account_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易单号，全局唯一',
  `log_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '流水编号，全局唯一',
  `account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账户编号',
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户编号',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户名称',
  `counterparty_user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '目标用户编号',
  `counterparty_username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '目标用户名称',
  `counterparty_account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '目标账户编号',
  `amount` decimal(30, 6) NOT NULL COMMENT '交易金额',
  `balance` decimal(30, 6) NULL DEFAULT NULL COMMENT '交易之后的余额',
  `trade_type` tinyint NULL DEFAULT NULL COMMENT '交易类型：100创建账户 >0收入 <0 支出',
  `status` tinyint NULL DEFAULT NULL COMMENT '交易状态',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '交易描述',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trade_no_idx`(`trade_no`) USING BTREE,
  UNIQUE INDEX `log_no_idx`(`log_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 130 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for red_envelope_good
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_good`;
CREATE TABLE `red_envelope_good`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `envelope_type` tinyint NOT NULL COMMENT '红包类型：普通红包，碰运气红包,过期红包',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名称',
  `user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户编号, 红包所属用户 ',
  `account_no` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '发送者支付账户',
  `blessing` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '祝福语',
  `amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '红包总金额',
  `each_amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '单个红包金额，碰运气红包无效',
  `quantity` int UNSIGNED NOT NULL COMMENT '红包总数量 ',
  `remain_amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '红包剩余金额额',
  `remain_quantity` int UNSIGNED NOT NULL COMMENT '红包剩余数量 ',
  `expired_at` datetime(3) NOT NULL COMMENT '过期时间',
  `order_status` tinyint NOT NULL COMMENT '红包/订单状态：0 创建、1 发布启用、2过期、3失效',
  `order_type` tinyint NOT NULL COMMENT '订单类型：发布单、退款单 ',
  `pay_status` tinyint NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `envelope_no_idx`(`envelope_no`) USING BTREE,
  INDEX `id_user_idx`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1281 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for red_envelope_item
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_item`;
CREATE TABLE `red_envelope_item`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `item_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包订单详情编号 ',
  `envelope_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `receive_username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '红包接收者用户名称',
  `receive_user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包接收者用户编号 ',
  `amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '收到金额',
  `quantity` int UNSIGNED NOT NULL COMMENT '收到数量：对于收红包来说是1 ',
  `remain_amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '收到后红包剩余金额',
  `remain_quantity` int NOT NULL COMMENT '收到后红包数量',
  `account_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包接收者账户ID',
  `pay_status` tinyint NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `order_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '订单描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `item_no_idx`(`item_no`) USING BTREE,
  INDEX `envelope_no_idx`(`envelope_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 171 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
