CREATE TABLE contest_problem_info
(
    contest_problem_id INT AUTO_INCREMENT COMMENT '竞赛题目信息表ID',
    contest_id         INT       NOT NULL COMMENT '竞赛ID',
    problem_id         INT       NOT NULL COMMENT '题目ID',
    created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (contest_problem_id),
    UNIQUE KEY (contest_id, problem_id),
    INDEX (contest_id),
    INDEX (problem_id)
) COMMENT '竞赛题目信息表';