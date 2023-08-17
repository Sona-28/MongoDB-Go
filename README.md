# MongoDBConnection to Go
**********************************
1. Create a common method that takes two parametres
first - databasename
second - collection name

This method will return mongo collection

## Operators
*************
1. **$unwind** - this operator will be used to split the array documents into seperate documents for each element in the array.
    Eg: db.getCollection("students").aggregate([{$unwind:"$course"}])
2. **$pipeline** - when you are performing the aggregation, pipeline will give you the different stages to perform the aggregation
3. **$match** - will be acting as the search query that can be integrated with aggregation pipeline
4. **$lookup** - will be used to perform the join between two tables
