

CREATE TABLE IF NOT EXISTS `t_session_info` (
                                                session_id INT NOT NULL AUTO_INCREMENT COMMENT '会话号',
                                                user_id VARCHAR(50) NOT NULL COMMENT '用户身份号',
    title VARCHAR(50) NOT NULL COMMENT '会话标题',
    session_status INT NOT NULL DEFAULT 1 COMMENT '会话状态 1/2/3',
    create_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`session_id`)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `t_session_detail` (
                                                  user_id VARCHAR(50) NOT NULL COMMENT '用户身份号',
                                                  session_id INT NOT NULL  COMMENT '会话号',
                                                  question TEXT NOT NULL COMMENT '问题',
                                                  answer TEXT NOT NULL COMMENT '答案',
                                                  msgSize INT NOT NULL COMMENT '消息大小',
                                                  create_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
