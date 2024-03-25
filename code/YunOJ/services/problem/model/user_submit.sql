CREATE TABLE user_submit
(
    submit_id   INT AUTO_INCREMENT COMMENT '提交id',
    user_id     INT       NOT NULL COMMENT '用户id',
    problem_id  INT       NOT NULL COMMENT '题目id',
    code        TEXT      NOT NULL COMMENT '代码',
    language    TINYINT   NOT NULL DEFAULT 1 COMMENT '语言 1-c++',
    result      TINYINT   NOT NULL DEFAULT 0 COMMENT '结果',
    time        INT       NOT NULL DEFAULT 0 COMMENT '运行时间 单位ms',
    memory      INT       NOT NULL DEFAULT 0 COMMENT '内存占用 单位MB',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (submit_id),
    INDEX user_id_index (user_id),
    INDEX problem_id_index (problem_id),
    INDEX user_id_problem_id_index (user_id, problem_id),
    FOREIGN KEY (user_id) REFERENCES user_info (userid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (problem_id) REFERENCES problem (problem_id) ON DELETE CASCADE ON UPDATE CASCADE
) COMMENT '用户提交题目信息表';
