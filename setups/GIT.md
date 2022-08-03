# Setting up GIT with SSH

## Installing GIT

### Windows

https://git-scm.com/

Install for windows, and choose to also include GitBash

### MacOS / Linux

run "sudo apt install git"

## Setup SSH key

- Open your terminal (gitbash for windows)
- Type "ssh-keygen" and just accept all the default values
- Type "cat ~/.ssh/id_rsa.pub" and copy the content 
- Head to https://github.com/settings/keys
- Click "New SSH key" and print the content of the file above into the Key field

## Done
You can now use the SSH clone command to clone repositories and interact with github
