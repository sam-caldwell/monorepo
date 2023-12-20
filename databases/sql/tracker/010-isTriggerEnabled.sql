/*
 * 010-isTriggerEnabled.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * return whether the given trigger is enabled
 */
create or replace function isTriggerEnabled(tableName text, triggerName text) returns boolean as
$$
declare
    currentState boolean;
begin
    select trg.tgenabled = 'O' into currentState
    from pg_trigger trg
             join pg_class tbl on trg.tgrelid = tbl.oid
    where trg.tgname = triggerName
      and tbl.relname = tableName;
    return currentState;
end;
$$ language plpgsql;
