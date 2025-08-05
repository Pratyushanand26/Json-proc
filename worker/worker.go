package worker

type RawJson struct {
	id     int
	name   string
	score  int
	status string
}

func Newentry(in <-chan RawJson, out chan<- RawJson) {
	
	data := <-in
	if data.score <= 50 {
		data.status = "fail"
	}else{
		data.status="pass"
	}

}
