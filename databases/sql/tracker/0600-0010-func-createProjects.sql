/*
 * 0020-0010-func-createProjects.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createProjects(name varchar(64), iconId uuid, workflowId uuid,
                                          description text) returns uuid as
$$
declare
    newId uuid;
begin
    newId := gen_random_uuid();
    insert into projects (id, name, iconId, workflowId, description)
    values (newId, name, iconId, workflowId, description);
    return newId;
end;
$$ language plpgsql;
