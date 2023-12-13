/*
 * 0100-func-boundsCheckLower.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function boundsCheckLower(n integer, lowerBound integer) returns boolean as
$$
begin
    if (n < lowerBound) then
        raise exception 'lower bounds check error(boundary: %)(value:%)',lowerBound,n;
    end if;
    return true;
end;
$$ language plpgsql;
