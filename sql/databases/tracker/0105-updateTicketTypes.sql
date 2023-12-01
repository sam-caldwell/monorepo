/*
 * 0105-func-updateTicketType.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketType(typeId uuid, name varchar(64), iconId uuid, workflowId uuid,
                                            description text) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set name=name,iconid=iconId,workflowId=workflowId,description=description where id=typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
