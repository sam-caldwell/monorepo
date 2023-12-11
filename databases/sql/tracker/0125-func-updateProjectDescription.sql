/*
 * 0125-func-updateProjectDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectDescription(projectId uuid, descriptiveText text) returns integer as
$$
declare
    count integer;
begin
    update projects set description=descriptiveText where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
