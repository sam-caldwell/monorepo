/*
 * 022-getEntity.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function getEntity(entityId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select json_build_object(
                   'id', id,
                   'type',type,
                   'created',created,
                   'context',context
               ) as entity
    into result
    from entity
    where id=entityId
    limit 1;
    return result;
end;
$$ language plpgsql;
