# Singleton

A compoenent which is instantiated only once.

## Tips
* For some components it only make sense to have one in the system
  - Database repository
  - Object factory

* the construction call is expensive
  - We only do it once
  - We give everyone the same instance

* Want to prevent anyone creating additional copies
* Need to take care of lazy instantiation

## summary

* Lazy one-time initialization using sync.Once
* Adhere to DIP: depend on interfaces, not concrete types
* Singleton is not scary
