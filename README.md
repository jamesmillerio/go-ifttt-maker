# go-ifttt-maker
A simple Go library for interacting with the IFTTT Maker Channel.

    //Create an instance of our Maker struct
    var maker = new(GoIFTTTMaker.MakerChannel)
    var key = "your ifttt maker key"
    var event = "your ifttt maker event"

    //Set the values to send
    maker.Value1 = ""
    maker.Value2 = ""
    maker.Value3 = ""
    maker.Value4 = ""

    //Send our request
    maker.Send(key, event)
