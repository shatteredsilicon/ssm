#!/bin/bash

## setup as we do in pmm-framework
mysql -h 127.0.0.1 -u root -pps -e "SET GLOBAL slow_query_log='ON';"
mysql -h 127.0.0.1 -u root -pps -e "INSTALL PLUGIN QUERY_RESPONSE_TIME_AUDIT SONAME 'query_response_time.so';"
mysql -h 127.0.0.1 -u root -pps -e "INSTALL PLUGIN QUERY_RESPONSE_TIME SONAME 'query_response_time.so';"
mysql -h 127.0.0.1 -u root -pps -e "INSTALL PLUGIN QUERY_RESPONSE_TIME_READ SONAME 'query_response_time.so';"
mysql -h 127.0.0.1 -u root -pps -e "INSTALL PLUGIN QUERY_RESPONSE_TIME_WRITE SONAME 'query_response_time.so';"
mysql -h 127.0.0.1 -u root -pps -e "SET GLOBAL query_response_time_stats=ON;"
mysql -h 127.0.0.1 -u root -pps -e "GRANT SELECT, PROCESS, SUPER, REPLICATION CLIENT, RELOAD ON *.* TO 'pmm-agent'@'%';"
