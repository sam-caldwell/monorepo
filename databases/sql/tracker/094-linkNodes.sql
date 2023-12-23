/*
 * 094-linkNodes.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Given two nodes (LHS and RHS) representing a prev node (left hand) and next node (right hand)
 * this function will point lhs.nextId->rhs and rhs.prevId->lhs.
 *
 * This is needed when deleting workflow steps.
 */
create or replace function linkNodes(lhs uuid, rhs uuid) returns boolean as
$$
begin

    update workflowsteps set nextstepid=rhs where id=lhs;
    update workflowsteps set prevstepid=lhs where id=rhs;
    return true;
end;
$$ language plpgsql;
