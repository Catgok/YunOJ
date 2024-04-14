CREATE TABLE contest_rank_info
(
    id              INT AUTO_INCREMENT COMMENT '竞赛排名信息ID',
    contest_id      INT       NOT NULL COMMENT '竞赛ID',
    user_id         INT       NOT NULL COMMENT '用户ID',
    problem_id      INT       NOT NULL COMMENT '题目ID',
    submit_times    INT       NOT NULL DEFAULT 0 COMMENT '提交次数',
    try_times       INT       NOT NULL DEFAULT 0 COMMENT '尝试次数',
    is_pass         BOOL      NOT NULL DEFAULT FALSE COMMENT '是否通过',
    first_pass_time TIMESTAMP COMMENT '首次通过时间',
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY (contest_id, user_id, problem_id)
) COMMENT '竞赛排名信息表';
