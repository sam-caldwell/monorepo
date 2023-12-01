/*
 * 0120-func-getWorkflowByTeamId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getWorkflowByTeamId(workflowTeamId varchar(64)) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', teamId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'description', description
        )) as workflow
    into result
    from workflow
    where ownerId == workflowTeamId;
    return result;

end ;
$$ language plpgsql;
