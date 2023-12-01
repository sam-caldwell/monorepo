/*
 * 0099-table-logs.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A table containing all logs for history.
 */
create table if not exists logs
(
    id       serial4       not null primary key,
    priority logPriority   not null default 'info',
    userId   uuid          not null,
    message  varchar(2048) not null,
    foreign key (userId) references users (id)
);
