#!/bin/bash

sudo bash -c 'cat > /etc/systemd/system/blog.service <<- EOM
[Unit]
Description=Provision infrastructure
[Service]
User=nurlashko
Group=nurlashko
ExecStart=/bin/bash -c "bash /home/nurlashko/provision/run.sh && sudo rm -rf /home/nurlashko/provision/run.sh "
EOM'

sudo bash -c 'cat > /etc/systemd/system/blog.timer <<- EOM
[Unit]
Description=Check and provision if necessary [Every 10 seconds]
[Timer]
OnCalendar=*:*:0/5
Unit=blog.service
[Install]
WantedBy=default.target
EOM'

sudo systemctl enable /etc/systemd/system/blog.timer
sudo systemctl start blog.timer

docker kill $(docker ps -q)
docker system prune -f

docker pull gcr.io/kouzoh-p-nurlashko/nurlashko/provisioner
docker run --name provisioner -d -v /home/nurlashko/provision:/opts/provision/tmp gcr.io/kouzoh-p-nurlashko/nurlashko/provisioner
