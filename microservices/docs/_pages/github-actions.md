---
title: Github Actions CI/CD
author: Tao He
date: 2022-02-06
category: Jekyll
layout: post
---

![image](https://user-images.githubusercontent.com/10639020/169696584-b91a38e6-7ae8-493c-afc8-d84716626e5e.png)

Workflow configurations can be found [here](https://github.com/NurlashKO/blog/tree/main/.github/workflows).

All pipelines follow the same pattern:
- Watch changes in `microservices/${name}` directory
- Build new container image
- Push image to a remove registry

We use long polling from the target server as our CD approach,
which means that once its being pushed, it will be automatically deployed.

GCR service account keys stored as secrets in this repo.

Workflows can be triggered manually on main or any other branch.