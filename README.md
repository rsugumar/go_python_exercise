PRE-REQUISITE:
--------------
    This solution is done in docker, hence docker to be installed to verify the solution.

NOTE:
-----
    - Uses python 3 and go from docker latest images (docker hub)
    - config.json file: This contains the supported currencies symbol and the pre-defined ranges

DOCKER PULL:
------------
    To pull from Docker Hub:
        command: docker pull rsukumar/iamplus

RUN WITH DOCKER:
----------------
Option 1:
    Running with the supported command line arguments
    docker run --rm rsukumar/iamplus:latest python findRange.py -c ./config.json -i "SGD 12.34"

Option 2:
    Running with the volumes option to reflect the local changes into the container (using -v)
    docker run --rm -v ${PWD}/config.json:/usr/src/app/config.json rsukumar/iamplus:latest python findRange.py -c ./config.json -i "SGD 12.34"

GITHUB:
-------
    URL: https://github.com/rsugumar/go_python_exercise

OPTIONAL INFO:
--------------

BUILD:
------
    - To build the docker image:
        command: "docker build --rm -f "Dockerfile" -t rsukumar/iamplus:latest ."
