create table if not exists ticketTypes (
    id uuid primary key default gen_random_uuid(),
    -- each workflow has a unique name --
    name varchar(64) not null unique,
    -- a uuid representing the icon for the workflow.
    iconId uuid not null,
    -- descriptive text --
    description text,
    foreign key (iconId) references icons(id)
);
