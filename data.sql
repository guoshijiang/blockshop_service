/*
 Navicat Premium Data Transfer

 Source Server         : my8.0
 Source Server Type    : MySQL
 Source Server Version : 80014
 Source Host           : localhost:3356
 Source Schema         : blockshop

 Target Server Type    : MySQL
 Target Server Version : 80014
 File Encoding         : 65001

 Date: 04/08/2021 21:53:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'fa-list',
  `is_show` tinyint(4) NOT NULL DEFAULT '1',
  `sort_id` int(11) NOT NULL DEFAULT '1000',
  `log_method` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '不记录',
  PRIMARY KEY (`id`),
  KEY `admin_menu_url` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, '后台首页', 'admin/index/index', 'fa-home', 1, 99, '不记录');
INSERT INTO `admin_menu` VALUES (2, 0, '系统管理', 'admin/sys', 'fa-desktop', 1, 1099, '不记录');
INSERT INTO `admin_menu` VALUES (3, 2, '用户管理', 'admin/admin_user/index', 'fa-user', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (4, 3, '添加用户界面', 'admin/admin_user/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (5, 3, '修改用户界面', 'admin/admin_user/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (6, 3, '删除用户', 'admin/admin_user/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (7, 2, '角色管理', 'admin/admin_role/index', 'fa-group', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (8, 7, '添加角色界面', 'admin/admin_role/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (9, 7, '修改角色界面', 'admin/admin_role/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (10, 7, '删除角色', 'admin/admin_role/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (11, 7, '角色授权界面', 'admin/admin_role/access', 'fa-key', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (12, 2, '菜单管理', 'admin/admin_menu/index', 'fa-align-justify', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (13, 12, '添加菜单界面', 'admin/admin_menu/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (14, 12, '修改菜单界面', 'admin/admin_menu/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (15, 12, '删除菜单', 'admin/admin_menu/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (16, 2, '操作日志', 'admin/admin_log/index', 'fa-keyboard-o', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (17, 16, '日志详情', 'admin/admin_log/view', 'fa-search-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (18, 2, '个人资料', 'admin/admin_user/profile', 'fa-smile-o', 1, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (19, 0, '订单管理', 'admin/order/mange', 'fa-first-order', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (20, 19, '订单管理', 'admin/order/index', 'fa-first-order', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (55, 3, '修改头像', 'admin/admin_user/update_avatar', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (56, 3, '添加用户', 'admin/admin_user/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (57, 3, '修改用户', 'admin/admin_user/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (58, 3, '用户启用', 'admin/admin_user/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (59, 3, '用户禁用', 'admin/admin_user/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (60, 3, '修改昵称', 'admin/admin_user/update_nickname', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (61, 3, '修改密码', 'admin/admin_user/update_password', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (62, 7, '创建角色', 'admin/admin_role/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (63, 7, '修改角色', 'admin/admin_role/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (64, 7, '启用角色', 'admin/admin_role/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (65, 7, '禁用角色', 'admin/admin_role/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (66, 7, '角色授权', 'admin/admin_role/access_operate', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (67, 12, '创建菜单', 'admin/admin_menu/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (68, 12, '修改菜单', 'admin/admin_menu/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (148, 0, '商品管理', 'admin/goods/mange', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (149, 148, '商品列表', 'admin/goods/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (150, 148, '商品分类', 'admin/goods/category/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (151, 149, '添加商品', 'admin/goods/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (152, 149, '编辑商品', 'admin/goods/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (153, 149, '更新商品', 'admin/goods/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (154, 149, '删除商品', 'admin/goods/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (155, 150, '添加商品分类', 'admin/goods/category/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (156, 150, '编辑商品分类', 'admin/goods/category/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (157, 150, '更新商品分类', 'admin/goods/category/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (158, 150, '删除商品分类', 'admin/goods/category/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (159, 0, '区块链', 'admin/block/manage', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (160, 159, '币种管理', 'admin/asset/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (161, 160, '币种-创建', 'admin/asset/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (162, 160, '币种-更新', 'admin/asset/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (163, 160, '币种-编辑', 'admin/asset/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (164, 160, '币种-删除', 'admin/asset/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (165, 0, '商户管理', 'admin/merchant/manage', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (166, 165, '商户列表', 'admin/merchant/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (167, 166, '商户-创建', 'admin/merchant/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (168, 166, '商户-编辑', 'admin/merchant/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (169, 166, '商户-更新', 'admin/merchant/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (170, 166, '商户-删除', 'admin/merchant/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (171, 0, '用户管理', 'admin/user/manage', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (172, 171, '用户列表', 'admin/user/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (173, 172, '用户-添加', 'admin/user/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (174, 172, '用户-编辑', 'admin/user/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (175, 172, '用户-更新', 'admin/user/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (176, 172, '用户-删除', 'admin/user/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (177, 159, '用户钱包', 'admin/wallet/user/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (178, 159, '资产记录', 'admin/wallet/record/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (179, 149, '商品-添加界面', 'admin/goods/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (180, 166, '添加商户-界面', 'admin/merchant/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (181, 150, '添加商品分类-界面', 'admin/goods/category/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (182, 148, '商品属性', 'admin/goods/type/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (183, 182, '商品属性-添加界面', 'admin/goods/type/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (184, 182, '商品属性-创建', 'admin/goods/type/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (185, 182, '商品属性-编辑界面', 'admin/goods/type/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (186, 182, '商品属性-更新', 'admin/goods/type/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (187, 182, '商品属性-删除', 'admin/goods/type/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (188, 0, '其他', 'admin/news/manage', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (189, 188, '公告管理', 'admin/news/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (190, 189, '添加公告-界面', 'admin/news/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (191, 189, '添加公告-创建', 'admin/news/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (192, 189, '编辑公告-界面', 'admin/news/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (193, 189, '编辑公告-更新', 'admin/news/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (194, 189, '删除公告-删除', 'admin/news/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (195, 0, '论坛', 'admin/forum/manage', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (196, 195, '论坛管理', 'admin/forum/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (197, 195, '论坛分类', 'admin/forum/category/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (198, 197, '添加论坛分类-界面', 'admin/forum/category/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (199, 197, '添加论坛分类-创建', 'admin/forum/category/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (200, 197, '编辑论坛分类-界面', 'admin/forum/category/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (201, 197, '编辑论坛分类-更新', 'admin/forum/category/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (202, 197, '论坛分类-删除', 'admin/forum/category/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (203, 188, '工单列表', 'admin/message/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (204, 203, '查看工单', 'admin/message/history', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (205, 203, '回复工单', 'history/message/send', 'fa-list', 0, 1000, '不记录');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
