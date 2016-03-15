# ContestBot

Spark + Tropo gaming app

Can be used for Song or Quizz contests


## prep work

Add the ContestBot to a Spark room

POST ... // adds the bot webhook


## starting up a contest

Invite participants to your room

/launch : launches a new contest
- Bot says "new contest starting in XX minutes"
- an audio is picked from the library
- the room is called and the audio is played

## participants

/guess <audio title> : participants make their attempt
- Bot says whether participant found anwer or not
-- nice try <participant name>
-- try again <participant name>

after XX minutes, the bot gives an answer

## add a contest

/contribute <phone number> <contest answer>
- you get called back to record your contest quizz, song, phrase or whatever will give clues to participants about the contest answer


# Future thoughts

- add contest answer match %age rather than strict

- add a retry option : the audio is replayed, and participants are given one more chance to win the contest

- contest rating

- contributors get notified via SMS when their contest is used / found


