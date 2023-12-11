/*
 * 0125-func-deleteProjects.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteProjects(projectId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from projects where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
