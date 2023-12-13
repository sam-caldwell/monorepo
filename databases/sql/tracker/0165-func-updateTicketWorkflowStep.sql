/*
 * 0165-func-updateTicketWorkflowStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketWorkflowStep(ticketId uuid, workflowStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticket set workflowStepId=workflowStepId where id=ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
