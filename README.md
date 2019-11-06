[한국어](README.kr.md)
[日本語](README.jp.md)
# TourOfGo

## command sequence
### set .env file for host machine volumes permission
- `sh setDotEnv.sh`

### start docker compose
- `docker-compose up -d --build`

### after docker-compose up -d --build
- write terraform template at host machine `./workspace`
- execute terraform template 
  - in docker container
    - `docker exec -it tour-of-go /bin/sh`
  - outside of docker container
    - `docker exec tour-of-go go run section1/1.go`