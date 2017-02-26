#HTTP="http://127.0.0.1:8080"
HTTP="http://ec2-35-167-129-44.us-west-2.compute.amazonaws.com:8080"

curl -X POST -d '{"Score":1337,"PlayerName":"Sean Plott"}' $HTTP/api/live/topic
