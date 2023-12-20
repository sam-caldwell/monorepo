/*
 * 010-isTriggerEnabled.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * return whether the given trigger is enabled
 */
create or replace function isTriggerEnabled(triggerName text) returns boolean as
$$
declare
    currentState varchar;
begin
    currentState := (select tgenabled
                     from pg_trigger
                     where tgname = lower(triggerName)
                     limit 1);
    return currentState = 'O';
end;
$$ language plpgsql;
