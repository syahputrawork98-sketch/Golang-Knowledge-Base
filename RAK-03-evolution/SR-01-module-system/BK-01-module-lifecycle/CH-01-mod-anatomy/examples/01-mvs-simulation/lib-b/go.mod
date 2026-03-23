module example.com/lib-b

go 1.22

require example.com/lib-a v1.1.0
replace example.com/lib-a => ../lib-a-v11
