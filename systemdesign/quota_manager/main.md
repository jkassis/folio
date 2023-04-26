```
                        _______________
                   _____/               \_____
                 _/                           \___
                /         QUOTA SERVICE           \
                |                        _________/
                 \__            ________/
                    \__________/

             +------------+  +------------+  +-----
             | S1 Storage |  | S2 Storage |  | ....
             -------------+  +------------+  +-----
Application  +----+  +----+  +----+  +----+  +----+
Services     | S1 |  | S1 |  | S2 |  | S2 |  | ....
             +----+  +----+  +----+  +----+  +----+
                  \    |     /           ___--
                   \   |    /       ___--
                    \  |   /    __--
                     \ |  /__---
                     Client
```


GENERAL REQUIREMENTS
------------
## Functional
+ get up to date metrics on utilization per customer... in aggregate, but also by provider.
+ get / set quotas per customer in aggregate
+ rate-limit / apply back pressure to write requests that increase the utilization
+ always allow deletes / removal
+ allow customer to see utilization by storage provider / backend

## Security
+ don't let other customers see each other's quota / utilization
+ allow admins to see everything (from cloud service provider's perspective)

## Non-Functional
+
+

## Utilization Data Pipeline
* Function
  + Publish, store, and serve utilization metrics
* Data: utilization by customer
  + Logging: no
  + Metrics: yes... these are gauges, not counters, etc.
* Options
  + Option: Data Push
    could provide fresher data, but you manage sockets
  + Option: Data Pull/Scrape
    resilient at scale. don't worry about sockets. less fresh. ok since data is changing a lot.
  + Option: Shared Infra for Publish / Scrape / Subscribe / Warehouse / Serve
    + Potentially not optimal from performance perspective.
    + Competing with noisy neighbors.
  * Option: Dedicated Infra for Publish / Scrape / Subscribe / Warehouse / Serve



MICROSERVICES / COMPONENTS
--------------------------
## Utilization Data Publisher
* Function
  + Publish actual utilization metrics
* Data: telemetry: writes / deletes
* Challenges
  + Hybrid / Multi-Cloud
* Deployment
  + To Data Storage Providers
    easier to integrate at deeper level. fewer parties to coordinate

## Utilization Data Scraper
* Function
  + Scrape utilization metrics and store in Quota/Utilization Server

## Quota/Utilization Data Server
* Function
  + Manage quotas per customer
  + Store current utilization
  + Serve utilization and quota
* Scale
  + number of storage provider systems... < 1k
  + customers... << BB << MM (service-accounts).
* Deployment
  + K8S StatefulSet with raft'd dedicated backend
  + K8S Single instance replicaset with scalable cloud data source

## RateLimit Client
* Function
```
  func (rl *RateLimiter) writeOK(size int64, obj Datastore) error
```
* Deployment: API
    + Use gRPC / Swagger / Thrist / etc.
      Distribution requires delivering polyglot solution. Maintenancy challenge. Maybe mitigates with gRPC / Swagger. Standard IDL.

## Rate Limit Server
* Scale
  + QPS == Writes Per Second --> Assume Infinite...
* Security
  + users can request permission to write to their own quota
* Deployment
  + K8S ReplicaSet/Deployment allows horizontal scaling
  + Stateless for the horizontally scaled component
+ Has a cache or copy of quota values per customer
  + Keep entire dataset of limits on the rate limit
  + if too much, we can look at caching
    LRU Caching strategy & freshness timeout / SLO for quota changes
+ challenges
  + limiting small number of large write requests
  + limiting a large number of small write requests
+ solution for concurrency
  + limit writeOK maximum request size (eg 10mb at a time) and require multiple authorizations over time during large writes
  + shard rate-limiter by user and add smarts to RL to keep a "fuzzy" quota
    + assume that all requested writes go through
    + fuzzy usage = known actual usage + requested usage
    + if fuzzy usage < quota => reject
    + periodicaly fetch actual usage (SLO territory)
