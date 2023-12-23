/*
 * 152-updateTicketAssignee.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateTicketAssignee(ticketId uuid, assignee uuid) returns integer as
$$
declare
    count integer;
begin
    update tickets set assignedUserId=assignee where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
