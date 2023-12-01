/*
 * 0105-func-updateTicketTypeWorkflowId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketTypeWorkflowId(typeId uuid, workflowId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set workflowId=workflowId where id=typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
