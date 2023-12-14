/*
 * 080-Properties.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple key-value property table
 */

create table if not exists propertyKeys
(
    id      uuid primary key not null,
    name    varchar(64)      not null,
    -- --
    created timestamp        not null default now(),
        foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * stringProperties
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create table if not exists stringProperties
(
    id    uuid primary key not null,
    value text             not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * numericProperties
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create table if not exists numericProperties
(
    id    uuid primary key not null,
    value integer          not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index ndxPropertyKeysName on propertykeys (name);
create unique index ndxPropertyKeysCreated on propertykeys (created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createPropertyKey()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createPropertyKey(propertyName varchar(64)) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId:=gen_random_uuid();
    insert into propertyKeys (id, name) values (propertyId, propertyName);
    return propertyId;
end
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createIntegerProperty()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createIntegerProperty(propertyName varchar(64), propertyValue integer) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := (select createPropertyKey(propertyName) as id);

    insert into numericProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createStringProperty()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createStringProperty(propertyName varchar(64), propertyValue text) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := (select createPropertyKey(propertyName) as id);
    insert into stringProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deletePropertyByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deletePropertyByName(propertyName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    delete from propertyKeys where id=(select id from propertyKeys where name=propertyName);
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getIntegerProperty()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getIntegerProperty(propertyName varchar(64)) returns integer as
$$
declare
    result integer;
begin
    select value as v into result
    from numericProperties
    where id in (select id
                 from propertyKeys
                 where name = propertyName);
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getStringProperty()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getStringProperty(propertyName varchar(64)) returns text as
$$
declare
    result text;
begin
    select value as v into result
    from stringProperties
    where id in (select id
                 from propertyKeys
                 where name = propertyName);
    return result;
end;
$$ language plpgsql;
