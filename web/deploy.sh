rm -rf dist
yarn
yarn build
ssh winkhub "rm -rf /opt/wink-local/dist"
scp -r dist winkhub:/opt/wink-local/
