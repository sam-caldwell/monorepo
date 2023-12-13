/*
 * 0145-func-updateTicketSummary.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketDescription(ticketId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update ticket set description=description where id=ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
