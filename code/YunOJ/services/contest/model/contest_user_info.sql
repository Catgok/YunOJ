CREATE TABLE contest_user_info
(
    contest_user_id INT AUTO_INCREMENT COMMENT '竞赛用户信息ID',
    user_id         INT         NOT NULL COMMENT '用户ID',
    contest_id      INT         NOT NULL COMMENT '竞赛ID',
    user_name       VARCHAR(16) NOT NULL COMMENT '用户名称',
    join_status     TINYINT     NOT NULL DEFAULT 0 COMMENT '用户参加竞赛的状态 0:未参加 1:已参加 2:已结束',
    created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (contest_user_id),
    UNIQUE KEY (contest_id, user_id),
    INDEX (contest_id),
    INDEX (user_id),
    FOREIGN KEY (contest_id) REFERENCES contest_info (contest_id),
    FOREIGN KEY (user_id) REFERENCES user_info (userid)
) COMMENT '竞赛用户信息表';
