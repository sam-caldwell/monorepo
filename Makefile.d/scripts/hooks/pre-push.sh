#!/bin/bash -e

remote="$1"
url="$2"

zero=$(git hash-object --stdin </dev/null | tr '[0-9a-f]' '0')

while read local_ref local_oid remote_ref remote_oid
do
        if test "$local_oid" = "$zero"
        then
                # Handle delete
                :
        else
                if test "$remote_oid" = "$zero"
                then
                        # New branch, examine all commits
                        range="$local_oid"
                else
                        # Update to existing branch, examine new commits
                        range="$remote_oid..$local_oid"
                fi

                # Check for WIP commit
                commit=$(git rev-list -n 1 --grep '^WIP' "$range")
                if test -n "$commit"
                then
                        echo >&2 "Found WIP commit in $local_ref, not pushing"
                        exit 1
                fi
        fi
done

exit 0
