// Go program to illustrate the concept of multiple interfaces
package main

import &quot;fmt&quot;

// interface 1
type AuthorDetails interface {
	details()
}

// interface 2
type AuthorArticles interface {
	articles()
}

// structure
type author struct {
	a_name	string
	branch	string
	college	string
	year	int
	salary	int
	particles	int
	tarticles	int
}

// implementing method of the interface 1
func (a author) details() {
	fmt.Printf(&quot;Author Name: %s&quot;, a.a_name)
    	fmt.Printf(&quot;\nBranch: %s and passing year: %d&quot;, a.branch, a.year)
    	fmt.Printf(&quot;\nCollege Name: %s&quot;, a.college)
    	fmt.Printf(&quot;\nSalary: %d&quot;, a.salary)
	fmt.Printf(&quot;\nPublished articles: %d&quot;, a.particles)
}

// Implementing method
// of the interface 2
func (a author) articles() {

    pendingarticles := a.tarticles - a.particles
    fmt.Printf(&quot;\nPending articles: %d&quot;, pendingarticles)
}

// Main value
func main() {

    // Assigning values
    // to the structure
    values := author{
        a_name:    &quot;Mickey&quot;,
        branch:    &quot;Computer science&quot;,
        college:   &quot;XYZ&quot;,
        year:      2012,
        salary:    50000,
        particles: 209,
        tarticles: 309,
    }

    // Accessing the method
    // of the interface 1
    var i1 AuthorDetails = values
    i1.details()

    // Accessing the method
    // of the interface 2
    var i2 AuthorArticles = values
    i2.articles()

}
