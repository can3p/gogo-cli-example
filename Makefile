shell:
	flyctl postgres connect -a gogo-cli-exampledb

tunnel:
	flyctl proxy 5432 -a gogo-cli-exampledb
