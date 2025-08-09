-- Tags table
CREATE TABLE IF NOT EXISTS `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'tag name',
  `color` varchar(20) DEFAULT '#666666' COMMENT 'tag color',
  `is_active` tinyint(1) DEFAULT 1 COMMENT 'is active',
  `sort_order` int(11) DEFAULT 0 COMMENT 'sort order',
  `created_at` bigint(20) DEFAULT 0 COMMENT 'created at',
  `updated_at` bigint(20) DEFAULT 0 COMMENT 'updated at',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tags table';

/*
顺路拼车更便捷 
运动搭子一起嗨 
发现身边的美食 
拼餐省钱更有趣 
租房互助找好房 
闲置好物低价转让
*/
insert into tags (name, color, is_active, sort_order) values ('顺路拼车更便捷', '#666666', 1, 1);
insert into tags (name, color, is_active, sort_order) values ('运动搭子一起嗨', '#666666', 1, 2);
insert into tags (name, color, is_active, sort_order) values ('发现身边的美食', '#666666', 1, 3);
insert into tags (name, color, is_active, sort_order) values ('拼餐省钱更有趣', '#666666', 1, 4);
insert into tags (name, color, is_active, sort_order) values ('租房互助找好房', '#666666', 1, 5);
insert into tags (name, color, is_active, sort_order) values ('闲置好物低价转让', '#666666', 1, 6);


-- 发帖表
CREATE TABLE IF NOT EXISTS `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `post_tags` varchar(255) NOT NULL DEFAULT '' COMMENT '发帖标签',
  `content` text NOT NULL COMMENT '发帖内容',
  `images` text  NULL COMMENT '图片URL列表，JSON格式',
  `latitude` decimal(10,6) DEFAULT 0 NOT NULL COMMENT '纬度',
  `longitude` decimal(10,6) DEFAULT 0 NOT NULL COMMENT '经度',
  `like_count` int(11) DEFAULT 0 NOT NULL COMMENT '点赞数量',
  `is_active` tinyint(1) DEFAULT 1 NOT NULL COMMENT '是否匿名',
  `created_at` bigint(20) DEFAULT 0 NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) DEFAULT 0 NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_tags` (`post_tags`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='发帖表';

-- 发帖点赞表
CREATE TABLE IF NOT EXISTS `post_likes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `post_id` int(11) NOT NULL DEFAULT 0 COMMENT '发帖ID',
  `is_like` tinyint(1) DEFAULT 0 NOT NULL COMMENT '是否点赞 0 否 1 是',
  `created_at` bigint(20) DEFAULT 0 NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_post_id` (`post_id`),
  UNIQUE KEY `uk_user_post` (`user_id`, `post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='发帖点赞表';

-- 评分表
CREATE TABLE IF NOT EXISTS `ratings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `post_id` int(11) NOT NULL DEFAULT 0 COMMENT '发帖ID',
  `rating` int(11) NOT NULL DEFAULT 0 COMMENT '评分',
  `created_at` bigint(20) DEFAULT 0 NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) DEFAULT 0 NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评分表';

-- 评分平均值表
CREATE TABLE IF NOT EXISTS `ratings_ave` (
  `post_id` int(11) NOT NULL DEFAULT 0 COMMENT '发帖ID',
  `rating_ave` decimal(10,2) NOT NULL DEFAULT 0 COMMENT '评分平均值',
  `rating_count` int(11) NOT NULL DEFAULT 0 COMMENT '评分总人数',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评分平均值表';

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar_id` int(11) NOT NULL DEFAULT 0 COMMENT '头像ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 图片表
CREATE TABLE IF NOT EXISTS `imgs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `base64` text NOT NULL COMMENT '图片base64',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图片表';