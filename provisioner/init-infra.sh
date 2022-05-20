#!/bin/bash

cat > /etc/systemd/system/blog.service <<- EOM
[Unit]
Description=Provision infrastructure
[Service]
ExecStart=/bin/sh /home/nurlashko/provision/run.sh
ExecPost=/bin/rm -rf /home/nurlashko/provision/run.sh
EOM

cat > /etc/systemd/system/blog.timer <<- EOM
[Unit]
Description=Check and provision if needed every 10 seconds
[Timer]
OnBootSec=10
OnUnitActiveSec=10
AccuracySec=1ms
[Install]
WantedBy=timers.target
EOM

docker run -d -v /home/nurlashko/provision:/opts/provision gcr.io/kouzoh-p-nurlashko/nurlashko/provisioner
