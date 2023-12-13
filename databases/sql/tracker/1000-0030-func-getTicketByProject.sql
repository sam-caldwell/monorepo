/*
 * 0060-0030-func-getTicketByProject.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTicketByProject(projectId uuid, pageLimit integer,
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
    where projectId == projectId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;