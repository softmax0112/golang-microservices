bucketname=`cat bucketname`
cd ../lambda
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip lambda.zip main
echo 'Please type version :'
read version
aws s3 cp lambda.zip s3://$bucketname/v$version/lambda.zip