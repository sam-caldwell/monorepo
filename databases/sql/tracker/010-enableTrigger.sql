/*
 * 010-enableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Enable a given SQL trigger
 */
create or replace function enableTrigger(tableName text,triggerName text) returns boolean as
$$
begin
    if not isTriggerEnabled(triggerName) then
        execute format('alter table %I enable trigger %I', lower(tableName), lower(triggerName));
    end if;
    return true;
end;
$$ language plpgsql;
