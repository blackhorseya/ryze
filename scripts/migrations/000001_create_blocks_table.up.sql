create table if not exists blocks
(
    number      bigint unsigned not null,
    hash        binary(32)      not null,
    parent_hash binary(32)      not null,
    timestamp   timestamp(6)    not null,
    constraint pk_blocks
        primary key (hash)
) engine = innodb
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci;

create index if not exists idx_blocks_number on blocks (number);
create index if not exists idx_blocks_timestamp on blocks (timestamp);