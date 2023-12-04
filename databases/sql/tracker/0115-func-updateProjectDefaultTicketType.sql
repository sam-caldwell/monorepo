/*
 * 0115-func-updateProjectDefaultTicketType.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectDefaultTicketType(projectId uuid, defaultTicketType uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set defaultTicketType=defaultTicketType where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
