CREATE TABLE `admin_api` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `method` enum('POST','GET') COLLATE latin1_general_ci NOT NULL DEFAULT 'GET',
  `secure` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `url` varchar(50) COLLATE latin1_general_ci NOT NULL COMMENT '/api',
  `description` text COLLATE latin1_general_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`method`,`url`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_api_query` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `api_id` int(10) unsigned NOT NULL DEFAULT '0',
  `query` text COLLATE latin1_general_ci,
  `params` text COLLATE latin1_general_ci,
  `property` varchar(20) COLLATE latin1_general_ci DEFAULT '',
  `sequence` smallint(5) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `api_id` (`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(20) COLLATE latin1_general_ci NOT NULL,
  `status` enum('enabled','disabled') COLLATE latin1_general_ci NOT NULL DEFAULT 'disabled',
  PRIMARY KEY (`id`),
  KEY `created_at` (`created_at`),
  KEY `name` (`name`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_group_page` (
  `group_id` int(10) unsigned NOT NULL,
  `page_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`group_id`,`page_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_group_api` (
  `group_id` INT UNSIGNED NOT NULL,
  `api_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`group_id`, `api_id`));

CREATE TABLE `admin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `page_id` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(50) COLLATE latin1_general_ci NOT NULL,
  `icon` text COLLATE latin1_general_ci,
  `sequence` smallint(5) unsigned NOT NULL DEFAULT '0',
  `status` enum('draft','active','suspended','deleted') COLLATE latin1_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `page_id` (`page_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_page` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(50) COLLATE latin1_general_ci NOT NULL,
  `url` varchar(100) COLLATE latin1_general_ci DEFAULT NULL,
  `title` varchar(50) COLLATE latin1_general_ci NOT NULL DEFAULT '',
  `description` text COLLATE latin1_general_ci,
  `layout` longtext COLLATE latin1_general_ci,
  `status` enum('draft','publish') COLLATE latin1_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `url` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_page_api` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `page_id` int(10) unsigned NOT NULL,
  `api_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `page_id_api_id` (`page_id`,`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `password_expired_at` timestamp NULL DEFAULT NULL,
  `initial` varchar(3) COLLATE latin1_general_ci DEFAULT NULL,
  `username` varchar(50) COLLATE latin1_general_ci NOT NULL,
  `secret` varchar(32) COLLATE latin1_general_ci NOT NULL,
  `name` varchar(20) COLLATE latin1_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `initial` (`initial`),
  KEY `created_at` (`created_at`),
  KEY `username` (`username`),
  KEY `password_expired_at` (`password_expired_at`),
  FULLTEXT KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_user_group` (
  `user_id` int(10) unsigned NOT NULL,
  `group_id` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`group_id`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

CREATE TABLE `admin_user_session` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `expired_at` timestamp NULL DEFAULT NULL,
  `status` enum('online','expired') COLLATE latin1_general_ci NOT NULL DEFAULT 'online',
  `client` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `created_at` (`created_at`),
  KEY `expired_at` (`expired_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

INSERT INTO admin_user VALUES('1', NOW(), NULL, NULL, NULL, 'admin', 'f6fdffe48c908deb0f4c3bd36c032e72', 'Admin');

INSERT INTO admin_group VALUES('1', NOW(), NULL, 'admin', 'enabled');

INSERT INTO admin_user_group VALUES('1', '1', NOW(), NULL);

INSERT INTO admin_page VALUES('1', NOW(), NULL, 'admin_menu', '/admin/menu', 'Menu', NULL, '{\"kind\": \"admin.menu\"}', 'publish'), ('2', NOW(), NULL, 'admin_groups', '/admin/groups', 'Groups', NULL, NULL, 'publish'), (
'3', NOW(), NULL, 'admin_users', '/admin/users', 'Users', NULL, NULL, 'publish'), ('4', NOW(), NULL, 'admin_pages', '/admin/pages', 'Pages', NULL, '{\"kind\": \"admin.pages\"}', 'publish'), ('100', NOW(), NULL, 'logout', '/user/logout', '', NULL, NULL, 'publish');

INSERT INTO admin_page_api VALUES('1', '1', '1'), ('2', '1', '2'), ('3', '1', '3'), ('4', '1', '4'), ('5', '1', '5'), ('6', '4', '6');

INSERT INTO admin_group_page VALUES('1', '1'), ('1', '2'), ('1', '3'), ('1', '4');

INSERT INTO admin_group_api VALUES('1', '1'), ('1', '2'), ('1', '3'), ('1', '4'), ('1', '5'), ('1', '6');

INSERT INTO admin_menu VALUES('1', '0', '0', 'Settings', 'settings', '0', 'active'), ('2', '1', '1', 'Menu', 'menu', '0', 'active'), ('3', '1', '2', 'Groups', 'groups', '1', 'active'), ('4', '1', '3', 'Users', 'group', '2', 'active'), ('5', '1', '4', 'Pages', 'web', '3', 'active'), ('100', '0', '100', 'Logout', 'logout', '100', 'active');

INSERT INTO admin_api VALUES('1', 'GET', '1', '/admin/menu', NULL), ('2', 'GET', '1', '/admin/menu/add', 'Add Menu'), ('3', 'GET', '1', '/admin/menu/reorder', 'Reorder Menu'), ('4', 'GET', '1', '/admin/menu/update', 'Update Menu'), ('5', 'GET', '1', '/admin/menu/status', 'Publish Menu'), ('6', 'GET', '1', '/admin/pages', 'View Pages');

