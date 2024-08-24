module example.com/main

go 1.23.0

replace example.com/greetings => ../greetings

replace example.com/hello_world => ../hello_world

require example.com/greetings v0.0.0-00010101000000-000000000000
