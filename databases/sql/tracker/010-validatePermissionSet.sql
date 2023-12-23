/*
 * 010-validatePermissionSet.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function validatePermissionSet(pOwner permissions, pTeam permissions, pEveryone permissions)
    returns boolean as
$$
begin
    /*
     * ToDo: implement permission set validation
     */
    return true;
end;
$$
