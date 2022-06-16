-- gb_demo -- 
CREATE DATABASE IF NOT EXISTS `gb_base` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_base`;

CREATE TABLE IF NOT EXISTS `story` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `title` varchar(255) NOT NULL DEFAULT '',
            `content` text NOT NULL,
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `seq_user` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- gb_user_shard_1 --

CREATE DATABASE IF NOT EXISTS `gb_user_shard_1` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_user_shard_1`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '', 
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- gb_user_shard_2 --

CREATE DATABASE IF NOT EXISTS `gb_user_shard_2` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_user_shard_2`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- gb_user_shard_3 --

CREATE DATABASE IF NOT EXISTS `gb_user_shard_3` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_user_shard_3`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- gb_user_shard_4 --

CREATE DATABASE IF NOT EXISTS `gb_user_shard_4` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_user_shard_4`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- gb_log_shard_1 --

CREATE DATABASE IF NOT EXISTS `gb_log_shard_1` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_log_shard_1`;

-- gb_log_shard_2 --

CREATE DATABASE IF NOT EXISTS `gb_log_shard_2` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `gb_log_shard_2`;