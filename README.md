# ContestBot

This gaming app can be used for various types of contests: Quizz, Songs

It leverages the Cisco Spark client where the contest is settled, and the Tropo API to play the contest in a Spark Room.


## setup the bot

- launch Cisco Spark, create a Contest Room, invite Contest participants

- install the [ContestBot](spark/Readme.md) 


## starting up a contest

/launch &lt;contest name&gt; : launches a new contest
- Bot says "new contest XXXX starting"
- a contest is picked from the library (udio challenge, answer, creator)
- the room is called by Tropo, participants join, and the audio is played

Quizz example : I travel on the network. My best friend is a scripting language. CISCO APIs leverage me a lot. I am, I am ... 


## participants

/guess &lt;contest attempt&gt; : participants take their chance
- ContestBot says whether participant found the contest answer or not
-- nice try from <participant name> 
-- try again <participant name>
-- BRAVO <participant name>

after XX minutes or tries, the bot would give the answer to all participants 


## submit a contest

/contribute &lt;phone number&gt; &lt;contest answer&gt;
- you get called back to record your contest quizz, song, phrase or whatever will give clues to participants about the contest answer


# Future thoughts

- add answer accuracy match rather than strict

- add a retry option : the audio is replayed, and participants are given one more chance to win the contest

- contest rating by the participants (likes)

- contributors get notified via SMS when their contest is answered or liked by participants


# How to run the ContestBot by yourself

## Configure Tropo
 
see Readme.md in tropo/
 
 
## Start the Webhook

- set both BOT_SPARK_TOKEN and BOT_TROPO_TOKEN env variables with the tokens provided by Spark and Tropo Developper Portals

``` bash
export BOT_SPARK_TOKEN="MmIzYTk0MWYtZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZTVhZWExOGEtM2R"
export BOT_TROPO_TOKEN="73716c7756656ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ4b76464c77544b5a69764673657a4c6574"
```

- install localtunnel
- git clone repo
- go to spark folder
- make run



