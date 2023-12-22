/*
 * 122-addTicketTypeToProject.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function addTicketTypeToProject(thisPid uuid, thisTid uuid) returns uuid as
$$
declare
    associationId uuid;
begin
    associationId := (select createEntity('ticketTypeAssociation'::entityType));
    insert into projectTicketTypes (id, projectId, ticketTypeId) values (associationId, thisPid, thisTid);
    return associationId;
end;
$$ language plpgsql;
