# How to send notifications to the notification service

The script below outlines how to send notifications to the notification service. 
```
curl --request POST \
  --url http://localhost:3000/notifications \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data '{
  "type": "Warning",
  "title": "Backup failed",
  "description": "Backup failed due to a network error"
}'
```

> [!NOTE]
> The type property supports the following content: Warning, Info, Debug, Fatal, Error
> Furthermore title and descriptions are required

There is also an [OpenApi specification](/api/notifications-openapi-v3.json) that you can import into your tool (such as Postman)

