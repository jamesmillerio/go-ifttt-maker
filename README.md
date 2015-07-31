# go-ifttt-maker
A super simple Go library for interacting with the IFTTT Maker Channel.

# Usage
    //Create an instance of our Maker struct
    maker := new(GoIFTTTMaker.MakerChannel)
    key, event := "your maker key", "your maker event"

    //Set the values to send
    maker.Value1 = "my first value"
    maker.Value2 = "my second value"
    maker.Value3 = "my third value"

    //Send our request
    maker.Send(key, event)
