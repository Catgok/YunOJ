# todo 考虑是否需要落库！！！

CREATE TABLE judge_case
(
    case_id     INT AUTO_INCREMENT COMMENT '测试用例id',
    problem_id  INT       NOT NULL COMMENT '题目id',
    input       TEXT      NOT NULL COMMENT '输入文件路径',
    output      TEXT      NOT NULL COMMENT '输出文件路径',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (case_id),
    INDEX problem_id_index (problem_id),
    FOREIGN KEY (problem_id) REFERENCES problem (problem_id) ON DELETE CASCADE ON UPDATE CASCADE
) COMMENT '测试用例表';
