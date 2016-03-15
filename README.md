# ContestBot

Spark + Tropo gaming app

Can be used for various contests type: Song, Quizz,...


## setup the bot

- start Spark, create a Contest Room, invite participants

- install the ContestBot (see spark/Readme.md) 

## starting up a contest

/launch : launches a new contest
- Bot says "new contest starting in XX minutes"
- an audio is picked from the library
- the room is called and the audio is played

## participants

/guess &lt;contest attempt&gt; : participants make their attempt
- Bot says whether participant found the contest answer or not
-- nice try <participant name>
-- try again <participant name>

after XX minutes, the bot gives an answer

## add a contest

/contribute &lt;phone number&gt; &lt;contest answer&gt;
- you get called back to record your contest quizz, song, phrase or whatever will give clues to participants about the contest answer


# Future thoughts

- add contest answer match %age rather than strict

- add a retry option : the audio is replayed, and participants are given one more chance to win the contest

- contest rating

- contributors get notified via SMS when their contest is used / found


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



