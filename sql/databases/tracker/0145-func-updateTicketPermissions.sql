/*
 * 0145-func-updateTicketSummary.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketPermissions(ticketId uuid, owner permissions, team permissions,
                                                   everyone permissions) returns integer as
$$
declare
    count integer;
begin
    update ticket set owner=owner, team=team, everyone=everyone where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
