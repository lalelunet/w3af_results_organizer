BEGIN TRANSACTION;
CREATE TABLE "vulns" (
	`vuln_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`vuln_url`	TEXT UNIQUE,
	`vuln_request`	TEXT,
	`vuln_description`	TEXT,
	`vuln_categorie`	INTEGER,
	`vuln_severity`	INTEGER,
	`vuln_project`	INTEGER,
	`vuln_state`	INTEGER,
	`vuln_comment`	TEXT,
	`vuln_scan`	INTEGER,
	`vuln_plugin`	INTEGER
);
CREATE TABLE `severity` (
	`severity_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`severity_name`	TEXT
);
INSERT INTO `severity` VALUES (1,'High'),
 (2,'Medium'),
 (3,'Low'),
 (4,'Information');
CREATE TABLE `scans` (
	`scan_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`scan_project`	INTEGER,
	`scan_date`	TEXT
);
CREATE TABLE `projects` (
	`project_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`project_name`	TEXT,
	`project_description`	TEXT
);
CREATE TABLE `plugins` (
	`plugins_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`plugins_categories`	INTEGER,
	`plugins_name`	TEXT UNIQUE,
	`plugins_description`	TEXT
);
CREATE TABLE "categories" (
	`categories_id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`categories_name`	TEXT,
	`categories_description`	INTEGER
);
INSERT INTO `categories` VALUES (1,'audit',NULL),
 (2,'infrastructure',NULL),
 (3,'bruteforce',NULL),
 (4,'grep',NULL),
 (5,'evasion',NULL),
 (6,'output',NULL),
 (7,'crawl',NULL),
 (8,'auth',NULL);
COMMIT;
