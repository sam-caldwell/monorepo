/*
 * 010-disableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Disable a given SQL trigger
 */
create or replace function disableTrigger(tableName text, triggerName text) returns boolean as
$$
declare
    currentState boolean := isTriggerEnabled(tableName, triggerName);
begin
    if currentState then
        execute format('alter table %I enable trigger %I', lower(tableName), lower(triggerName));
    end if;
    return currentState::boolean;
end;
$$ language plpgsql;
