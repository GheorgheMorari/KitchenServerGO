docker build ./ -t kitchen_image
docker stop kithcen_container
docker run -d --rm -p 8000:8000 --name kitchen_container kitchen_image go run main