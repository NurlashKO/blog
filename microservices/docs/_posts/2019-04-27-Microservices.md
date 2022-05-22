---
title: Microservices
author: Tao He
date: 2019-04-22
category: Jekyll
layout: post
---

Writing documentation on such an early stage is complete waste of time.

Take `blog` microservice below as an example. `Endpoint`, `Source code`, ~~protobufs~~, could become outdated tomorrow.
Description doesn't make any sense now, and just slows development.

At this stage your source code and common practices should be the best documentation.
Later you or any other service owner might add related docs, runbooks, etc.,
but it should be done in microservice's own "context"(repo/subdir), not a common "docs" bucket.

Here I am going to just introduce you to common practices that exist in this project.

- All requests from `public` to `internal` network pass through `ingress` microservice. Configuration can be found [here](https://github.com/NurlashKO/blog/blob/main/microservices/ingress/nginx.conf).
  
- Your CI workflow should located in https://github.com/NurlashKO/blog/tree/main/.github/workflows/{name}-ci.yaml
  - e.g. For blog microservices it is [here](https://github.com/NurlashKO/blog/blob/main/.github/workflows/blog-ci.yml).
  
- All provisioning configs located in `https://github.com/NurlashKO/blog/blob/main/microservices/provisioner/services/{microservice_name}.sh`
  - e.g. For blog microservice it will be [here](https://github.com/NurlashKO/blog/blob/main/microservices/provisioner/services/blog.sh)
  - This script is going to be sourced by `provisioner`'s [run.sh](https://github.com/NurlashKO/blog/blob/main/microservices/provisioner/run.sh#L23) here.
  
You might notice that everything on this page can go outdated as well and that's OK.
It is just 1 page update, not 10 individual microservices.

To step aside from My Personal Blog.
Microservices are for big corporations.
Changes above will probably be reviewed by different departments at your company.
You might need permission from security team to register DNS entries,
ask help from CI/CD folks to help you with building pipeline, discuss with SRE about how to setup monitoring/alerting and etc.

Everything here is IMHO.