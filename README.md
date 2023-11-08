# Demo with Elasticsearch

```
GET /
GET _cat
GET _cat/indices

POST logs/_doc
{
  "service_name": "service A",
  "level": "info",
  "message": "Hello World",
  "@timestamp": "2023-11-08T11:12:43.061+0800"
}

GET logs/_search

GET logs/_search
{
  "query": {
    "match": {
      "message": "hello"
    }
  }
}


GET logs/_mapping
```