# Samples of starting monitoring of services
# 1. on startup
# PMM_AGENT_PRERUN_SCRIPT="pmm-admin status --wait=10s;pmm-admin add mysql mysql-container mysql:3306 --username=root --password=ps --query-source=perfschema" docker-compose up pmm-server pmm-client mysql
# 2. after startup
# docker-compose up pmm-server pmm-client mysql
# docker-compose exec pmm-client pmm-admin add mysql mysql-container mysql:3306 --username=root --password=ps --query-source=perfschema

