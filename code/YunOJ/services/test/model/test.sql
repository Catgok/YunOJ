create table test
(
    id       int auto_increment comment '主键',
    info     varchar(128)   not null,
    intVal   int            not null,
    floatVal decimal(12, 4) not null,
    primary key (id)
);
