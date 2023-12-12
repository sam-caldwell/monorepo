/*
 * 0109-func-createWorkflow.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createWorkflow(name varchar(64), iconId uuid, ownerId uuid, teamId uuid, owner permissions,
                                          team permissions, everyone permissions, description text) returns uuid as
$$
declare
    workflowId uuid;
begin
    workflowId := gen_random_uuid();
    insert into workflow (id, name, iconId, ownerId, teamId, owner, team, everyone, description)
    values (workflowId, name, iconId, ownerId, teamId, owner, team, everyone, description);
    return workflowId;
end;
$$ language plpgsql;
