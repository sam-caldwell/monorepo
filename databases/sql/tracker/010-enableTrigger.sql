/*
 * 010-enableTrigger.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Enable a given SQL trigger
 */
create or replace function enableTrigger(tableName text,triggerName text) returns void as
$$
declare
    currentState boolean;
begin
    if not isTriggerEnabled(tableName,triggerName) then
        select trg.tgenabled = 'O' into currentState
        FROM pg_trigger trg
                 JOIN pg_class tbl ON trg.tgrelid = tbl.oid
        WHERE trg.tgname = triggerName
          AND tbl.relname = tableName;

        -- If the trigger is currently enabled, temporarily disable it
        IF currentState THEN
            EXECUTE format('ALTER TABLE %I DISABLE TRIGGER %I', tableName, triggerName);
        END IF;
    end if;
end;
$$ language plpgsql;
