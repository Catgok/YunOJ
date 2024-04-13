CREATE TABLE contest_problem_info
(
    id         INT AUTO_INCREMENT COMMENT 'ID',
    problem_id INT       NOT NULL KEY COMMENT '题目ID',
    contest_id INT       NOT NULL COMMENT '竞赛ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY (contest_id, problem_id)
) COMMENT '竞赛题目信息表';