create table if not exists workflow_steps
(
    id          uuid primary key default gen_random_uuid(),
    workflowId  uuid        not null,
    name        varchar(64) not null,
    nextStep    uuid        not null,
    description text
)
