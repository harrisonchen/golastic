# golastic

API endpoints interfacing with [ElasticSearch](https://www.elastic.co/products/elasticsearch) [Go client](https://github.com/belogik/goes). Run on [Docker](https://www.docker.com/) containers.

### Running the client on Docker

If you do not have Docker and doctor-compose CLI setup, please refer to the [Docker Setup](#docker-setup) section.

Run:

```
docker-compose up
```

Pinging the root url should return status 200
```
curl $(boot2docker ip):8080
```

### Docker Setup

1) Install [Boot2Docker](https://github.com/boot2docker/osx-installer/releases/download/v1.7.1/Boot2Docker-1.7.1.pkg)
  
    After installation, run:
    
    ```
    boot2docker init
    boot2docker up
    boot2docker shellinit
    ```

    Suggestion: Place the boot2docker env variables into your .bash_profile to load every new shell session
    
2) Install docker-compose CLI
 
    ```
    curl -L https://github.com/docker/compose/releases/download/1.3.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
    ```
 
