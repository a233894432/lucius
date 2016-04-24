# Lucius
CRUD for mgo - golang mongodb driver
Package for quick plug'n play use of mgo - mongodb driver for golang language.

Provides functions for basic manipulation of mongodb documents in CRUD (Create, Read, Update, Delete) pattern.

## Install
```sh
go get github.com/vedrans/lucius.git
```

## Usage
Before using db host and name needs to be setup with public `SetDB` function which accepts two string parameters (db host and name) or same can be done using two separate public functions `SetDBHost` and `SetDBName`.

#### Create
`Create` function accepts document in `bson.M` format (from `labix.org/v2/mgo/bson` package) and collection name as string parameter. It returns error  if document couldn't be saved for some reason or nil otherwise.

#### FindByID
`FindByID` function finds one document that is matching given ID and collection name. Function accepts two parameters: ID (in `bson.ObjectId` type) and collection name (in string type) and returns two parameters document in `bson.M` format and error if query failed, nil otherwise.

#### FindBy
`FindBy` function finds one document that is matching given condition and collection name. This function accepts condition writen in `bson.M` format as first parameter and collection name as second parameter. It returns two parameters: document in `bson.M` format and error if query failed, nil otherwise. 

#### FindAll
`FindAll` function returns all documents in specific collection as an array of `bson.M` objects as first parameter on return. It requests collection name as `string` and returns error as second parameter or nil if query was successful. 

#### Update
`Update` function replaces matched document with provided. It accepts condition as `bson.M`, new document  as `bson.M` and collection name as `string`. It returns error if query failed, nil otherwise.

#### Patch
As difference from `Update` function, `Patch` just updates matched document with provided. It accepts condition as `bson.M`, update as `bson.M` and collection name as `string`. It returns error if query failed, nil otherwise. It basically is merging existing document with provided update. If some parameter exists in both documents update one is replacing existing one, if existing document doesn't have parameter update document have  it will be simply added.

#### Delete
`Delete` function deletes matched document from collection. It accepts id of document in `bson.ObjectId` format and collection name as `string` and returns error if query fails otherwise nil.

## Licence
This package is covered by GNU GENERAL PUBLIC LICENSE Version 3 licence.

## Contributing
If you would like to improve this package all pull requests are very welcome as long as covered by test.