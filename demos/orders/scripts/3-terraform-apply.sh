cd ../terraform
echo 'Please type version :'
read version
terraform init
terraform apply -var="app_version=$version" -input=false -auto-approve