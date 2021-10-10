package utils

import (
	"fmt"
	"time"
)

/*
* This is a simple command line "progress reporter".
* It's usage is (intentionally) very simple.
* Example:
* pr := utils.ProgressReporter{}
* pr.SetTotal(1000)
* for i:=0;i<1000;i++ {
*    // do something slow
*    DoSomethingSlow()
*    pr.Inc()
*    pr.Print()
* }
* The above sippet will print a rate and an ETA once a second.
 */
type ProgressReporter struct {
	lastPrinted   time.Time
	eta           time.Time
	etaCalced     time.Time
	checkpoint    time.Time
	checkpointCtr uint64
	start         time.Time
	total         uint64
	done          uint64
}

func (p *ProgressReporter) SetTotal(total uint64) {
	p.total = total
}
func (p *ProgressReporter) Set(a uint64) {
	if p.done == 0 {
		p.start = time.Now()
		p.checkpoint = p.start
	}
	p.done = a
	if time.Since(p.checkpoint) > (time.Duration(300) * time.Second) {
		p.checkpoint = time.Now()
		p.checkpointCtr = p.done
	}

}
func (p *ProgressReporter) Inc() {
	p.Add(1)
}
func (p *ProgressReporter) Add(a uint64) {
	if p.done == 0 {
		p.start = time.Now()
		p.checkpoint = p.start
	}
	p.done = p.done + a
	if time.Since(p.checkpoint) > (time.Duration(300) * time.Second) {
		p.checkpoint = time.Now()
		p.checkpointCtr = p.done
	}

}
func (p *ProgressReporter) Eta() time.Time {
	if time.Since(p.etaCalced) < (time.Duration(5) * time.Second) {
		return p.eta
	}
	elapsed := time.Since(p.checkpoint)
	f := float64(elapsed)
	z := float64(p.done - p.checkpointCtr)
	f = float64(elapsed) / z
	f = f * float64(p.total)
	p.eta = time.Now().Add(time.Duration(f))
	p.etaCalced = time.Now()
	return p.eta
}

// return true if it actually printed stuff
func (p *ProgressReporter) PrintSingleLine() bool {
	s := p.String()
	if s == "" {
		return false
	}
	fmt.Printf("%c%s", byte(13), s)
	return true
}
func (p *ProgressReporter) Print() bool {
	s := p.String()
	if s == "" {
		return false
	}
	fmt.Println(s)
	return true
}

func (p *ProgressReporter) Rate() float64 {
	elapsed := float64(time.Since(p.checkpoint) / time.Second)
	z := float64(p.done - p.checkpointCtr)
	f := z / (elapsed)
	return f
}
func (p *ProgressReporter) String() string {
	if (time.Since(p.lastPrinted)) < (time.Duration(1) * time.Second) {
		return ""
	}
	p.lastPrinted = time.Now()
	eta_s := p.Eta().Format("2006-01-02 15:04:05")
	perc := float32(float32(p.done) / float32(p.total) * float32(100))
	return fmt.Sprintf("Processing %d of %d (%2.1f%%), ETA: %v, %s/sec", p.done, p.total, perc, eta_s, PrettyNumber(uint64(p.Rate())))
	//	return fmt.Sprintf("Processing %d of %d (%2.1f%%), ETA: %v, %.1f/sec", p.done, p.total, perc, eta_s, p.Rate())
}
