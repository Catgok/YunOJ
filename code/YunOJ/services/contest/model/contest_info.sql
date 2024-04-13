CREATE TABLE contest_info
(
    contest_id   INT AUTO_INCREMENT COMMENT '竞赛ID',
    contest_name VARCHAR(100) NOT NULL UNIQUE COMMENT '竞赛名称',
    description  TEXT COMMENT '竞赛描述',
    start_time   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '竞赛开始时间',
    end_time     DATETIME COMMENT '竞赛结束时间',
    create_time  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (contest_id)
) COMMENT '竞赛信息表';