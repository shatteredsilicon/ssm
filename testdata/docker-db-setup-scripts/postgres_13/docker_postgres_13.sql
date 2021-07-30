#!/bin/bash

## Percona-distribution postgresql
create extension pg_stat_statements;
create extension pg_stat_monitor;
SELECT pg_reload_conf();
