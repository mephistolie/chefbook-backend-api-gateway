dir=$(basename "$(pwd)")
if [ "$dir" == "scripts" ]
then
cd ..
fi
swag init -q --parseDependency -g internal/transport/http/handler/handler.go
