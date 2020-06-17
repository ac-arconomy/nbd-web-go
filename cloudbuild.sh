cloud-build-local --dryrun=false --push --substitutions SHORT_SHA=test,\
_SHORT_SERVICE_NAME=template,\
_GCP_CONNECTION_INSTANCE=gembet:asia-northeast1:gembet,\
_DATABASE_NAME=template,\
_MIGRATE_DATABASE="true",\
_REGION=asia-northeast1,\
_DART_REPO_NAME=golang-template-service-dart-sdk .
