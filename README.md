# proglog
Simple log

- Add log 
```
curl --request POST \
  --url http://127.0.0.1:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
	"record": {
		"value": "TGV0J3MgR28gIzEK"
	}
}'
```

- Read log by offset
```
curl --request GET \
  --url http://127.0.0.1:8080/ \
  --header 'Content-Type: application/json' \
  --data '{"offset": 3}'
```
