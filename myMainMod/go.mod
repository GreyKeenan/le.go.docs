module example.com/myMainMod

go 1.22.2

replace example.com/myLibMod => ../myLibMod

require example.com/myLibMod v0.0.0-00010101000000-000000000000
