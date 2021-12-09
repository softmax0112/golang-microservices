echo "Please type your bucket name"
read bucketname
cat > bucketname <<EOF
$bucketname
EOF
bucketname=`cat bucketname`
aws s3api create-bucket --bucket=$bucketname --region=us-east-1
cd ../terraform
cat > variables.tf <<EOF
variable "app_version" {}
variable "s3_bucket" {
  default = "$bucketname"
}
EOF