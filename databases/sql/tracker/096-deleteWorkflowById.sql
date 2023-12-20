/*
 * 060-deleteWorkflowById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create teams of users.
 */
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowById(thisWorkflowId uuid) returns integer as
$$
declare
    count integer;
begin
    if deleteWorkflowPreCheck(thisWorkflowId) then
        delete from workflowsteps where workflowid = thisWorkflowId;
        delete from workflows where id = thisWorkflowId;
        get diagnostics count = ROW_COUNT;
        return count;
    else
        return 0; -- error state
    end if;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowPreCheck(thisWorkflowId uuid) returns boolean as
$$
declare
    result boolean := true;
begin
    /*
     * This is currently a placeholder.  This function should be overloaded later
     * once ticketTypes are defined so that we can prevent any workflow from being
     * deleted until all related ticketTypes are removed.
     */
    return result;
end;
$$ language plpgsql;

