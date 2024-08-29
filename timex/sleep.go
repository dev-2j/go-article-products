package timex

// func SleepWithCancel(ctx context.Context, ttl time.Duration) {

// 	// https://stackoverflow.com/questions/55135239/how-can-i-sleep-with-responsive-context-cancelation

// 	// ctx, cancel := context.WithCancel(context.Background())
// 	// defer cancel()
// 	go func() {
// 		// t := time.Now()
// 		select {
// 		case <-ctx.Done():
// 		case <-time.After(ttl):
// 		}
// 		// fmt.Printf("here after: %v\n", time.Since(t))
// 	}()

// }
