// Trace latest script version deployed
log("20160316-1546, newcontest.js");

// Start outbound call
call("sip:" + room_sip + ";transport=tcp", { "timeout":60, "callerID":botname } );

// Typical delay for the Spark clients to be notified of the incoming call
wait(4000);

// Welcome message to all participants
for (i=0;i < replays; i++) {
    say("Dear participants, welcome to the Contest. Please leave your keyboards for a minute, listen and be the first to answer the contest.");
}

// Wait for the signal to launch the contest
// Note: this could be a secret known to the contest launcher only
ask("Are you ready to go ? say yes, wait or abort", {
    choices:"yes, abort",
    repeat:3,
    terminator:"#",
    onBadChoice: function(event) {
        say("I'm sorry, I didn't understand what you said.");
    },
    onTimeout: function(event) {
        say("I'm sorry. I didn't hear anything.");
    },
    onChoice: function(event) {
        switch (event.value + "") {
            case 'abort':
                say("Too bad. Hope we'll see you again. Bye");
                hangup();
                break;
            case 'yes':
                break;
        }
    }
});

// Starting challenge
say("Challenge is starting");
wait (1000);
say("Place your guess in the Spark Room with the syntax / guess followed by your answer");
wait (2000);
say("Let's go");
wait (2000);

// Play audio file
// say("http://soundbible.com/mp3/I%20Love%20You%20Daddy-SoundBible.com-862095235.mp3")

// or speak a challenge
say("I travel on networks.");
wait(2000);

say("CISCO APIs leverage me a lot.");
wait(2000);

say("My best friend is a scripting language.");
wait(2000);

say("I am");

// Leave a few seconds to the participants
wait(5000);

// End of challenge
say("challenge is over");

// Close call
hangup();