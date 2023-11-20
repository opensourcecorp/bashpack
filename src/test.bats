#!/usr/bin/env bats

# bats doesn't like this being fancier for some reason, so this pathing expects
# to be called from repo root as `make test`
source "./src/main.sh"

@test "_sanitize-pkg-name works for various git protocol patterns" {
  for given_path in \
    'http://github.com/org/name' \
    'https://github.com/org/name' \
    'git://github.com/org/name' \
    'git@github.com:org/name' \
    'https://github.com/org/name.git' \
  ; do
    want='github.com/org/name'
    got="$(_sanitize-pkg-name ${given_path})"
    [[ "${got}" == "${want}" ]] || {
      printf 'got == %s, want == %s\n' "${got}" "${want}"
      return 1
    }
  done
}

@test "_cache gets a package" {
  want="${BASHPACK_LIB}/github.com/opensourcecorp/ezlog"
  rm -rf "${want}"
  _cache 'https://github.com/opensourcecorp/ezlog'
  [[ -d "${want}" ]] || {
    printf 'Directory %s should exist but does not\n' "${want}"
    return 1
  }
}

@test "_cache knows about an existing package" {
  want="${BASHPACK_LIB}/github.com/opensourcecorp/ezlog" # should exist from previous test
  EZLOG_LEVEL=debug run _cache 'https://github.com/opensourcecorp/ezlog'
  [[ "${output}" =~ 'already cached' ]] || {
    printf '_cache thinks directory %s is not cached but it should be\n' "${want}"
    printf "\$output was:\n"
    printf '%s' "${output}"
    return 1
  }
}

@test "_mainpath works for a bashpack package" {
  want="${BASHPACK_LIB}/github.com/opensourcecorp/ezlog/src/main.sh" # should exist from a previous test
  got="$(_mainpath https://github.com/opensourcecorp/ezlog)"
  [[ "${got}" == "${want}" ]] || {
    printf 'got == %s, want == %s\n' "${got}" "${want}"
    return 1
  }
}

@test "_mainpath works (as well as it can) for a non-bashpack package" {
  want="${BASHPACK_LIB}/github.com/opensourcecorp/osc-infra/baseimg/scripts/build/main.sh"
  got="$(_mainpath https://github.com/opensourcecorp/osc-infra)"
  [[ "${got}" == "${want}" ]] || {
    printf 'got == %s, want == %s\n' "${got}" "${want}"
    return 1
  }
}
