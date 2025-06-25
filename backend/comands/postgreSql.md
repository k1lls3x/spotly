sudo pacman -Syu postgresql postgresql-pgcron

sudo -iu postgres initdb --locale=en_US.UTF-8 -D /var/lib/postgres/data

sudo systemctl start postgresql.service
sudo systemctl enable postgresql.service

# Отредактируй postgresql.conf, добавь:
# shared_preload_libraries = 'pg_cron'

sudo systemctl restart postgresql.service

# Далее в psql
psql -U postgres
postgres=# CREATE EXTENSION pg_cron;
