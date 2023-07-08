
IGNORE_FILE=src/scripts/exclude-list.txt
./vars.sh

cd ../..
rsync -avz -e "ssh -i $KEY_PATH" --exclude-from $IGNORE_FILE ./ $USER@$IP_ADDR:/home/$USER/mse
# ssh -i $KEY_PATH $USER@$IP_ADDR "cd /home/$USER/mse/src/scripts; ./start-prod.sh"