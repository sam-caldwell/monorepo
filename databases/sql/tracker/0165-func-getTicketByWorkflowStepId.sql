/*
 * 0165-func-getTicketByWorkflowStepId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTicketByWorkflowStepId(stepId uuid, pageLimit integer,
                                                     pageOffset integer) returns jsonb as
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
    where workflowStepId == stepId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
