package evaluate

import "time"

func SetT0(t time.Time) {
	T0 = t
	WriteBreakToFile()
	WriteToFile("T0 (Trigger deposit): %s", T0)
}

func SetT0a(t time.Time) {
	T0a = t

	WriteToFile("T0a (Finish deposit): %s", t)

	
}

func SetT1(t time.Time) {
	T1 = t
	WriteToFile("T1 (Relayer caught deposit event): %s", t)

	if !T0a.IsZero() {
		WriteToFileEvaluate("Step 1 (finish deposit to relayer caught): %s", T1.Sub(T0a))
	}
}

func SetT2(t time.Time) {
	T2 = t
	WriteToFile("T2 (Trigger/Start vote): %s", t)
}

func SetT2a(t time.Time) {
	T2 = t
	WriteToFile("T2 (Trigger/Start vote): %s", t)
}

// func SetT2a(t time.Time) {
// 	T2a = t
// 	WriteToFile("T2a (Finish vote): %s", t)
// }

// func SetTimeT2a(t time.Time) {
// 	T2a = t
// 	WriteToFile("This is the last vote -> executing ... (T2a): %s", T2)
// }

func SetT3(t time.Time) {
	T3 = t
	WriteToFileWithCustomPath("./../example/log_new.txt", "T3 (Finish execute - Executed): %s", T3)
	// if !T2a.IsZero() {
	// 	WriteToFileEvaluate("Step 1 (finish deposit to relayer caught): %s", T3.Sub(T2a))
	// }
}

func SetT4(t time.Time) {
	T4 = t
	WriteToFile("Time finish deposit (T4): %s", T4)
}
