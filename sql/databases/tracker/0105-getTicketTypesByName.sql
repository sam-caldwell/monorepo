/*
 * 0105-func-getTicketTypesByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getTicketTypesByName(typeName varchar(64)) returns jsonb as
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
    where name == typeName;
    return result;

end ;
$$ language plpgsql;
