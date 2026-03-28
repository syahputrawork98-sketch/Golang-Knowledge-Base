module example.com/localreplace/app

go 1.22

require example.com/localreplace/lib v0.0.0

replace example.com/localreplace/lib => ../lib
