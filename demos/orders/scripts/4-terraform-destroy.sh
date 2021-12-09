cd ../terraform
echo 'Please type version :'
read version
terraform destroy -var="app_version=$version" -input=false -auto-approve