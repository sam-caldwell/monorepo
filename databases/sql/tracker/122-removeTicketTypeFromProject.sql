/*
 * 122-removeTicketTypeFromProject.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function removeTicketTypeFromProject(thisPid uuid, thisTid uuid) returns integer as
$$
declare
    count integer;
begin
    delete from projectTicketTypes where projectId = thisPid and ticketTypeId = thisTid;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
