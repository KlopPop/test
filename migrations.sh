#!/bin/bash
chmod +x 000001_create_tables.down.sql
chmod +x 000001_create_tables.up.sql
su postgres
psql -U postgres -d postgres -a -f ./000001_create_tables.up.sql
psql -U postgres -d postgres -a -f ./000001_create_tables.up.sql