/*
 * 0115-func-updateProjectPermOwner
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectPermOwner(projectId uuid, owner permissions) returns integer as
$$
declare
    count integer;
begin
    update projects set owner=owner where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
