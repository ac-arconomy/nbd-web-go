## Gembet golang service template

#### To regenerate proto and rest from api.proto
1. change to root folder of project
2. sh ./download-scripts  #This downloads required scripts from repo
3. sh ./gembet-scripts/generate-proto-api.sh   #This generate grpc and rest libaries

#### Encrypt file into Google KMSs
gcloud kms encrypt --plaintext-file=rsa_ssh.key \
--ciphertext-file=rsa_ssh.key.enc \
--location=global --keyring=gembet --key=bitbucket-ssh

