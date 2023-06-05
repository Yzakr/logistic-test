### Welcome!
I made a first approach of what you are looking for the matching test for drivers and addresses.
Considerations:
- I've used Golang for dev
- Just used internal libraries.
- I included a Makefile to group all the main commands.
- I made a some kind of CLEAN arch approach.
- Ready to be used in API env just by adding the router.
- To run the project you will be asked for the filenames but can run by default using repository's files

For running:
````
  make run
````
For building:
```
  make build
```
For Testing:
```
  make test
```

Models:
You will find the structs definition for drivers, address and shipment options.

Repositories:
For data retrieving from files but can be changed for another type.

Usecases:
For business logic and testing some functions there.
