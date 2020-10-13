-- demo -- 
CREATE DATABASE IF NOT EXISTS `demo` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `demo`;

CREATE TABLE IF NOT EXISTS `story` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `title` varchar(255) NOT NULL DEFAULT '',
            `content` text NOT NULL,
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- user_shard_1 --

CREATE DATABASE IF NOT EXISTS `user_shard_1` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `user_shard_1`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '', 
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- user_shard_2 --

CREATE DATABASE IF NOT EXISTS `user_shard_2` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `user_shard_2`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- user_shard_3 --

CREATE DATABASE IF NOT EXISTS `user_shard_3` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `user_shard_3`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- user_shard_4 --

CREATE DATABASE IF NOT EXISTS `user_shard_4` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `user_shard_4`;

CREATE TABLE IF NOT EXISTS `user_info` (
            `id` int(10) NOT NULL AUTO_INCREMENT,
            `name` varchar(255) NOT NULL DEFAULT '',         
            `dtime` int(10) NOT NULL DEFAULT 0,
            PRIMARY KEY (`id`),
            KEY `dtime` (`dtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

-- log_shard_1 --

CREATE DATABASE IF NOT EXISTS `log_shard_1` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `log_shard_1`;

-- log_shard_2 --

CREATE DATABASE IF NOT EXISTS `log_shard_2` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `log_shard_2`;