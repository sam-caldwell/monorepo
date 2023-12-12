/*
 * 0100-func-boundsCheck.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function boundsCheck(n integer, lowerBound integer, upperBound integer) returns boolean as
$$
begin
    if ((n < lowerBound) or (n > upperBound)) then
        raise exception 'bounds check error(boundary: % to %s)(value:%)',lowerBound, upperBound,n;
    end if;
    return true;
end;
$$ language plpgsql;
