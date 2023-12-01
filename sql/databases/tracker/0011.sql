create table if not exists workflow_actions
(
    id          uuid primary key default gen_random_uuid(),
    stepId      uuid        not null,
    name        varchar(64) not null,
    action      text not null,
    description text
)
