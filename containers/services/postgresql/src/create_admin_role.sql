do
$$
    begin
        -- Create an admin account if it doesn't exist
        -- (password should be reset later.  This should run only once.)
        if not exists (select rolname from pg_roles where rolname = 'admin') then
            create role admin with superuser login password 'admin';
            grant all privileges on database postgres to admin;
        end if;
    end
$$;
