# WebHook registration

##  Pick your room ID

- Go to https://developer.ciscospark.com/endpoint-rooms-get.html

- Run the request in test mode

GET https://api.ciscospark.com/v1/rooms
Content-type	application/json 
Authorization	Bearer <your spark token here>

200 OK



##  Add the webhook

- Go to https://developer.ciscospark.com/endpoint-webhooks-post.html

- Fill in the form and run the request in test mode

POST https://api.ciscospark.com/v1/webhooks
Content-type	application/json 
Authorization	Bearer <your spark token here>
{
	"name": "ContestChallenge",
	"targetUrl": "https://contestbot.localtunnel.me",
	"resource": "messages",
	"event": "created",
	"filter": "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vZDNmN2YwYTAtZThiZS0xMWU1LTllZjYtZjFmMDhjMjNjNDI5"
}

200 OK
{
	"id": "Y2lzY29zcGFyazovL3VzL1dFQkhPT0svZTg5ZmFmZGYtNzNjOC00NDRkLWJkZGMtMGFlMGQ1OWNhMmMy",
	"name": "ContestChallenge",
	"resource": "messages",
	"event": "created",
	"targetUrl": "https://contestbot.localtunnel.me",
	"filter": "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vZDNmN2YwYTAtZThiZS0xMWU1LTllZjYtZjFmMDhjMjNjNDI5"
}

