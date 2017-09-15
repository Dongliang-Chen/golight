
package logger

import (
	"testing"
	"bytes"
	"time"
//	"os"
)


//Simple test case for Print
func TestPrint(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Print("hello world", 23)	
		tResults("TestPrint", `{"l":"debug","m":"hello world23"}`+"\n", out, t)
	})
	t.Run("Empty", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Print("")	
		tResults("TestPrint", `{"l":"debug"}`+"\n", out, t)
	})
}

//Simple test case for Printf
func TestPrintf(t *testing.T) {
	t.Run("Hello", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Printf("hello world %v", 23)	
		tResults("TestPrint", `{"l":"debug","m":"hello world 23"}`+"\n", out, t)
	})
	t.Run("Empty", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Printf("")	
		tResults("TestPrint", `{"l":"debug"}`+"\n", out, t)
	})
}

// Test cases for Sub
func TestSublogger(t *testing.T) {
	t.Run("sublogger", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().Str("k0","v0").Str("kk","vv").Logger()
		log1 := log.With().Str("k1", "v123").Str("k2", "v2").Logger()
		
		log.Print("TestSublogger")
		tResults("TestSublogger", `{"l":"debug","k0":"v0","kk":"vv","m":"TestSublogger"}`+"\n", out, t)

		out.Reset()
		log1.Print("TestSublogger")
		tResults("TestSublogger", `{"l":"debug","k0":"v0","kk":"vv","k1":"v123","k2":"v2","m":"TestSublogger"}`+"\n", out, t)
	})
}

