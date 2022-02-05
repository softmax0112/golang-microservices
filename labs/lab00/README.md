# Lab 00 - Getting Set up for Exercises

In this first exercise we'll make sure that we're all set up with our AWS credentials and access to AWS.

## Log into the AWS Console

1. Log in to AWS using the link, username, and password provided to you
1. In the top bar, near the right, you'll see your username/alias @ tf-intro - clicking on that will display a dropdown
1. In the dropdown, click on "My Security Credentials"
1. This will take you to your security credentials screen/tab; feel free to change your password if you like, you'll be using this account for the next 2 days.
1. Use the provided access key/secret key (or click "Create access key" to create a new one)
1. If you've clicked "Create access key", an access key and secret will be created for you, **copy the Access key ID and Secret access key (or download a CSV file), we'll use them in setting your environment up below**
1. If you've clicked "Create access key", close out of the modal/pop-up

## Launch your Environment (Option 1 - Cloud9 in AWS)

1. In the top bar of the AWS Console, in the center, you'll see a search box; click on it, and type "Cloud9" which will filter available services in the search list. Click on "Cloud9" which will take you to where we can create your environment.
1. Click on "Create Environment"
1. Give your environment a unique name (your student alias is suggested) and, optionally, a description. Click "Next step"
1. Keep the settings as their defaults on this screen, then click "Next step"
1. Review your settings on the next screen, and then click "Create environment"
1. Wait for your environment to start. In this step, AWS is provisioning an EC2 instance on which your IDE environment will run. This gives us the distinct advantage of having a consistent and controlled environment for development regardless of client hardware and OS. It also allows us to connect to our instances and AWS's API without worrying about port availability in a corporate office. :-)
1. Once your IDE loads, you should see a Welcome document. Your instructor will give you a walkthrough of the visible panel. Feel free to take a moment to read through the welcome document.

## Set up your environment credentials to connect to AWS

Place the following in your `~/.bash_profile` or `~/.zshrc` (if on a Mac) file at the bottom, and replace the values in brackets with your generated creds:
```
export AWS_ACCESS_KEY_ID=[The access key ID provided or that you created above]
export AWS_SECRET_ACCESS_KEY=[The secret access key provided or that you created above]
export AWS_DEFAULT_REGION=us-east-2
```

Then source your new `.bash_profile` or `.zshrc` (if on a Mac) and ensure environment has the appropriate env vars set:
```
source ~/.bash_profile
env | grep AWS
```
or
```
source ~/.zshrc
env | grep AWS
```

The printenv above should output something like:
```
AWS_SECRET_ACCESS_KEY=xxxxxxx
AWS_DEFAULT_REGION=us-east-2
AWS_ACCESS_KEY_ID=xxxxxx
```

Having done that, we should be ready to move on!
