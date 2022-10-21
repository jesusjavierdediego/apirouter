#!/bin/bash

if [[ -f "gitea" ]]
then
    echo "The test IT Git repo DOES NOT EXIST in the filesystem. The Integration Test is not possible"
else
    echo "API ROUTER integration tests start"
    docker-compose up -d
    echo "Docker containers for ITs ready."
    sleep 10
    sh prepareTests.sh
    echo "Preparation to tests OK"
    echo "Integration tests start"
    PROFILE=dev go test xqledger/apirouter/apilogger -v 2>&1 | go-junit-report > ../testreports/apilogger.xml
    PROFILE=dev go test xqledger/apirouter/configuration -v 2>&1 | go-junit-report > ../testreports/configuration.xml
    PROFILE=dev go test xqledger/apirouter/utils -v 2>&1 | go-junit-report > ../testreports/utils.xml
    PROFILE=dev go test xqledger/apirouter/grpcclient -v 2>&1 | go-junit-report > ../testreports/grpcclient.xml
    PROFILE=dev go test xqledger/apirouter/kafka -v 2>&1 | go-junit-report > ../testreports/kafka.xml
    PROFILE=dev go test xqledger/apirouter/rest -v 2>&1 | go-junit-report > ../testreports/rest.xml
    PROFILE=dev go test xqledger/apirouter/gitapiclient -v 2>&1 | go-junit-report > ../testreports/gitapiclient.xml
    PROFILE=dev go test xqledger/apirouter/admindb -v 2>&1 | go-junit-report > ../testreports/admindb.xml
    echo "Integration tests complete"
    echo "Cleaning up..."
    cd ../integration-tests
    docker-compose down
    echo "Clean up complete. Bye!"
fi