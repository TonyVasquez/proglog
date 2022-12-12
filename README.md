# Distributed log

Proglog is a distributed append only log which comprises a segment, a store file and an index file.

- Record - the data stored into log
- Store - the file where records saved
- Index - comprises offset and position for record in the store
- Segment - combine a store and an index together
- Log - is an abstraction which combine all segments


# API

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
