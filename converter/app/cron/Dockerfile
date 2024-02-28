from golang:latest

RUN apt -y update
RUN apt -y upgrade
RUN apt-get install -y cron
COPY cronjob /etc/cron.d/cronjob

# Donner les droits appropriés au fichier cron
RUN chmod 0644 /etc/cron.d/cronjob

# Appliquer le fichier de configuration cron
RUN crontab /etc/cron.d/cronjob

COPY app /app

CMD ["cron", "-f"]

