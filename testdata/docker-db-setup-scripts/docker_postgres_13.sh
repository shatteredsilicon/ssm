#!/bin/bash

## Percona-distribution postgresql
psql -h localhost -U postgres -c 'create extension pg_stat_statements'
psql -h localhost -U postgres -c 'create extension pg_stat_monitor'
psql -h localhost -U postgres -c 'SELECT pg_reload_conf();'
