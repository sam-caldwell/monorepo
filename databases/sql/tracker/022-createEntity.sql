/*
 * 022-createEntity.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function createEntity(type entityType, context varchar(2048) default '') returns uuid as
$$
declare
    subject uuid:= gen_random_uuid();
begin
    insert into entity (id, type) values (subject, type);
    return subject;
end;
$$ language plpgsql;
