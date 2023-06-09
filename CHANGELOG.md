## 0.1.5 (2023-05-08)

### Perf

- async handle new block when received

## 0.1.4 (2023-05-08)

### Feat

- **repo**: add send `new_block` message to kafka (#15)
- impl publish new block to kafka
- add kafkax and inject into repo

### Fix

- gen-go

## 0.1.3 (2023-05-07)

### Feat

- impl get block by hash

## 0.1.2 (2023-05-07)

### Feat

- impl get block by hash in repo
- bind get block by hash api

### Fix

- ignore health check logger

## 0.1.1 (2023-05-07)

### Fix

- dial ethclient with websocket

## 0.1.0 (2023-05-07)

### Feat

- add health api
- inject migrate into repo
- inject ethclient into repo
- impl list blocks api
- add restful main file
- impl list blocks in biz
- impl list blocks in repo
- create new block into database
- impl create new block in repo
- inject mariadb and execute migration
- start block listener
- impl listen new block in biz (#8)
- impl listen new block in biz
- add block biz impl
- impl listen new block in repo (#7)
- impl listen new block in repo
- init ethclient when new repo
- add impl file for repo
- add listener block adapter
- impl block listener (#6)
- impl service for main file
- print config options
- add wire inject for main file
- add block listener cmd
- add listener adapter
- define block repo interface
- create blocks table (#5)
- add migration
- add mariadb client
- define block biz interface
- add biz inteface
- add internal pkg (#4)
- add httpx server
- add log
- add config
- add some pkg (#3)
- add adapters
- add response
- add netx
- add httpx
- add contextx and er
- add domain entity (#2)
- add transaction entity
- add block entity
- add pb for domain entity

### Refactor

- restful folder
- repository injection (#9)
- add app servicer
- rename adapter
