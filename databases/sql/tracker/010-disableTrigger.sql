/*
 * 010-disableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Disable a given SQL trigger
 */
create or replace function disableTrigger(tableName text, triggerName text) returns boolean as
$$
begin
    if isTriggerEnabled(triggerName) then
        execute format('alter table %I disable trigger %I', lower(tableName), lower(triggerName));
    end if;
    return true;
end;
$$ language plpgsql;
