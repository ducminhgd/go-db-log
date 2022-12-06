# Logging using Database

I recommend this only use for meaningful logs.

## Data manipulated event log

This kind of logging is used for records history of creating or changing data.

| Field       | Data type | Required | Nullable | Allow empty | Description                                                                 |
| ----------- | :-------: | :------: | :------: | :---------: | --------------------------------------------------------------------------- |
| **id**      |  String   |   Yes    |    No    |     No      | UUID represent for an event, should be PK                                   |
| event_type  |  String   |   Yes    |    No    |     No      | Type of event, one of [`CREATED`, `UPDATED`, `DELETED`]                     |
| object_type |  String   |   Yes    |    No    |     No      | Type of object, should be name of model, such as: `user`, `contact`, `role` |
| object_id   |  String   |   Yes    |   Yes    |     No      | ID object, in failed creating case, it can be `NULL`                        |
| data        |  String   |    No    |   Yes    |     Yes     | Changed or created data                                                     |
| result      |  Integer  |   Yes    |    No    |     No      | `0`: failed, `1`: success. Default `1`                                      |
| timestamp   | Timestamp |   Yes    |    No    |     No      | Timestamp when an event is created                                          |
