bucketname=`cat bucketname`
aws s3 rm s3://$bucketname --recursive
aws s3api delete-bucket --bucket=$bucketname