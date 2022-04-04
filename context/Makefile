.PHONY: tests stage

tests:
	@make -C functions/analysis tests
	@make -C functions/store tests
	@make -C functions/stats tests

stage:
	@make -C functions/analysis stage
	@make -C functions/store stage 
	@make -C functions/stats stage