# zerorpc_demo
一个grpc的小demo
```
curl -X GET  http://192.168.1.52/v1/domain/all
curl -X POST -H "Content-Type: application/json" -d "{\"username\": \"david\",\"password\": \"123456\"}" http://192.168.1.52/v1/user/login
```
# 服务发现
基于 Kubernetes Endpoints 的服务发现
https://mp.weixin.qq.com/s/-WaWJaM_ePEQOf7ExNJe7w
# istio
```
./deploy-istio #deploy配置 包含istio gateway dr vs配置
```
![image](https://github.com/david966524/zerorpc_demo/assets/121029437/3157c782-fafc-40a3-9ec4-4b9c3583332e)

# jaeger
```
链路追踪配置
etc中
```

![image](https://github.com/david966524/zerorpc_demo/assets/121029437/51677fa6-1510-42f8-9320-96aef81eb411)
