### Setting up Gober
Build docker image
```
docker build --tag gober .
```

List images
```
docker image ls
```

Run container
```
docker run --publish 8080:8080 gober
```


### Setting up OSRM
1) Pre-process the Singapore data extract
```
docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-extract -p /opt/car.lua /data/Singapore.osm.pbf
```

2) Partition the data extract
```
docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-partition /data/Singapore.osrm
docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-customize /data/Singapore.osrm

```

3) Run the OSRM router server
```
docker run -t -i -p 8081:5000 -v "${PWD}:/data" osrm/osrm-backend osrm-routed --algorithm mld /data/Singapore.osrm
```

4) Make a request to test
```
curl "http://127.0.0.1:8081/route/v1/driving/13.388860,52.517037;13.385983,52.496891?steps=true"

```