CREATE TABLE problem
(
    problem_id   INT AUTO_INCREMENT COMMENT '题目id',
    title        VARCHAR(48) NOT NULL UNIQUE COMMENT '题目名称',
    time_limit   INT         NOT NULL DEFAULT 1000 COMMENT '时间限制 单位ms',
    memory_limit INT         NOT NULL DEFAULT 128 COMMENT '空间限制 单位MB',
    description  TEXT        NOT NULL COMMENT '题目描述',
    hard_level   TINYINT     NOT NULL DEFAULT 0 COMMENT '难度等级 0-简单,1-中等,2-困难',
    is_delete    TINYINT     NOT NULL DEFAULT 0 COMMENT '是否删除 0-否,1-是',
    submit_count INT         NOT NULL DEFAULT 0 COMMENT '提交次数',
    pass_count   INT         NOT NULL DEFAULT 0 COMMENT '通过次数',
    create_time  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (problem_id),
    INDEX (title),
    INDEX (create_time)
) COMMENT '题目信息表' AUTO_INCREMENT = 1001;
