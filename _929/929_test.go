package _929

import "testing"

func Test929_1(t *testing.T){
	ss := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}

	expected := 2

	emails := numUniqueEmails(ss)
	if emails != expected{
		t.Errorf("expected  %d, but %d", expected, emails)
	}
}
