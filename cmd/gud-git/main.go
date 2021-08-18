package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/alxbse/gud-git/pkg/rules"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	var targetRevision string
	var rule string
	flag.StringVar(&targetRevision, "target-revision", "refs/heads/main", "target revision")
	flag.StringVar(&rule, "rule", "", "rule")
	flag.Parse()

	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	master, err := repo.ResolveRevision(plumbing.Revision(targetRevision))
	if err != nil {
		panic(err)
	}

	masterCommit, err := repo.CommitObject(*master)
	if err != nil {
		panic(err)
	}

	headCommit, err := repo.CommitObject(head.Hash())
	if err != nil {
		panic(err)
	}

	mergeBases, err := masterCommit.MergeBase(headCommit)
	if err != nil {
		panic(err)
	}

	logOptions := git.LogOptions{
		From: head.Hash(),
	}

	refIter, err := repo.Log(&logOptions)
	if err != nil {
		panic(err)
	}

	knownVerbs := []string{
		"add",
		"update",
		"delete",
		"remove",
		"disable",
		"revert",
		"use",
		"fix",
		"refactor",
		"generate",
		"move",
		"release",
		"replace",
		"upgrade",
	}

	maxLength := 64

	prepositions := []string{
		"of",
		"in",
		"on",
	}

	var rules []Rule

	switch rule {
	case "NoIssueNumber":
		rules = []Rule{
			&RuleNoIssueNumber{},
		}
	case "CapitalizedWord":
		rules = []Rule{
			&RuleCapitalizedWord{},
		}
	case "ImperativeMood":
		rules = []Rule{
			&RuleImperativeMood{},
		}
	default:
		rules = []Rule{
			&RuleNoFixup{},
			&RuleNoMerge{},
			&RuleNoLeadingWhitespace{},
			&RuleNoSingleWord{},
			&RuleNoIssueNumber{},
			&RuleNoConventionalSpec{},
			&RuleKnownVerb{
				KnownVerbs: knownVerbs,
			},
			&RuleTitleLength{
				MaxLength: maxLength,
			},
			&RuleBaseFormVerb{},
			&RuleCapitalizedWord{},
			&RuleNoTrailingPunct{},
			&RuleImperativeMood{
				Prepositions: prepositions,
			},
		}
	}

	hasErrors := false
	err = refIter.ForEach(func(c *object.Commit) error {
		if c.Hash == mergeBases[0].Hash {
			return object.ErrCanceled
		}

		messages := []string{}
		for _, rule := range rules {
			valid, message := rule.Validate(c)
			if !valid {
				hasErrors = true
				messages = append(messages, message)
				if rule.BreakWhenInvalid() {
					break
				}
			}
		}
		if len(messages) >= 1 {
			fmt.Printf("\033[33mcommit %s\033[0m\n\n", c.Hash)
			for _, message := range messages {
				fmt.Printf("%s\n\n", message)
			}
		}
		return nil
	})

	if hasErrors {
		os.Exit(1)
	}
}
