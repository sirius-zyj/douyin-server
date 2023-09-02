douyin-server
===

# Structure

![structure](https://img1.imgtp.com/2023/09/02/o2w0ui2z.png)


# How to take a quick start?


**We strongly recommend deploying using Docker.**
***

1. Download
```
git clone https://github.com/sirius-zyj/douyin-server.git
```

2. Generate binary file
+ Take your own aliyun OSS and set **OSSAK\OSSSK** in config/config.yml and docker/microservices/etcd/config.yml
![1693657047972.png](https://img1.imgtp.com/2023/09/02/tWo2r4r9.png)
```
sh run.sh

cd rpc
sh build_all_service.sh

cd docker
sh copy.sh
```

3. Deploy

+ Make sure that ports **3306, 2379, 6379, 8080, and 8880-8886 are not occupied** on your local machine.
 <br />3306:mysql
 <br />2379:redis
 <br />6370:etcd
 <br />8080:router
 <br />8880-8886:microservices

 + Make suer that you have set up **docker** and **docker-compose**

 ```
 cd docker
 sudo docker-compose up -d
 ```
