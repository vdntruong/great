## Gateway

### Run at docker

```bash
docker run -d -p 8080:8080 -p 80:80 -v $PWD/api-gateway/traefik.yml:/etc/traefik/traefik.yml traefik:v3.1
```

Traefik searches for [static configuration file](https://doc.traefik.io/traefik/getting-started/configuration-overview/#configuration-file).

Generate password:
```bash
htpasswd -nb <username> <password>
```
