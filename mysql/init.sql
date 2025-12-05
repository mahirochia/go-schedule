-- 创建数据库
CREATE DATABASE IF NOT EXISTS `FilmSite` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `FilmSite`;

-- 创建日程表
CREATE TABLE IF NOT EXISTS `schedule` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '日程ID',
    `user_id` BIGINT NOT NULL DEFAULT 0 COMMENT '用户ID',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `year` SMALLINT NOT NULL DEFAULT 1970 COMMENT '年',
    `month` TINYINT NOT NULL DEFAULT 1 COMMENT '月(1-12)',
    `day` TINYINT NOT NULL DEFAULT 1 COMMENT '日(1-31)',
    `start_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
    `end_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
    `content` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '日程安排内容',
    `priority` TINYINT NOT NULL DEFAULT 0 COMMENT '优先级(0-低,1-中,2-高)',
    `status` INT NOT NULL DEFAULT 1 COMMENT '状态：1-未开始，2-进行中，3-已结束，4-已完成',
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_year_month_day` (`year`, `month`, `day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='日程表';

CREATE TABLE `news` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '新闻ID',
                        `news_id` varchar(100) NOT NULL COMMENT '新闻唯一标识',
                        `title` varchar(500) NOT NULL COMMENT '新闻标题',
                        `creator` varchar(100) DEFAULT NULL COMMENT '作者',
                        `source` varchar(200) DEFAULT NULL COMMENT '新闻来源',
                        `content` longtext COMMENT '新闻内容',
                        `publish_time` datetime DEFAULT NULL COMMENT '新闻发布时间',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uk_news_id` (`news_id`),
                        KEY `idx_publish_time` (`publish_time`),
                        KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='新闻数据表';