/*
 * 010-enableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Enable a given SQL trigger
 */
create or replace function enableTrigger(tableName text,triggerName text) returns boolean as
$$
declare
    currentState boolean:=isTriggerEnabled(tableName,triggerName);
begin
    if not currentState then
        execute format('alter table %I disable trigger %I', lower(tableName), lower(triggerName));
    end if;
    return currentState::boolean;
end;
$$ language plpgsql;
