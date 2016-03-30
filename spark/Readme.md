#  Installation

Add the webhook to the room you want the bot installed


##  Pick a room id

- Go to https://developer.ciscospark.com/endpoint-rooms-get.html, signin, toggle Test Mode, and run the request

![Run request from the Spark Developer Portal](../doc/ListRoomsInPostman.png)

- Identify the Room you want to start the contest from on the right pane and pick its Room id

- Note that the following HTTP request is forge by the Web portal on your behalf

``` bash
GET https://api.ciscospark.com/v1/rooms
Content-type	application/json 
Authorization	Bearer <your spark token here>

200 OK
{
	"items": [
		{
			"id": "Y2lzY29zcGFyazovL3VzL1JPT00vZDNmN2YwYTAtZThiZS0xMWU1LTllZjYtZjFmMDhjMjNjNDI5",
			"title": "Bot Contest Room",
			"created": "2016-03-13T01:56:40.234Z",
			"lastActivity": "2016-03-17T10:04:25.600Z",
			"isLocked": false
		},
		...
		]
}
``` 


##  Register the "Contest Bot" webhook

- Go to https://developer.ciscospark.com/endpoint-webhooks-post.html, signin, toggle Test Mode (if it is not already turned on)

- Fill in the form with your room id and run the request

![Fill the form and run request from the Spark Developer Portal](../doc/AddWebHookViaPostman.png)

The following HTTP request is forge by the Web portal on your behalf

``` bash
POST https://api.ciscospark.com/v1/webhooks
Content-type	application/json 
Authorization	Bearer <your spark token here>
{
	"name": "ContestBot",
	"targetUrl": "http://spark-bot.appspot.com/spark",
	"resource": "messages",
	"event": "created",
	"filter": "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vZDNmN2YwYTAtZThiZS0xMWU1LTllZjYtZjFmMDhjMjNjNDI5"
}

200 OK
{
	"id": "Y2lzY29zcGFyazovL3VzL1dFQkhPT0svZTg5ZmFmZGYtNzNjOC00NDRkLWJkZGMtMGFlMGQ1OWNhMmMy",
	"name": "ContestBot",
	"resource": "messages",
	"event": "created",
	"targetUrl": "http://spark-bot.appspot.com/spark,
	"filter": "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vZDNmN2YwYTAtZThiZS0xMWU1LTllZjYtZjFmMDhjMjNjNDI5"
}
``` 
