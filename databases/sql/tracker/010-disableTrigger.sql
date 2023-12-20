/*
 * 010-disableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Disable a given SQL trigger
 */
create or replace function disableTrigger(tableName text,triggerName text) returns void as
$$
declare
    currentState boolean;
begin
    if isTriggerEnabled(tableName,triggerName) then
    -- Check if the trigger is currently enabled
    select trg.tgenabled = 'O' into currentState
    FROM pg_trigger trg
             JOIN pg_class tbl ON trg.tgrelid = tbl.oid
    WHERE trg.tgname = triggerName
      AND tbl.relname = tableName;

    -- If the trigger is currently enabled, temporarily disable it
    IF currentState THEN
        EXECUTE format('ALTER TABLE %I ENABLE TRIGGER %I', tableName, triggerName);
    END IF;
    end if;
end;
$$ language plpgsql;
