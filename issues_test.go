package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// ( •_•)
func TestWorksForGitlabHttpsRemotes(t *testing.T) {
	test_case := `origin   git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte/issues`
	assert.Equal(t, expected, GetUrl(test_case), "works for https remote")
}

// ( •_•)>⌐■-■
func TestWorksForSshRemotes(t *testing.T) {
	test_case := `origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)
origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (push)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte/issues`
	assert.Equal(t, expected, GetUrl(test_case), "works for ssh remote")
}

// (⌐■_■)
func TestWorksForHerokuRemotes(t *testing.T)  {
	test_case := `heroku	https://git.heroku.com/murmuring-island-99377.git (fetch)
heroku	https://git.heroku.com/murmuring-island-99377.git (push)
origin	https://github.com/syzer/poker-player-go (fetch)
origin	https://github.com/syzer/poker-player-go (push)`
	expected  := `https://github.com/syzer/poker-player-go/issues`
	assert.Equal(t, expected, GetUrl(test_case), "works for with heroku remotes too")
}

// ┬─┬ ノ(º_ºノ)
func TestWorksForHttpsRemotes(t *testing.T) {
	test_case := `origin  https://github.com/syzer/issues.git (fetch)`
	// issues inception
	expected := `https://github.com/syzer/issues/issues`
	assert.Equal(t, GetUrl(test_case), expected, "works for https remote")
}