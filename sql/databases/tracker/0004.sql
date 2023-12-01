DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'log_priority') THEN
            CREATE TYPE log_priority AS ENUM ('critical', 'warn', 'info', 'debug');
        END IF;
    END
$$;

create table if not exists log
(
    id       serial4       not null primary key,
    priority log_priority  not null default 'info',
    message  varchar(2048) not null
);
