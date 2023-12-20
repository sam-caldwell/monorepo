/*
 * 890-preChecks.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * these are overloaded pre-check functions used before some operations
 */
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteAvatarPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteAvatarPreCheck(entityId uuid) returns boolean as
$$
begin
    return (select count(id) from users where avatarId = entityId) == 0;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteIconPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteIconPreCheck(entityId uuid) returns boolean as
$$
begin
    /*
     * This is currently a placeholder.  This function should be overloaded later
     * as record consumers are defined so that we can prevent any record here from being
     * deleted until all related records are removed.
     */
    return ((select count(id) from teams where iconId = entityId) == 0) and
           ((select count(id) from workflows where iconId = entityId) == 0);
end;
$$ language plpgsql;
