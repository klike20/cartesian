Cartesian API
=============

Create an API server in [go](https://golang.org/). It will deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane. Take a look at https://en.wikipedia.org/wiki/Cartesian_coordinate_system if you need a refresher on this concept.

It must have an api route at `/api/points` that accepts a `GET` request with the following parameters, and returns a JSON list of points that are within `distance` from `x,y`, using the Manhattan distance method. The points should be returned in order of increasing distance from the search origin.
- `x` integer (required). This represents the `x` coordinate of the search origin.
- `y` integer (required). This represents the `y` coordinate of the search origin.
- `distance` integer (required). This represents the Manhattan distance; points within `distance` from `x` and `y` are returned, points outside are filtered out.

The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks). It is defined as the sum of the horizontal and vertical distances between points on a grid. Formally, where `p1 = (x1, y1)` and `p2 = (x2, y2)`, `distance(p1,p2) = |x1-x2| + |y1-y2|`.

On startup, the API server should read a list of points from `data/points.json`.




Explanation:

* This is an application which works according to the indication

* It has 2 endpoint
 * http://localhost:10000/api/points
    This one will need to receive 3 parameters like this: http://localhost:10000/api/points?distance=20&y=0&x=0
      will return all the points which are within Manhattan distance of 20 from (0,0) point
 
 * http://localhost:10000/api/points/path
    This one will need one paramter, like this: http://localhost:10000/api/points/path?path=newdata.json
    This one will change the points datasource from data.json (which is the initial one) to newdata.json (these files need to be in the same directory of the main file). For now there are these 2 files as sources

* The validation is done via the specific url (otherwise a 404 will be returned)

* The intention was to define a design based on single responsabilities and separation of concern, that's why 3 more package where added: entitied (for required structures), loader (for read the file a create the array of points) and calculator (for calculate the distance and identify the points which match the conditions). This additiona package need to be put and installed in the gopath


* Integration and Unit Test where added to test the api comsuption and the main business logic related with the calculation and the file loading

* the main2.txt file includes also a first approach without extra packages focus on the functionality

* Some commented code will be found, only with the purpose to show different alternatives for the implementation

* in the entities package, some commented code can be seen with the intentation to show how the implementation could be extended in case the coordinates system changes so the distance should be calculates in other way based on the coordinates for that paticular system (cilindric, spheric, polar, etc)
