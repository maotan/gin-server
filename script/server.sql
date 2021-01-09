CREATE TABLE `user` (
    `id` bigint NOT NULL COMMENT '用户id',
    `account` varchar(64) NOT NULL COMMENT '账号',
    `mobile` varchar(11) NOT NULL COMMENT '手机号',
    `nick` varchar(128) DEFAULT NULL COMMENT '昵称',
    `password` varchar(64) DEFAULT NULL COMMENT '密码',
    `status` tinyint DEFAULT '0' COMMENT '状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;