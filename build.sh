VERSION=$(git rev-parse HEAD)
SERVICE=$(basename $PWD)
TAG=563473344515.dkr.ecr.eu-central-1.amazonaws.com/$SERVICE:$VERSION

echo "Building image: $TAG"

docker build . -t $TAG
docker push $TAG
