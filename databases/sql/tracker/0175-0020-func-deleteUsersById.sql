/*
 * 0004-0020-func-deleteUsersById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteUsersById(userId uuid) returns integer as
$$
declare
    count integer;
begin
    --Note: a user should not be deletable if it is associated with any
    --      project, ticket, workflow, team (as owner)
    delete from users where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
