CREATE TABLE contest_user_info
(
    id           INT AUTO_INCREMENT COMMENT '竞赛用户信息ID',
    user_id      INT       NOT NULL COMMENT '用户ID',
    contest_id   INT       NOT NULL COMMENT '竞赛ID',
    join_status  TINYINT   NOT NULL DEFAULT 0 COMMENT '用户参加竞赛的状态 0:未参加 1:已参加 2:已结束',
    contest_rank INT       NOT NULL DEFAULT 0 COMMENT '用户在竞赛中的排名',
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY (contest_id, user_id),
    FOREIGN KEY (contest_id) REFERENCES contest_info (contest_id),
    FOREIGN KEY (user_id) REFERENCES user_info (userid)
) COMMENT '竞赛用户信息表';
