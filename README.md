# 图片验证码


提供图片验证码生成与验证服务。



Docker service 
```
docker service create --name sms --network cluster --replicas 1  -d hub.unionlive.com/sms:latest
docker service update --force --update-parallelism 1 --update-delay 3s --image hub.unionlive.com/sms:latest sms
docker service update  --replicas 3  sms
```


# Change Log

## v1.0.0 
    [Release Date 2019-12-08]
- [feature] 创建项目
