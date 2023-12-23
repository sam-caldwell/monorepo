/*
 * 152-getTicketsByPermission.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTicketsByPermission(o permissions, t permissions, e permissions,
                                                  recordCount integer default 1000,
                                                  startAt integer default 0) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorid,
            'assignedUserId', assignedid,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', summary,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from tickets
    where isPermissible(owner, o)
       or isPermissible(team, t)
       or isPermissible(everyone, e)
    limit recordCount offset startAt;
    return result;
end;
$$ language plpgsql;
