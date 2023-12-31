/*
 * 020-entity.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * An entity is a universal object identifier (uuid) for the entire tracker database.
 * This is intended as the basis of our accountability system.  The entity system provides
 * a write-only object registry with timestamps and context.
 */
create table if not exists entity
(
    id      uuid          not null primary key,
    /*
     * The context is a text description of the action the subject took on the object.
     */
    type    entityType    not null,
    created timestamp     not null default now(),
    context varchar(2048) not null default '',
    constraint prohibit_zero_uuid check (id <> '00000000-0000-0000-0000-000000000000')
);
/*
 * entity indexes
 */
create index if not exists ndxEntityType on entity (type);
create index if not exists ndxEntityContext on entity (context);
create index if not exists ndxEntityCreated on entity (created);
/*
 * ensure entity table cannot be deleted...writes only
 */
create or replace trigger preventEntityDelete
    before delete
    on entity
    for each row
execute function preventDelete();

create or replace trigger preventEntityUpdate
    before update
    on entity
    for each row
execute function preventUpdate();

