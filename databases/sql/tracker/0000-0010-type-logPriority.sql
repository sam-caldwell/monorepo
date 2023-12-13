/*
 * 0000-1000-type-logPriority.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'logPriority') THEN
            CREATE TYPE logPriority AS ENUM ('critical', 'warn', 'info', 'debug');
        END IF;
    END
$$;
