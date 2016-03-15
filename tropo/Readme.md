# Installation

- Create a Tropo application at http://tropo.com
- Add the newcontest.js script
- Generate the Tropo Token and add it to a BOT_TROPO_TOKEN env variable
   
   
## To launch a new contest
 
The Tropo JS script takes a 3 arguments : token, number, audio

The script can be invoked manually :

``` bash
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" 
-d 'token=<token>&number=sip:<roomSIP@ciscospark.com>&audio=http://soundbible.com/mp3/I%20Love%20You%20Daddy-SoundBible.com-862095235.mp3' 
"https://api.tropo.com/1.0/sessions?action=create"
```
