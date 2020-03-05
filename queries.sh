curl -X POST -H 'Accept: application/json' -d "name=user1" -c "/tmp/token" http://localhost:8080/user/login
# curl -X GET -H 'Accept: application/json' -b "/tmp/token" http://localhost:8080/users
curl -X GET -H 'Accept: application/json' -b "/tmp/token" http://localhost:8080/user/view/1
# # curl -X GET -H 'Accept: application/xml' http://localhost:8080/user/view/1

# # create
curl -X POST -H 'Accept: application/json' -d "name=userT1&mail=user4@test.com" -b "/tmp/token" http://localhost:8080/user/create
# curl -X GET -H 'Accept: application/json' -b "/tmp/token" http://localhost:8080/users
# curl -X GET -H 'Accept: application/json' -b "/tmp/token" http://localhost:8080/user/view/4
