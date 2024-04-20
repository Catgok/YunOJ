CREATE TABLE user_info
(
    userid     INT AUTO_INCREMENT COMMENT '用户ID',
    username   VARCHAR(16)  NOT NULL UNIQUE COMMENT '用户名称',
    phone      varchar(18)  NOT NULL UNIQUE COMMENT '手机号码',
    user_type  TINYINT      NOT NULL DEFAULT 2 COMMENT '用户类型：1-教练，2-队员',
    password   VARCHAR(255) NOT NULL COMMENT '密码',
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (userid),
    INDEX (created_at)
) COMMENT '用户信息表';