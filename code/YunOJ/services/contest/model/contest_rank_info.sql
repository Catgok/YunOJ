CREATE TABLE contest_rank_info
(
    contest_rank_id INT AUTO_INCREMENT COMMENT '竞赛排名信息ID',
    contest_id      INT         NOT NULL COMMENT '竞赛ID',
    user_id         INT         NOT NULL COMMENT '用户ID',
    user_name       VARCHAR(16) NOT NULL COMMENT '用户名称',
    problem_id      INT         NOT NULL COMMENT '题目ID',
    submit_times    INT         NOT NULL DEFAULT 0 COMMENT '提交次数',
    try_times       INT         NOT NULL DEFAULT 0 COMMENT '尝试次数',
    is_pass         BOOL        NOT NULL DEFAULT FALSE COMMENT '是否通过',
    first_pass_time INT COMMENT '首次通过时间 单位s',
    created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (contest_rank_id),
    UNIQUE KEY (contest_id, user_id, problem_id),
    INDEX (contest_id),
    INDEX (user_id),
    INDEX (problem_id),
    FOREIGN KEY (contest_id) REFERENCES contest_info (contest_id),
    FOREIGN KEY (user_id) REFERENCES user_info (userId),
    FOREIGN KEY (problem_id) REFERENCES problem (problem_id)
) COMMENT '竞赛排名信息表';
