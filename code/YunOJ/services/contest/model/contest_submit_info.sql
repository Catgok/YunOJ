CREATE TABLE contest_submit_info
(
    id         INT AUTO_INCREMENT COMMENT '竞赛提交信息ID',
    contest_id INT       NOT NULL COMMENT '竞赛ID',
    submit_id  INT       NOT NULL COMMENT '提交ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY (contest_id,submit_id),
    UNIQUE KEY (submit_id),
    FOREIGN KEY (contest_id) REFERENCES contest_info (contest_id),
    FOREIGN KEY (submit_id) REFERENCES user_submit (submit_id)
) COMMENT '竞赛提交信息表';
