-- 用户表
CREATE TABLE `user` (
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                        `username` VARCHAR(50) NOT NULL COMMENT '用户名',
                        `password` VARCHAR(255) NOT NULL COMMENT '密码',
                        `phone` VARCHAR(20) NOT NULL COMMENT '手机号',
                        `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
                        `real_name` VARCHAR(50) DEFAULT NULL COMMENT '真实姓名',
                        `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像',
                        `gender` TINYINT(1) DEFAULT '0' COMMENT '性别:0-未知,1-男,2-女',
                        `age` TINYINT UNSIGNED DEFAULT NULL COMMENT '年龄',
                        `occupation` VARCHAR(50) DEFAULT NULL COMMENT '职业',
                        `income_level` TINYINT UNSIGNED DEFAULT NULL COMMENT '收入水平:1-5级',
                        `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uniq_username` (`username`),
                        UNIQUE KEY `uniq_phone` (`phone`),
                        KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 用户偏好表
CREATE TABLE `user_preference` (
                                   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                                   `district_preference` VARCHAR(500) DEFAULT NULL COMMENT '区域偏好',
                                   `price_min` DECIMAL(10,2) DEFAULT NULL COMMENT '最低价格',
                                   `price_max` DECIMAL(10,2) DEFAULT NULL COMMENT '最高价格',
                                   `house_type` VARCHAR(100) DEFAULT NULL COMMENT '户型偏好',
                                   `area_min` DECIMAL(6,2) DEFAULT NULL COMMENT '最小面积',
                                   `area_max` DECIMAL(6,2) DEFAULT NULL COMMENT '最大面积',
                                   `tags` VARCHAR(500) DEFAULT NULL COMMENT '标签偏好',
                                   `facilities` VARCHAR(500) DEFAULT NULL COMMENT '设施要求',
                                   `commuting_time` INT UNSIGNED DEFAULT NULL COMMENT '通勤时间要求(分钟)',
                                   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `uniq_user_id` (`user_id`),
                                   CONSTRAINT `fk_user_preference_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户偏好表';

-- 房源表
CREATE TABLE `house` (
                         `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `title` VARCHAR(100) NOT NULL COMMENT '房源标题',
                         `description` TEXT COMMENT '房源描述',
                         `price` DECIMAL(10,2) NOT NULL COMMENT '月租金',
                         `area` DECIMAL(6,2) NOT NULL COMMENT '面积(㎡)',
                         `house_type` VARCHAR(20) NOT NULL COMMENT '户型',
                         `floor` VARCHAR(20) DEFAULT NULL COMMENT '楼层',
                         `total_floors` TINYINT UNSIGNED DEFAULT NULL COMMENT '总楼层',
                         `district` VARCHAR(50) NOT NULL COMMENT '区域',
                         `address` VARCHAR(200) NOT NULL COMMENT '详细地址',
                         `longitude` DECIMAL(10,7) DEFAULT NULL COMMENT '经度',
                         `latitude` DECIMAL(10,7) DEFAULT NULL COMMENT '纬度',
                         `facilities` VARCHAR(500) DEFAULT NULL COMMENT '设施',
                         `tags` VARCHAR(500) DEFAULT NULL COMMENT '标签',
                         `main_image` VARCHAR(255) DEFAULT NULL COMMENT '主图',
                         `images` TEXT COMMENT '图片列表',
                         `contact_name` VARCHAR(50) NOT NULL COMMENT '联系人',
                         `contact_phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
                         `status` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-待租,2-已租,3-下架',
                         `view_count` INT UNSIGNED DEFAULT '0' COMMENT '浏览数',
                         `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
                         PRIMARY KEY (`id`),
                         KEY `idx_district` (`district`),
                         KEY `idx_price` (`price`),
                         KEY `idx_house_type` (`house_type`),
                         KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源表';

-- 房源特色表
CREATE TABLE `house_feature` (
                                 `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                                 `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                                 `feature_type` VARCHAR(50) NOT NULL COMMENT '特色类型',
                                 `feature_value` VARCHAR(100) NOT NULL COMMENT '特色值',
                                 `weight` TINYINT UNSIGNED DEFAULT '1' COMMENT '权重(1-10)',
                                 `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_house_id` (`house_id`),
                                 KEY `idx_feature_type` (`feature_type`),
                                 CONSTRAINT `fk_house_feature_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源特色表';

-- 用户行为表
CREATE TABLE `user_behavior` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                 `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                                 `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                                 `behavior_type` TINYINT UNSIGNED NOT NULL COMMENT '行为类型:1-浏览,2-收藏,3-联系,4-预约看房',
                                 `duration` INT UNSIGNED DEFAULT NULL COMMENT '浏览时长(秒)',
                                 `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '行为时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_house_id` (`house_id`),
                                 KEY `idx_behavior_type` (`behavior_type`),
                                 KEY `idx_created_at` (`created_at`),
                                 CONSTRAINT `fk_user_behavior_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
                                 CONSTRAINT `fk_user_behavior_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户行为表';

-- 收藏表
CREATE TABLE `favorite` (
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                            `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                            `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `uniq_user_house` (`user_id`, `house_id`),
                            KEY `idx_user_id` (`user_id`),
                            KEY `idx_house_id` (`house_id`),
                            CONSTRAINT `fk_favorite_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
                            CONSTRAINT `fk_favorite_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收藏表';

-- 推荐记录表
CREATE TABLE `recommendation` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                                  `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                                  `strategy` VARCHAR(50) NOT NULL COMMENT '推荐策略',
                                  `score` DECIMAL(5,4) NOT NULL COMMENT '推荐分数',
                                  `rank` INT UNSIGNED NOT NULL COMMENT '推荐排名',
                                  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '推荐时间',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_user_id` (`user_id`),
                                  KEY `idx_house_id` (`house_id`),
                                  KEY `idx_strategy` (`strategy`),
                                  KEY `idx_created_at` (`created_at`),
                                  CONSTRAINT `fk_recommendation_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
                                  CONSTRAINT `fk_recommendation_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='推荐记录表';

-- 预约看房表
CREATE TABLE `appointment` (
                               `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                               `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                               `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                               `appointment_time` DATETIME NOT NULL COMMENT '预约时间',
                               `contact_phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
                               `status` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-待确认,2-已确认,3-已取消,4-已完成',
                               `note` VARCHAR(200) DEFAULT NULL COMMENT '备注',
                               `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               KEY `idx_user_id` (`user_id`),
                               KEY `idx_house_id` (`house_id`),
                               KEY `idx_status` (`status`),
                               CONSTRAINT `fk_appointment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
                               CONSTRAINT `fk_appointment_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预约看房表';

-- 反馈表
CREATE TABLE `feedback` (
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                            `house_id` INT UNSIGNED NOT NULL COMMENT '房源ID',
                            `type` TINYINT UNSIGNED NOT NULL COMMENT '反馈类型:1-房源信息错误,2-推荐不准确,3-其他',
                            `content` TEXT NOT NULL COMMENT '反馈内容',
                            `status` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-待处理,2-已处理',
                            `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_user_id` (`user_id`),
                            KEY `idx_house_id` (`house_id`),
                            KEY `idx_status` (`status`),
                            CONSTRAINT `fk_feedback_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
                            CONSTRAINT `fk_feedback_house` FOREIGN KEY (`house_id`) REFERENCES `house` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='反馈表';