/*
 * 900-logging.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'logPriority') THEN
            CREATE TYPE logPriority AS ENUM ('critical', 'warn', 'info', 'debug');
        END IF;
    END
$$;
create table if not exists logs
(
    id       serial4       not null primary key,
    priority logPriority   not null default 'info',
    userId   uuid          not null,
    message  varchar(2048) not null,
    -- --
    created  timestamp        not null default now(),
    foreign key (userId) references users (id)
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index ndxLogsPriority on logs(priority);
create index ndxLogsUserId on logs(userId);
create index ndxLogsMessage on logs(message);
create index ndxLogsCreated on logs(created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createworkflows()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace trigger preventEntityDelete
    before delete
    on logs
    for each row
execute function preventDelete();
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * writeLog()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function writeLog(p logPriority, uid uuid, msg text) returns bool as
$$
begin
    insert into logs(priority, userId, message) values (p, uid, msg);
    return true;
end;
$$ language plpgsql;
