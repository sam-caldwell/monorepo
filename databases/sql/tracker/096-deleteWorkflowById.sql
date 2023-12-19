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
    -- ToDo: we should not be able to delete any workflows if it is mapped to a ticketType
    delete from workflows where id = thisWorkflowId;
    delete from workflowsteps where workflowid=thisWorkflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
