DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'logPriority') THEN
            CREATE TYPE logPriority AS ENUM ('critical', 'warn', 'info', 'debug');
        END IF;
    END
$$;
