# Example VM Configuration

Starting from Ubuntu 22.04 LTS Desktop an example setup procedure is as follows.

Go to settings and disable "Automatic Screen Lock" (the VM does not need to lock, the desktop does this already).

In a terminal run the following

```bash
sudo apt update
sudo apt upgrade -y
sudo apt install git git-lfs build-essential tree vim shellcheck

# Install GO
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
rm go1.23.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile

# Install VSCode and extensions for it
sudo apt-get install wget gpg
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -D -o root -g root -m 644 packages.microsoft.gpg /etc/apt/keyrings/packages.microsoft.gpg
echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" |sudo tee /etc/apt/sources.list.d/vscode.list > /dev/null
rm -f packages.microsoft.gpg

sudo apt install apt-transport-https
sudo apt update
sudo apt install code

code --install-extension=golang.go
code --install-extension=timonwong.shellcheck
code --install-extension=DavidAnson.vscode-markdownlint
code --install-extension=shd101wyy.markdown-preview-enhanced
```

To install the guest additions in an Ubuntu guest run `sudo apt install open-vm-tools` and then reboot the guest VM.
