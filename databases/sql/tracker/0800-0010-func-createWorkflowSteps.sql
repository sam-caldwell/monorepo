/*
 * 0800-0010-func-createWorkflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createWorkflowSteps(name varchar(64), iconId uuid, prevStepId uuid, nextStepId uuid,
                                               description text) returns uuid as
$$
declare
    teamId uuid;
begin
    teamId := gen_random_uuid();
    insert into users (id, name, iconId, prevStepId, nextStepId, description)
    values (teamId, name, iconId, prevStepId, nextStepId, description);
    return teamId;
end;
$$ language plpgsql;