func TestAddKVIndividual(t* testing.T) {
	t.Run("string", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Str("string", "string").
			Strs("strings", []string{"str1","str2"}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualString", `{"l":"debug","string":"string","strings":["str1","str2"],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("bool", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Bool("bool", true).
			Bools("bools", []bool{true,false}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualBool", `{"l":"debug","bool":true,"bools":[true,false],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Int("int", -12).
			Ints("ints", []int{10,-11}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt", `{"l":"debug","int":-12,"ints":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int8", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Int8("int8", -12).
			Ints8("ints8", []int8{10,-11}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt8", `{"l":"debug","int8":-12,"ints8":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int16", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Int16("int16", -12).
			Ints16("ints16", []int16{10,-11}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt16", `{"l":"debug","int16":-12,"ints16":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Int32("int32", -12).
			Ints32("ints32", []int32{10,-11}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt32", `{"l":"debug","int32":-12,"ints32":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Int64("int64", -12).
			Ints64("ints64", []int64{10,-11}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt64", `{"l":"debug","int64":-12,"ints64":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint8", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Uint8("uint8", 12).
			Uints8("uints8", []uint8{123,124}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint8", `{"l":"debug","uint8":12,"uints8":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint16", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Uint16("uint16", 12).
			Uints16("uints16", []uint16{123,124}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint16", `{"l":"debug","uint16":12,"uints16":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Uint32("uint32", 12).
			Uints32("uints32", []uint32{123,124}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint32", `{"l":"debug","uint32":12,"uints32":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Uint64("uint64", 12).
			Uints64("uints64", []uint64{123,124}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint64", `{"l":"debug","uint64":12,"uints64":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("float32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Float32("float32", 12.32).
			Floats32("float32", []float32{-123,124.1}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualFloat32", `{"l":"debug","float32":12.32,"float32":[-123,124.1],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("float64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Float64("float64", 12.32).
			Floats64("float64", []float64{-123,124.1}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualFloat64", `{"l":"debug","float64":12.32,"float64":[-123,124.1],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("time", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Time("time", time.Time{}).
			Times("time", []time.Time{time.Time{},time.Time{}}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualTime", `{"l":"debug","time":"0001-01-01T00:00:00Z","time":["0001-01-01T00:00:00Z","0001-01-01T00:00:00Z"],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("duration", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Dur("duration", 1*time.Second).
			Durs("duration",  []time.Duration{1*time.Second, 2*time.Second}).Logger()
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualDuration", `{"l":"debug","duration":1000,"duration":[1000,2000],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})

}

func TestAddKV(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Str("string", "string").
			Bool("bool", true).
			Int("int", int(-1)).
			Int8("int8", int8(8)).
			Int16("int16", int16(16)).
			Int32("int32", int32(32)).
			Int64("int64", int64(64)).
			Uint("uint", uint(1)).
			Uint8("uint8", uint8(8)).
			Uint16("uint16", uint16(16)).
			Uint32("uint32", uint32(32)).
			Uint64("uint64", uint64(64)).
			Float32("float32", float32(32.32)).
			Float64("float64", float64(64.64)).
			Time("time", time.Time{}).
			Dur("duration", 1*time.Second).Logger()
		log.Print("TestAddKV")	
		tResults("TestAddSingle", `{"l":"debug","string":"string","bool":true,"int":-1,"int8":8,"int16":16,"int32":32,"int64":64,"uint":1,"uint8":8,"uint16":16,"uint32":32,"uint64":64,"float32":32.32,"float64":64.64,"time":"0001-01-01T00:00:00Z","duration":1000,"m":"TestAddKV"}`+"\n", out, t)
	})
	t.Run("Array", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Strs("string", []string{"str1","str2"}).
			Bools("bool", []bool{true,false}).
			Ints("int", []int{-1, 2}).
			Ints8("int8", []int8{8, 9}).
			Ints16("int16", []int16{16,17}).
			Ints32("int32", []int32{32,33}).
			Ints64("int64", []int64{64,65}).
			Uints("uint", []uint{1,2}).
			Uints8("uint8", []uint8{8,9}).
			Uints16("uint16", []uint16{16,17}).
			Uints32("uint32", []uint32{32,33}).
			Uints64("uint64", []uint64{64,65}).
			Floats32("float32", []float32{32.32,33.33}).
			Floats64("float64", []float64{64.64,65.65}).
			Times("time", []time.Time{time.Time{},time.Time{}}).
			Durs("duration", []time.Duration{1*time.Second, 2*time.Second}).Logger()
		log.Print("TestAddKV")	
		tResults("TestAddArray", `{"l":"debug","string":["str1","str2"],"bool":[true,false],"int":[-1,2],"int8":[8,9],"int16":[16,17],"int32":[32,33],"int64":[64,65],"uint":[1,2],"uint8":[8,9],"uint16":[16,17],"uint32":[32,33],"uint64":[64,65],"float32":[32.32,33.33],"float64":[64.64,65.65],"time":["0001-01-01T00:00:00Z","0001-01-01T00:00:00Z"],"duration":[1000,2000],"m":"TestAddKV"}`+"\n", out, t)
	})
	t.Run("EmptyArray", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false).With().
			Strs("string", []string{}).
			Bools("bool", []bool{}).
			Ints("int", []int{}).
			Ints8("int8", []int8{}).
			Ints16("int16", []int16{}).
			Ints32("int32", []int32{}).
			Ints64("int64", []int64{}).
			Uints("uint", []uint{}).
			Uints8("uint8", []uint8{}).
			Uints16("uint16", []uint16{}).
			Uints32("uint32", []uint32{}).
			Uints64("uint64", []uint64{}).
			Floats32("float32", []float32{}).
			Floats64("float64", []float64{}).
			Times("time", []time.Time{}).
			Durs("duration", []time.Duration{}).Logger()
		log.Print("TestAddKV")	
		tResults("TestAddEmptyArray", `{"l":"debug","string":[],"bool":[],"int":[],"int8":[],"int16":[],"int32":[],"int64":[],"uint":[],"uint8":[],"uint16":[],"uint32":[],"uint64":[],"float32":[],"float64":[],"time":[],"duration":[],"m":"TestAddKV"}`+"\n", out, t)
	})
}

func TestAddKVInterface(t *testing.T) {
	t.Run("KVPair", func(t *testing.T) {
		out := &bytes.Buffer{}
		e := New(out, false).With().
			Interface("interface", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}}).
			Bool("bool", true).Logger().Debug()
		e.Msg("done")
		tResults("TestAddKVInterface", `{"l":"debug","interface":{"a":"aa","b":19,"d":19,"g":{"c":"c","d":[1,2,3]}},"bool":true,"m":"done"}`+"\n", out, t)
	})
}

func TestMsg(t *testing.T) {
	t.Run("Debug", func(t *testing.T) {
		out := &bytes.Buffer{}
		e := New(out, false).Debug()
		e.Msg("s1");
		tResults("TestDebugStr", `{"l":"debug","m":"s1"}`+"\n", out, t)
	})
	t.Run("Info", func(t *testing.T) {
		out := &bytes.Buffer{}
		e := New(out, false).Info()
		e.Msg("s1");
		tResults("TestInfoStr", `{"l":"info","m":"s1"}`+"\n", out, t)
	})
	t.Run("Warn", func(t *testing.T) {
		out := &bytes.Buffer{}
		e := New(out, false).Warn()
		e.Msg("s1");
		tResults("TestInfoStr", `{"l":"warn","m":"s1"}`+"\n", out, t)
	})
	t.Run("Error", func(t *testing.T) {
		out := &bytes.Buffer{}
		e := New(out, false).Error()
		e.Msg("s1");
		tResults("TestInfoStr", `{"l":"error","m":"s1"}`+"\n", out, t)
	})
}


func TestLevel(t *testing.T) {
	t.Run("SetGlobalLevel", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		
		SetGlobalLevel(LogDebug)
		log.Error().Msg("s1");    tResults("SetGlobalLevel LogDebug", `{"l":"error","m":"s1"}`+"\n", out, t)

		out.Reset()
		SetGlobalLevel(LogFatal)
		log.Error().Msg("s1");    tResults("SetGlobalLevel LogFatal", "", out, t)

		out.Reset()
		SetGlobalLevel(LogInfo)
		log.Warn().Msg("s1");    tResults("SetGlobalLevel LogError", `{"l":"warn","m":"s1"}`+"\n", out, t)

		out.Reset()
		SetGlobalLevel(LogDisabled)
		log.Panic().Msg("s1");    tResults("SetGlobalLevel LogDisabled", "", out, t)

	})
	t.Run("level", func(t *testing.T) {
		out := &bytes.Buffer{}
		SetGlobalLevel(LogDebug)
		log := New(out, false).Level(LogDebug)
		log.Debug().Msg("s1");
		tResults("level LogDebug", `{"l":"debug","m":"s1"}`+"\n", out, t)

		log = log.Level(LogInfo)
		out.Reset()
		log.Debug().Msg("s1");
		tResults("level LogInfo with Debug", "", out, t)

		out.Reset()
		log.Warn().Msg("s1");
		tResults("level LogInfo with Warn", `{"l":"warn","m":"s1"}`+"\n", out, t)
		
		log = log.Level(LogDebug)
		out.Reset()
		log.Debug().Msg("s1");
		tResults("level LogDebug with Debug", `{"l":"debug","m":"s1"}`+"\n", out, t)

		log = log.Level(LogWarn)
		out.Reset()
		log.Info().Msg("s1");
		tResults("level LogInfo", "", out, t)

		out.Reset()
		log.Warn().Msg("s1");
		tResults("level LogWarn with Warn", `{"l":"warn","m":"s1"}`+"\n", out, t)

		out.Reset()
		log.Error().Msg("s1");
		tResults("level LogWarn with Error", `{"l":"error","m":"s1"}`+"\n", out, t)

		log = log.Level(LogDisabled)
		out.Reset()
		log.Panic().Msg("s1");
		tResults("level LogDisabled with Panic", "", out, t)

		if log.Panic().Enabled() != false {
			t.Errorf("level LogDisabled Panic Enabled failed")
		}
		log = log.Level(LogError)
		if log.Panic().Enabled() != true {
			t.Errorf("level LogError Panic Enabled failed")
		}
		if log.Warn().Enabled() != false {
			t.Errorf("level LogError Warn Enabled failed")
		}
	})	
}

/*
func TestFatal(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Fatal("m", "s1");    tResults("TestInfoStr", `{"l":"fatal","m":"s1"}`+"\n", out, t)
	})
}

func TestPanic(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Panic("m", "s1");    tResults("TestInfoStr", `{"l":"panic","m":"s1"}`+"\n", out, t)
	})
}
*/

func tResults(test, want string, out *bytes.Buffer, t *testing.T) {
	got := out.String()
	if got != want {
		t.Errorf("Log output failed on %v:\ngot:  %v\nwant: %v", test, got, want)
	}
}

