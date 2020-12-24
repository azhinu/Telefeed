#!/bin/sh


read -p "Enter Telefeed install path (Default /usr/local/bin/telefeed/ ): " REPLY
install={$REPLY:-/usr/local/bin/telefeed}

#Install software

echo "Installing server….."
sudo apt update  # To get the latest package lists
sudo apt install redis-server -y
mkdir -p $install
cp -rf ./* $install/
read -p "Enter your Telegram Token: " REPLY
echo "$REPLY" > $install/telefeed.bot
read -p "Enter your VK Token:
(If you want to use multiple tokens, you can add if later.)
" REPLY
echo "$REPLY" > $install/tokens.bot

#Install service config
echo "Removing old configuration.."
sudo rm /etc/systemd/system/telefeed.service

echo "Configurating services..."

echo "[Unit]
Description=Telefeed
After=redis.service

[Service]
Type=forking
WorkingDirectory=$install/

ExecStartPre=$install/boltsrv
ExecStart=$install/tgsrv
ExecStartPost=$install/postsrv publics
ExecStartPost=$install/postsrv feeds

KillMode=process
Restart=on-failure

[Install]
WantedBy=multi-user.target" | sudo tee --append /lib/systemd/system/telefeed.service > /dev/null
sudo chmod 644 /lib/systemd/system/telefeed.service


echo "Enabling services..."
sudo systemctl daemon-reload
sudo systemctl enable telefeed.service
sudo systemctl start telefeed.service
