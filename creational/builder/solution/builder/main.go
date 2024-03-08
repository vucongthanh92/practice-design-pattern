package main

import "design-pattern/creational/builder/solution/builder/internal"

func main() {
	director := internal.NewDirector()
	builder := internal.NewBuilder()

	service := director.BuilderService(builder)
	service.DoBusiness()
}
