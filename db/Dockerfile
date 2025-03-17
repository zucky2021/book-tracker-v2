FROM mysql:8.4 as build

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD mysqladmin ping -h localhost -u root -p${MYSQL_ROOT_PASSWORD} || exit 1

EXPOSE 3306
