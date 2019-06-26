# Design

In essence try to "binary search" this database to limit the number of wasted (read: empty) responses. Kind of like trying to binpack our responses to get all the data in the fewest number of calls.

Map Go inputs onto the USA Swimming unofficial API

Break the world into atoms, and make logic for the atoms to be bundled up into aggregates. If a request fails due to trying to return too many responses, "pop" it into smaller chunks until it returns data.

Store state in a database

Implement signal listener for cleanup
