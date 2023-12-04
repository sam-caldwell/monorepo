/*
 * 0125-func-getTicketByAssignee.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTicketByAssignee(assignedUserId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow into result
    from ticket
    where assignedUserId == assignedUserId;
    return result;
end;
$$ language plpgsql;
