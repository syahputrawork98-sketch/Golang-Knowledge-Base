module example.com/app

go 1.22

require (
	example.com/lib-a v1.2.0
	example.com/lib-b v1.0.0
)

replace (
	example.com/lib-a => ../lib-a-v12
	example.com/lib-b => ../lib-b
)
