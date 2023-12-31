/*
 * 010-boundsCheckLower.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function boundsCheckLower(n integer, lowerBound integer) returns boolean as
$$
begin
    if (n < lowerBound) then
        raise exception 'lower bounds check error';
    end if;
    return true;
end;
$$ language plpgsql;
