java:
	@GO111MODULE=on go build -o jvmgo cmd/java/main.go

javap:
	@GO111MODULE=on go build -o javap cmd/javap/main.go

.PHONY: java javap