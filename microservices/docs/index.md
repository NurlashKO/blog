---
layout: home
---

This documentation provides high level overview of Provisioning and CI/CD processes in
"My Personal Blog", as well as ~~collection of microservice documentations~~.

## Infrastructure overview

![image](https://user-images.githubusercontent.com/10639020/169694404-05f96c8f-328e-4027-b5a4-51f1d70d4c49.png)

While designing, I was trying to achieve next goals:
- Ship stuff to production as fast as possible with minimum efforts required from engineering(me).
- No microservice should require any maintenance. Reverting+Rebooting should fix 99.99% of issues.
- Deleting old and re-provisioning new machine should be done in a couple of clicks.
- Security

From the above you can already guess the purpose
of watcher, existance of init-infra.sh script and overall extensive usage of containerization technologies.
Feel free to explore more by selecting any topic in menu on the left.