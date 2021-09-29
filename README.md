# franz-size

Don't mind this, just for funsies on trying to refresh go and stuff, name is a pun from 'fun size'. Yes, funny.

WIP: Current state of Base Code is a Frankenstein from projects here and there 

Goals:
* Create (possibly) modern go service
* Create a franz-go kafka client to get some arbitrary metadata from arbitrary kafka topic
* Persist (or not to) persist the metadata in arbitrary topic as Datapoint of [tstorage](https://github.com/nakabonne/tstorage)
* Create some endpoints to serve the data persisted, maybe as 
http://baseurl/endpoint/label/EPOCH