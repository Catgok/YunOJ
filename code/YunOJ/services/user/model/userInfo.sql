CREATE TABLE userInfo
(
    userid     INT AUTO_INCREMENT COMMENT '用户ID',
    username   VARCHAR(16)  NOT NULL UNIQUE COMMENT '用户名称',
    email      VARCHAR(48)  NOT NULL UNIQUE DEFAULT '' COMMENT '电子邮件地址',
    phone      varchar(18)  NOT NULL UNIQUE COMMENT '手机号码',
    gender     TINYINT      NOT NULL        DEFAULT 1 COMMENT '性别：1-男，2-女',
    user_type  TINYINT      NOT NULL        DEFAULT 2 COMMENT '用户类型：1-教练，2-队员',
    password   VARCHAR(255) NOT NULL COMMENT '密码',
    avatar     VARCHAR(255) NOT NULL        DEFAULT '' COMMENT '头像',
    status     TINYINT      NOT NULL        DEFAULT 1 COMMENT '用户状态：1-激活，2-禁用，3-待审核',
    created_at TIMESTAMP    NOT NULL        DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP    NOT NULL        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (userid),
    UNIQUE INDEX idx_name (username),
    UNIQUE INDEX idx_phone (phone),
    UNIQUE INDEX idx_email (email)
) COMMENT '用户信息表';