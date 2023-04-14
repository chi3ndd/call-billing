# call-billing
## Build docker image
```
make build
```
## Run service + database
#### Update config.yaml
```yaml
app:
  address: "0.0.0.0:8910" // webservice address
adapter:
  mongo:
    address: "mongodb://vin:123qwer@mongo:27017" // mongodb url
```
#### Run
```
make deploy
```
## Remove service
```
make remove
```
