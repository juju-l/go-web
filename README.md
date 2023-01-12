
```
grpcurl -insecure -d '{"code":"debug"}' -proto wxService.proto 127.0.0.1:443 v1.Wx/Auth
methods=(Auth GetAvatarUrl GetBusinessInfo `#UpBusinessInfo` GetDetails GetOrder GetRepair) # .......
for method in ${methods[@]}; do

grpcurl -insecure \
-d '{
	"code": "debug",
	"openid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"entid": "001d6fdb-79b5-41ff-af29-99a0e8618c37",
	"url": "",
	"businessid": "",
	"name": "",
	"num": "",
	"type": "",
	"description": "",
	"detail": "",
	"id": "8d3df1e1-9891-4d23-8435-a62c3e3bf6a1"
}' \
-proto wxService.proto \
-v -H token:\ `echo -n|grpcurl -insecure -d '{"code":"debug"}' -proto wxService.proto 127.0.0.1:443 v1.Wx/Auth|jq .data.token -r` -emit-defaults \
127.0.0.1:443 \
v1.Wx/\
$method

done
```
