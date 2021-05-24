CMD=docker run --rm -i hadolint/hadolint

lint:
	(cd .. && .docker/hadolint.sh)



	# (cd .. && git ls-files | grep Dockerfile | xargs -0 -I {} sh -c "docker run --rm -i hadolint/hadolint < {}" )