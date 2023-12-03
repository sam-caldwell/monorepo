/*
 * 0105-func-getTicketTypesById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getTicketTypesById(typeId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'workflowId', workflowId,
            'description', description
        ))  as data into result
    from ticketTypes
    where id == typeId;
    return result;

end ;
$$ language plpgsql;
