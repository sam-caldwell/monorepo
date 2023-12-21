/*
 * 140-workflowActions.sql
 * (c) 2023 Sam Caldwell.  See License.txt.
 */
create table if not exists workflowActions
(
    id          uuid primary key not null,
    name        varchar(64)      not null,
    topic       varchar(2048)    not null, -- the topic to be published to the message processor
    message     varchar(2048)    not null, -- the message to send (formatting string)
    -- --
    created     timestamp        not null default now(),
    description text,
    foreign key (id) references entity (id) on delete restrict,
    constraint validateWorkflowActionName check (validName(name))
);
/*
 * add foreign key to workflowSteps (091-workflowsteps.sql)
 */
alter table workflowsteps
    add constraint fkStepToActions
        foreign key (action)
            references workflowActions(id);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowActionsName on workflowActions (name);
create index if not exists ndxWorkflowActionsTopic on workflowActions (topic);
create index if not exists ndxWorkflowActionsCreated on workflowActions (created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createWorkflowAction()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createWorkflowAction(actionName varchar(64), actionTopic varchar(64),
                                                actionMessage varchar(2048), actionDescription text)
    returns uuid as
$$
declare
    entityId uuid;
begin
    entityId := (select createEntity('workflow_action'::entityType));
    insert into workflowActions (id, name, topic, message, description)
    values (entityId, actionName, actionTopic, actionMessage, actionDescription);
    return entityId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowAction()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowAction(actionId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from workflowActions where id=actionId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
