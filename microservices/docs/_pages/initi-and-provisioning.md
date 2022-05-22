---
title: Initialization and Provisioning
author: Tao He
date: 2022-02-06
category: Jekyll
layout: post
---

### Initialization

Initialization is just a single script which registers `blog.service` and related `blog.timer`.

> **systemd** is a system and service manager for Linux operating systems. When run as first process on boot (as PID 1), it acts as init system that brings up and maintains userspace services. 

You can consider it as a program that wakes up periodically and does some job.
The only reason I am using it is because smallest period in CRON Tabs is 1 minute.

If you open [init-infra.sh](https://github.com/NurlashKO/blog/blob/main/init-infra.sh), you will notice 
that all `blog.service` does is running `/home/nurlashko/provision/run.sh` and then deleting it once operation succeeds.

### Provisioning

Another important task of `init-infra.sh` is starting latest [provisioner](https://github.com/NurlashKO/blog/tree/main/microservices/provisioner) microservice.

This part is pretty tricky, hope you are still with me up until now.

Provisioner microservice **starts** with `/home/nurlashko/provision/` folder mounted to its `WORKDIR`.
1. `Provisioner` copies his `run.sh` into mounted directory which pushes this file to a **host filesystem**.
2. On the host filesystem, `blog.service` eventually finds this file and executes it.
3. If provisioning finishes with success, `blog.service` deletes this file, which prevents it from re-provisioning it indefinitely.

In other words, `provisioner` microservice here is being used as a package delivery service.

This workflow especially powerful when combined with `Watcher` microservice.

`Watcher` pulls new version and restarts microservice once new version is available.
You guessed it.
Once there are any changes in the provisioning scripts(new microservice, config, etc.), 
changes will be pulled and re-provisioned accordingly.