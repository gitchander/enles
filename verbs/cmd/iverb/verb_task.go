package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gitchander/enles/verbs/v1/verbs"
)

var verbForms = []verbs.VerbForm{
	verbs.Infinitive,
	verbs.SimplePast,
	verbs.PastParticiple,
}

type verbTask struct {
	v  verbs.Verb
	vf verbs.VerbForm

	question string
	answer   string
}

var _ Task = &verbTask{}

func randVerbTask(r *rand.Rand, vs []verbs.Verb) *verbTask {

	v := randVerb(r, vs)
	vf := randVerbForm(r, verbForms)

	var (
		question = fmt.Sprintf("Enter %d forms of verb %q (%s): ", len(verbForms), v.Value(vf), vf)
		answer   = fmt.Sprintf("%s %s %s", v.Infinitive, v.SimplePast, v.PastParticiple)
	)

	return &verbTask{
		v:  v,
		vf: vf,

		question: question,
		answer:   answer,
	}
}

func (p *verbTask) Question() string {
	return p.question
}

func (p *verbTask) CheckAnswer(answer string) (success bool, failMessage string) {

	vfs := verbForms

	fields := strings.Fields(answer)

	if len(fields) != len(vfs) {

		//failMessage = fmt.Sprintf("You mast enter %d forms of the verb", len(vfs))
		failMessage = fmt.Sprintf("invalid number of fields: have %d, want %d", len(fields), len(vfs))

		return false, failMessage
	}

	for i, vf := range vfs {
		if fields[i] != p.v.Value(vf) {
			failMessage = fmt.Sprintf("Wrong, correct are: %q", p.answer)
			//failMessage = fmt.Sprintf("Wrong! There are: %s", p.answer)
			return false, failMessage
		}
	}

	return true, ""
}

//------------------------------------------------------------------------------
type verbTasker struct {
	r  *rand.Rand
	vs []verbs.Verb
}

var _ Tasker = &verbTasker{}

func newVerbTasker(r *rand.Rand, vs []verbs.Verb) *verbTasker {
	return &verbTasker{
		r:  r,
		vs: vs,
	}
}

func (p *verbTasker) Next() (Task, bool) {
	task := randVerbTask(p.r, p.vs)
	return task, true
}
