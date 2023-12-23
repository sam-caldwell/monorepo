/*
 * 152-getTicketByProject.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTicketByProject(thisProject uuid,
                                              recordCount integer default 1000,
                                              startAt integer default 0) returns jsonb as
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
            'subject', summary,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where projectId = thisProject
    limit recordCount offset startAt;
    return result;
end;
$$ language plpgsql;
