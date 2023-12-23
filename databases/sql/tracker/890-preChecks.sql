/*
 * 890-preChecks.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * these are overloaded pre-check functions used before some operations
 */
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteAvatarPreCheck() function
 *      avatars cannot be deleted if any users exist.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteAvatarPreCheck(entityId uuid) returns boolean as
$$
begin
    return (select count(id) from users where avatarId = entityId) = 0;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteIconPreCheck() function
 *      icons cannot be deleted if any of the following exist using the icon...
 *          - team
 *          - workflow
 *          - project
 *          - ticket
 *          - ticketType
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteIconPreCheck(entityId uuid) returns boolean as
$$
begin
    return ((select count(id) from teams where iconId = entityId) = 0) and
           ((select count(id) from workflows where iconId = entityId) = 0);
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteUsersByIdPreCheck() function
 *      users cannot be deleted if any related records exist in...
 *          - teams
 *          - teamMemberships (associating users and teams)
 *          - workflows
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteUsersByIdPreCheck(entityId uuid) returns boolean as
$$
begin
    return ((select count(id) from teams where ownerId = entityId) = 0) and
           ((select count(id) from teammemberships where userId = entityId)=0) and
           ((select count(id) from workflows where iconId = entityId) = 0);
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTeamPreCheck() function
 *      teams cannot be deleted if any related records exist in...
 *          - workflows
 *          - projects
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTeamPreCheck(entityId uuid) returns boolean as
$$
begin
    /*
     * This is currently a placeholder.  This function should be overloaded later
     * as record consumers are defined so that we can prevent any record here from being
     * deleted until all related records are removed.
     */
    return true;
end;
$$ language plpgsql;
