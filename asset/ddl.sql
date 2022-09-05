CREATE TABLE `tally_fund_revenues`
(
    `id`                   int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `date`                 date           NOT NULL DEFAULT '1970-01-01' COMMENT '日期',
    `money`                decimal(20, 4) NOT NULL DEFAULT 0 COMMENT '金额',
    `user_id`              int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `fund_classifytion_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
    `remark`               varchar(200)   NOT NULL DEFAULT '' COMMENT '备注',
    `updated_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                    `idx_date` (`user_id`,`date`,`fund_classifytion_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账收入表';

CREATE TABLE `tally_fund_expences`
(
    `id`                   int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `date`                 date           NOT NULL DEFAULT '1970-01-01' COMMENT '日期',
    `money`                decimal(20, 4) NOT NULL DEFAULT 0 COMMENT '金额',
    `user_id`              int(11) NOT NULL DEFAULT 0 COMMENT '用户',
    `fund_classifytion_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
    `remark`               varchar(200)   NOT NULL DEFAULT '' COMMENT '备注',
    `updated_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                    `idx_date` (`user_id`,`date`,`fund_classifytion_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账支出表';


CREATE TABLE `tally_budgets`
(
    `id`                   int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `start_day`            date           NOT NULL DEFAULT '1970-01-01' COMMENT '开始时间',
    `end_day`              date           NOT NULL DEFAULT '1970-01-01' COMMENT '结束时间',
    `money`                decimal(20, 4) NOT NULL DEFAULT 0 COMMENT '金额',
    `user_id`              int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `fund_classifytion_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
    `updated_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at`           datetime       NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                    `idx_date` (`user_id`,`start_day`,`end_day`,`fund_classifytion_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账预算表';


CREATE TABLE `tally_users`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `avatar`     varchar(100) NOT NULL DEFAULT '' COMMENT '头像',
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '用户名贵',
    `updated_at` datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at` datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账用户表';

CREATE TABLE `tally_fund_classifytions`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
    `type`       varchar(100) NOT NULL DEFAULT '' COMMENT '分类类，收入/支出',
    `updated_at` datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at` datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_name` (`type`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账分类表';

CREATE TABLE `tally_fund_accounts`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `bank_id`      varchar(100) NOT NULL DEFAULT '' COMMENT '银行id',
    `bank_account` varchar(100) NOT NULL DEFAULT '' COMMENT '银行帐户',
    `account_type` varchar(100) NOT NULL DEFAULT '' COMMENT '帐户类型，储蓄还是信用',
    `updated_at`   datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `created_at`   datetime     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_bank` (`bank_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账帐户表';


CREATE TABLE `tally_banks`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
    `updated_at` datetime     NOT NULL DEFAULT 0 COMMENT '创建时间',
    `created_at` datetime     NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='记账银行表';

