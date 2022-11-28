create table if not exists files
(
    id            uuid primary key not null default gen_random_uuid(),
    file_name     text unique      not null,
    file_path     text unique      not null,
    creation_time timestamp        not null default current_timestamp,
    update_time   timestamp        not null default current_timestamp
);

create index if not exists file_name_index on files(file_name);
