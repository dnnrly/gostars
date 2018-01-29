# GoStars!

## Simple plan

* Game resource - defines the game name and universe parameters
* Stars resource - location and name only for now, under a game resource
* Player resource
  * Fields
    * Player name
	* ID/location of corresponding user
	* Place for race attributes
	* Id of game
* Ship types
  * Just details of the construction
* Fleets
  * Fields
    * Fleet composition (sub table for ship type and count)
	* Waypoints ([{location, action at location, index}])
		* GET and POST only
		* Whole array serialized to single field?