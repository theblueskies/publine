.PHONY:
	build
	dev

runner:
	cd lambdas/runner && make build
	cd lambdas/cleaner && make build
