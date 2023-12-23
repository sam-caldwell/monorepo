/*
 * 072-addUserToTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function addUserToTeam(userId uuid, teamId uuid) returns integer as
$$
declare
    count         integer;
    associationId uuid;
begin
    associationId := (select createEntity('teamAssociation'::entityType));
    insert into teamMemberships (id, userId, teamId) values (associationId, userId, teamId);
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
