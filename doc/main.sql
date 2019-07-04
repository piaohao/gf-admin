CREATE DATABASE gf_admin CHARACTER SET utf8 COLLATE utf8_general_ci;

create table gf_admin.sys_permission
(
	id int auto_increment comment 'id'
		primary key,
	code tinytext null comment '编码',
	parent_code tinytext null comment '父编码',
	name tinytext null comment '名称',
	icon tinytext null comment '图标',
	view varchar(255) null comment '前端视图',
	url tinytext null comment '链接',
	priority int null comment '排序',
	level int null comment '层级',
	is_menu int null comment '是否为菜单',
	status varchar(1) default 'Y' null comment '状态',
	create_time timestamp default current_timestamp() null comment '创建日期',
	update_time timestamp default current_timestamp() null on update current_timestamp() comment '更新日期'
)
comment '权限表';

INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (1, 'system', 'root', '系统管理', 'user', null, '#', 1, 1, 1, 'Y', '2019-01-29 05:58:02', '2019-03-03 21:27:53');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (2, 'system/user', 'system', '用户管理', 'fa-user', '/admin/sys_user', '/admin/sys_user/index', 1, 2, 1, 'Y', '2019-01-29 05:58:02', '2019-03-03 20:46:14');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (3, 'system/role', 'system', '角色管理', 'fa-user', '/admin/sys_role', '/admin/sys_role/index', 2, 2, 1, 'Y', '2019-01-29 05:58:02', '2019-03-03 20:46:14');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (4, 'system/permission', 'system', '权限管理', 'fa-user', '/admin/sys_permission', '/admin/sys_permission/index', 3, 2, 1, 'Y', '2019-01-29 05:58:02', '2019-03-03 20:46:14');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (6, 'root', 'super', '根目录', 'fa-folder', null, '#', 1, 0, 1, 'Y', '2019-02-12 21:31:50', '2019-02-13 13:40:52');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (12, 'system/user/add', 'system/user', '添加用户', '', null, '/admin/sys_user/add', 0, null, 0, 'Y', '2019-02-19 13:53:13', '2019-02-21 15:46:35');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (13, 'system/user/edit', 'system/user', '编辑用户', '', null, '/admin/sys_user/edit', 0, null, 0, 'Y', '2019-02-19 14:01:47', '2019-02-21 15:46:35');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (14, 'system/user/remove', 'system/user', '删除用户', '', null, '/admin/sys_user/remove', 0, null, 0, 'Y', '2019-02-19 14:02:03', '2019-02-21 15:46:35');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (15, 'system/role/add', 'system/role', '添加角色', '', null, '/admin/sys_role/add', 0, null, 0, 'Y', '2019-02-19 14:02:25', '2019-02-21 15:46:35');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (16, 'system/role/edit', 'system/role', '编辑角色', '', null, '/admin/sys_role/edit', 0, null, 0, 'Y', '2019-02-19 14:02:40', '2019-02-21 15:46:35');
INSERT INTO  gf_admin.sys_permission (id, code, parent_code, name, icon, view, url, priority, level, is_menu, status, create_time, update_time) VALUES (17, 'system/role/remove', 'system/role', '删除角色', '', null, '/admin/sys_role/remove', 0, null, 0, 'Y', '2019-02-19 14:02:54', '2019-02-21 15:46:35');

create table  gf_admin.sys_role
(
	id int auto_increment comment 'id'
		primary key,
	name varchar(20) null comment '角色名',
	description varchar(20) null comment '描述',
	channel_type varchar(20) null comment '渠道类别',
	create_time timestamp default current_timestamp() null comment '创建日期',
	update_time timestamp default current_timestamp() null comment '更新日期'
)
comment '角色表';

INSERT INTO  gf_admin.sys_role (id, name, description, channel_type, create_time, update_time) VALUES (1, 'admin', '管理员', null, '2019-02-13 20:57:49', '2019-02-13 20:57:49');

create table  gf_admin.sys_role_permission
(
	id int auto_increment comment 'id'
		primary key,
	role_id int not null comment '角色id',
	permission_id int not null comment '权限id',
	constraint sys_role_permission_pk
		unique (role_id, permission_id)
)
comment '角色权限表';

INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (177, 1, 1);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (178, 1, 2);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (181, 1, 3);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (185, 1, 4);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (176, 1, 6);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (179, 1, 12);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (180, 1, 13);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (182, 1, 15);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (183, 1, 16);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (184, 1, 17);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (186, 1, 18);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (187, 1, 19);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (191, 1, 20);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (195, 1, 21);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (199, 1, 22);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (203, 1, 23);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (207, 1, 24);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (188, 1, 25);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (189, 1, 26);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (190, 1, 27);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (192, 1, 28);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (193, 1, 29);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (194, 1, 30);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (196, 1, 31);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (197, 1, 32);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (198, 1, 33);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (200, 1, 34);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (201, 1, 35);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (202, 1, 36);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (204, 1, 37);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (205, 1, 38);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (206, 1, 39);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (208, 1, 40);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (209, 1, 41);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (210, 1, 42);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (211, 1, 43);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (212, 1, 44);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (213, 1, 45);
INSERT INTO  gf_admin.sys_role_permission (id, role_id, permission_id) VALUES (214, 1, 46);

create table  gf_admin.sys_user
(
	id int auto_increment comment '用户id'
		primary key,
	username varchar(20) null comment '账户名',
	password varchar(50) not null comment '密码',
	salt varchar(50) null comment '密码盐',
	nickname varchar(20) null comment '昵称',
	mobile varchar(50) null comment '电话',
	status varchar(1) default 'Y' not null comment '状态',
	create_time timestamp default current_timestamp() null comment '创建日期',
	update_time timestamp default current_timestamp() not null comment '更新日期'
)
comment '用户表';

INSERT INTO  gf_admin.sys_user (id, username, password, salt, nickname, mobile, status, create_time, update_time) VALUES (1, 'admin', '47a049fb091b28c47aa604ad24583c29', '8xua3', 'admin', '1111', 'N', '2019-02-11 13:17:37', '2019-02-11 13:17:37');

create table  gf_admin.sys_user_role
(
	id int auto_increment comment 'id'
		primary key,
	user_id int not null comment '用户id',
	role_id int not null comment '角色id',
	constraint sys_user_role_pk
		unique (user_id, role_id)
)
comment '用户角色表';

INSERT INTO  gf_admin.sys_user_role ( user_id, role_id) VALUES ( 1, 1);

