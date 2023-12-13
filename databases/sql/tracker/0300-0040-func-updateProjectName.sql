/*
 * 0008-0040-func-updateProjectName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectName(projectId uuid, projectName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update projects set name=projectName where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
