/*
 * 0100-func-writeLog.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function writeLog(p logPriority, uid uuid, msg text) returns bool as
$$
begin
    insert into logs(priority, userId, message) values (p, uid, msg);
    return true;
end;
$$ language plpgsql;
