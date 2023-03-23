CREATE DATABASE `arp` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;


CREATE TABLE arp.rp_user
(
    username   varchar(25) NULL,
    avatar     varchar(200) NULL,
    email      varchar(100) NULL,
    full_name  varchar(100) NULL,
    updated_at TIMESTAMP NULL,
    status     BOOL NULL,
    source     varchar(15) NULL,
    password    varchar(200) NULL,
    CONSTRAINT user_PK PRIMARY KEY (username)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;
