/*
 * 015-isPermissionSubset.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return whether permission rhs is in the set of lhs.  See Markdown notes.
 * This is great for searching for objects which have permissions in the given subset.
 */
create or replace function isPermissionSubset(lhs permissions, rhs permissions) returns boolean as
$$
declare
    result boolean;
begin
    if lhs = 'none' then
        result = false;
    else
        case rhs
            when 'none' then result:=false;
            when 'read' then result := (lhs = 'read');
            when 'create' then result := (lhs in ('create', 'read'));
            when 'update' then result := (lhs in ('update', 'create', 'read'));
            when 'delete' then result := (lhs in ('delete', 'update', 'create', 'read'));
            else raise exception 'unsupported or unknown permission value';
            end case;
    end if;
    return result;
end;
$$ language plpgsql;
