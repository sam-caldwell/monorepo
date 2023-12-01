/*
 * 0006-table-ticketAudience.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * map tickets to teams/users and allow overriding permissions.
 */
create table if not exists ticketAudience
(
    id       uuid primary key     default gen_random_uuid(),
    ticketId uuid        not null,
    userId   uuid        null,
    teamId   uuid        null,
    team     permissions not null default 'none',
    everyone permissions not null default 'none',
    foreign key (ticketId) references ticket (id),
    foreign key (userId) references users (id),
    foreign key (teamId) references teams (id)
);
