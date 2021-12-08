use db_read;

drop table if exists tb_user;
create table tb_user(
	acc varchar(32) primary key comment '',
    pwd varchar(128) not null comment '',
    nickname varchar(32) comment ''
) comment = '用户列表';

drop table if exists tb_admin;
create table tb_admin(
	acc varchar(32) primary key comment '',
    pwd varchar(128) not null comment '',
    nickname varchar(32) comment ''
) comment = '管理员列表';

#权限暂时定为，文章CRUD，文章单词CRUD，普通管理员可通过注册申请。超级管理员可管理普通管理员（由站长使用）。
drop table if exists tb_admin_power;
create table tb_admin_power(
	acc varchar(32) not null,
    power varchar(32) not null,
    primary key (acc,power)
) comment = '管理员权限表';

drop table if exists tb_admin_book;
create table tb_admin_book(
	acc varchar(32) not null,
    book int not null,
    primary key (acc,book)
) comment = '管理员可管理文章列表';

drop table if exists tb_book;
create table tb_book(
	id int auto_increment primary key,
    fullname varchar(64) not null,
	shortname varchar(32) null,
    remark varchar(256) null,
    picture varchar(256) null
) comment = '文章';


drop table if exists tb_chapter;
create table tb_chapter(
	id int auto_increment primary key,
    book int not null,
    fullname varchar(64) not null,
	shortname varchar(32) null,
    remark varchar(256) null,
    picture varchar(256) null
) comment = '章节';

drop table if exists tb_paragraph;
create table tb_paragraph(
	id int auto_increment primary key,
    book int not null,
    chapter int not null,
    content text,
    picture varchar(256) null
) comment = '段落';

drop table if exists tb_word_list;
create table tb_word_list(
	id int auto_increment primary key,
    book int not null,
    chapter int not null,
    paragraph int not null,
    word varchar(64),
    sound varchar(256) null comment '发音'
) comment = '生词表，通用';

 drop table if exists tb_user_word_list;
create table tb_user_word_list(
	id int auto_increment primary key,
    user_id int not null,
    book int not null,
    chapter int not null,
    paragraph int not null,
    word varchar(64),
    sound varchar(256) null comment '发音'
) comment = '生词表，用户';




