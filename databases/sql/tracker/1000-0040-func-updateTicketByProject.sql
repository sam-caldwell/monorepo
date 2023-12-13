/*
 * 0165-func-updateTicketProject.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketProject(ticketId uuid, projectId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticket set projectId=projectId where id=ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
