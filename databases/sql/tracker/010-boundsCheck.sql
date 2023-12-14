/*
 * 010-boundsCheck.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * boundsCheck()
 *  given value (n), return true if n is within the bounds of lowerBound and upperBound
 *  otherwise raise an exception
 */
create or replace function boundsCheck(n integer, lowerBound integer, upperBound integer) returns boolean as
$$
begin
    if ((n < lowerBound) or (n > upperBound)) then
        raise exception 'bounds check error';
    end if;
    return true;
end;
$$ language plpgsql;
