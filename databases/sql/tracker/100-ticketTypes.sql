/*
 * 100-ticketTypes.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Identify the ticket classifications.
 */
create table if not exists ticketTypes
(
    id          uuid primary key not null,
    -- each workflow has a unique name --
    name        varchar(64)      not null,
    -- a uuid representing the icon for the workflow.
    iconId      uuid             not null,
    -- a project ticketType is mapped to a workflow --
    workflowId  uuid             not null,
    -- --
    created     timestamp        not null default now(),
    -- descriptive text --
    description text,
    foreign key (iconId) references icons (id),
    foreign key (workflowId) references workflows (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxTicketTypesName on tickettypes (name);
create index if not exists ndxTicketTypesCreated on tickettypes (created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createTicketType(typeName varchar(64), typeIconId uuid, wid uuid,
                                            typeDescription text) returns uuid as
$$
declare
    newId uuid;
begin
    newId := gen_random_uuid();
    insert into ticketTypes (id, name, iconId, workflowId, description)
    values (newId, typeName, typeIconId, wid, typeDescription);
    return newId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTicketTypeById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTicketTypeById(typeId uuid) returns integer as
$$
declare
    count integer;
begin
    -- We should not be able to delete a ticket type if it is in use.
    -- This means we cannot delete a ticket type if it is used in any workflow or ticket.
    delete from ticketTypes where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketTypesById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketTypesById(typeId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'name', name,
                   'iconId', iconId,
                   'workflowId', workflowId,
                   'description', description
               ) as data
    into result
    from ticketTypes
    where id == typeId
    limit 1;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketTypesByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
        )) as data
    into result
    from ticketTypes
    where name == typeName;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketType(typeId uuid, typeName varchar(64), typeIconId uuid, wid uuid,
                                            typeDescription text) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set name=typeName, iconid=typeIconId, workflowId=wid, description=typeDescription
                       where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketTypeDescription()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketTypeDescription(typeId uuid, typeDescription text) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set description=typeDescription where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketTypeIconId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketTypeIconId(typeId uuid, typeIconId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set iconid=typeIconId where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketTypeName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketTypeName(typeId uuid, typeName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set name=typeName where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketTypeWorkflowId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketTypeWorkflowId(typeId uuid, wid uuid) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set workflowId=wid where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;






