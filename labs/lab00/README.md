# Lab 00 - Getting Set up for Exercises

In this first exercise we'll make sure that we're all set up with our access to AWS (for the Lambda exercise) and initial GoLang project configuration.

## Log into the AWS Console

1. Log in to AWS using the Account ID, username, and password provided to you
1. In the top bar, near the right, you'll see your username/alias @ `<Account ID>` - clicking on that will display a dropdown
1. In the dropdown, click on "My Security Credentials"
1. This will take you to your security credentials screen/tab; feel free to change your password if you like, you'll be using this account for the next 3 days.

## Option 1 - Launch and Use a Cloud9 Environment

1. In the top bar of the AWS Console, in the center, you'll see a search box; click on it, and type "Cloud9" which will filter available services in the search list. Click on "Cloud9" which will take you to where we can create your environment.
1. Click on "Create Environment"
1. Give your environment a unique name (your student alias is suggested) and, optionally, a description. Click "Next step".
1. Select the `SSH Access` option
1. Keep the other settings at their defaults on this screen, then click "Next step"
1. Review your settings on the next screen, and then click "Create environment"
1. Wait for your environment to start. In this step, AWS is provisioning an EC2 instance on which your IDE environment will run. This gives us the distinct advantage of having a consistent and controlled environment for development regardless of client hardware and OS. It also allows us to connect to our instances and AWS's API without worrying about port availability in a corporate office. :-)
1. Once your IDE loads, you should see a Welcome document. Your instructor will give you a walkthrough of the visible panel. Feel free to take a moment to read through the welcome document.
1. Cloud9 automatically comes with `git`, `terraform`, `Docker`, and `GoLang` installed. This environment gives you access to deploy components to AWS and can also be used as a development environment.

## Option 2 - Use Local Installation of Terraform, Visual Studio Code or GoLand

1. Create a new folder for housing your lab work during the class. Recommendation is to use a separate folder for each service.
1. In the root of the folder, execute `go mod init <unique-path>` (e.g., `go mod init github.com/KernelGamut32/golang-microservices/<service-name>`)
1. You can build out your project for the service in that folder using the structure of your choice (e.g., https://github.com/golang-standards/project-layout)
1. Use `go mod tidy` in the root of the folder to refresh any required dependencies as you build out the service

When complete, you will have multiple options for executing the build out of the exercises - 1) you can use GoLang and an IDE (like VS Code or GoLand) locally; 2) you can use the Cloud9 environment for GoLang development and any AWS deployments you need to run.

With all of that done, we should be ready to move on!
