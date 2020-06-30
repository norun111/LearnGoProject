package main

func main() {

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(done <-chan interface{}, in <-chan interface{}) (_,_ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				//2つのチャネルのコピー変数
				var out1, out2 = out1, out2

				for i := 0; i < 2; i++ {
					select {
					case out1<-val:
						out1 = nil //それ以降の書き込みをブロック
					case out2<-val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)


}