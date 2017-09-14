
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

// Test cases for Sublogger
func TestSublogger(t *testing.T) {
	t.Run("sublogger", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("k0","v0")
		log.AddKV("kk","vv")
		log1 := log.Sublogger().AddKV("k1", "v123").AddKV("k2", "v2")
		
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
		log := New(out, false)
		log.AddKV("string", "string").
			AddKV("strings", []string{"str1","str2"})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualString", `{"l":"debug","string":"string","strings":["str1","str2"],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("bool", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("bool", true).
			AddKV("bools", []bool{true,false})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualBool", `{"l":"debug","bool":true,"bools":[true,false],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("int", -12).
			AddKV("ints", []int{10,-11})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt", `{"l":"debug","int":-12,"ints":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int8", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("int8", -12).
			AddKV("ints8", []int8{10,-11})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt8", `{"l":"debug","int8":-12,"ints8":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int16", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("int16", -12).
			AddKV("ints16", []int16{10,-11})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt16", `{"l":"debug","int16":-12,"ints16":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("int32", -12).
			AddKV("ints32", []int32{10,-11})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt32", `{"l":"debug","int32":-12,"ints32":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("int64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("int64", -12).
			AddKV("ints64", []int64{10,-11})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInt64", `{"l":"debug","int64":-12,"ints64":[10,-11],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint8", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("uint8", 12).
			AddKV("uints8", []uint8{123,124})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint8", `{"l":"debug","uint8":12,"uints8":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint16", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("uint16", 12).
			AddKV("uints16", []uint16{123,124})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint16", `{"l":"debug","uint16":12,"uints16":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("uint32", 12).
			AddKV("uints32", []uint32{123,124})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint32", `{"l":"debug","uint32":12,"uints32":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("uint64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("uint64", 12).
			AddKV("uints64", []uint64{123,124})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualUint64", `{"l":"debug","uint64":12,"uints64":[123,124],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("float32", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("float32", 12.32).
			AddKV("float32", []float32{-123,124.1})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualFloat32", `{"l":"debug","float32":12.32,"float32":[-123,124.1],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("float64", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("float64", 12.32).
			AddKV("float64", []float64{-123,124.1})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualFloat64", `{"l":"debug","float64":12.32,"float64":[-123,124.1],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("time", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("time", time.Time{}).
			AddKV("time", []time.Time{time.Time{},time.Time{}})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualTime", `{"l":"debug","time":"0001-01-01T00:00:00Z","time":["0001-01-01T00:00:00Z","0001-01-01T00:00:00Z"],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})
	t.Run("duration", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("duration", 1*time.Second).
			AddKV("duration",  []time.Duration{1*time.Second, 2*time.Second})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualDuration", `{"l":"debug","duration":1000,"duration":[1000,2000],"m":"TestAddKVIndividual"}`+"\n", out, t)
	})/*
	t.Run("interface", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("interface", {"a":"str","b":1,"c":[]int{1,2,3}})
		log.Print("TestAddKVIndividual")	
		tResults("TestAddKVIndividualInterface", `{"l":"debug","interface":1000,"m":"TestAddKVIndividual"}`+"\n", out, t)
	})*/
}

func TestAddKV(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("string", "string").
			AddKV("bool", true).
			AddKV("int", int(-1)).
			AddKV("int8", int8(8)).
			AddKV("int16", int16(16)).
			AddKV("int32", int32(32)).
			AddKV("int64", int64(64)).
			AddKV("uint", uint(1)).
			AddKV("uint8", uint8(8)).
			AddKV("uint16", uint16(16)).
			AddKV("uint32", uint32(32)).
			AddKV("uint64", uint64(64)).
			AddKV("float32", float32(32.32)).
			AddKV("float64", float64(64.64)).
			AddKV("time", time.Time{}).
			AddKV("duration", 1*time.Second)
		log.Print("TestAddKV")	
		tResults("TestAddSingle", `{"l":"debug","string":"string","bool":true,"int":-1,"int8":8,"int16":16,"int32":32,"int64":64,"uint":1,"uint8":8,"uint16":16,"uint32":32,"uint64":64,"float32":32.32,"float64":64.64,"time":"0001-01-01T00:00:00Z","duration":1000,"m":"TestAddKV"}`+"\n", out, t)
	})
	t.Run("Array", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("string", []string{"str1","str2"}).
			AddKV("bool", []bool{true,false}).
			AddKV("int", []int{-1, 2}).
			AddKV("int8", []int8{8, 9}).
			AddKV("int16", []int16{16,17}).
			AddKV("int32", []int32{32,33}).
			AddKV("int64", []int64{64,65}).
			AddKV("uint", []uint{1,2}).
			AddKV("uint8", []uint8{8,9}).
			AddKV("uint16", []uint16{16,17}).
			AddKV("uint32", []uint32{32,33}).
			AddKV("uint64", []uint64{64,65}).
			AddKV("float32", []float32{32.32,33.33}).
			AddKV("float64", []float64{64.64,65.65}).
			AddKV("time", []time.Time{time.Time{},time.Time{}}).
			AddKV("duration", []time.Duration{1*time.Second, 2*time.Second})
		log.Print("TestAddKV")	
		tResults("TestAddArray", `{"l":"debug","string":["str1","str2"],"bool":[true,false],"int":[-1,2],"int8":[8,9],"int16":[16,17],"int32":[32,33],"int64":[64,65],"uint":[1,2],"uint8":[8,9],"uint16":[16,17],"uint32":[32,33],"uint64":[64,65],"float32":[32.32,33.33],"float64":[64.64,65.65],"time":["0001-01-01T00:00:00Z","0001-01-01T00:00:00Z"],"duration":[1000,2000],"m":"TestAddKV"}`+"\n", out, t)
	})
	t.Run("EmptyArray", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("string", []string{}).
			AddKV("bool", []bool{}).
			AddKV("int", []int{}).
			AddKV("int8", []int8{}).
			AddKV("int16", []int16{}).
			AddKV("int32", []int32{}).
			AddKV("int64", []int64{}).
			AddKV("uint", []uint{}).
			AddKV("uint8", []uint8{}).
			AddKV("uint16", []uint16{}).
			AddKV("uint32", []uint32{}).
			AddKV("uint64", []uint64{}).
			AddKV("float32", []float32{}).
			AddKV("float64", []float64{}).
			AddKV("time", []time.Time{}).
			AddKV("duration", []time.Duration{})
		log.Print("TestAddKV")	
		tResults("TestAddEmptyArray", `{"l":"debug","string":[],"bool":[],"int":[],"int8":[],"int16":[],"int32":[],"int64":[],"uint":[],"uint8":[],"uint16":[],"uint32":[],"uint64":[],"float32":[],"float64":[],"time":[],"duration":[],"m":"TestAddKV"}`+"\n", out, t)
	})
}


func TestAddKVInterface(t *testing.T) {
	t.Run("KVPair", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.AddKV("interface", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}}).
			AddKV("bool", true)
		log.Debug("m", LogObj{"a":"aa","b":19})
		tResults("TestAddKVInterface", `{"l":"debug","interface":{"a":"aa","b":19,"d":19,"g":{"c":"c","d":[1,2,3]}},"bool":true,"m":{"a":"aa","b":19}}`+"\n", out, t)
	})
}
func TestDebug(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		out.Reset(); log.Debug("m", "s1");    tResults("TestDebugStr", `{"l":"debug","m":"s1"}`+"\n", out, t)
		out.Reset(); log.Debug("m", true);    tResults("TestDebugBool", `{"l":"debug","m":true}`+"\n", out, t)
		out.Reset(); log.Debug("m", int(3));  tResults("TestDebugInt", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", int8(3)); tResults("TestDebugInt8", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", int16(3));tResults("TestDebugInt16", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", int32(3));tResults("TestDebugInt32", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", int64(3));tResults("TestDebugInt64", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", uint(3)); tResults("TestDebugUint", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", uint8(3));tResults("TestDebugUint8", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", uint16(3));tResults("TestDebugUint16", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", uint32(3));tResults("TestDebugUint32", `{"l":"debug","m":3}`+"\n", out, t)
		out.Reset(); log.Debug("m", uint64(3));tResults("TestDebugUint64", `{"l":"debug","m":3}`+"\n", out, t)
	})
	t.Run("Array", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		out.Reset(); log.Debug("m", []string{"s1","s2"});  tResults("TestDebugStr", `{"l":"debug","m":["s1","s2"]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []bool{true,false});  tResults("TestDebugBool", `{"l":"debug","m":[true,false]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []int{1,2,3});  tResults("TestDebugInt", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []int8{1,2,3}); tResults("TestDebugInt8", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []int16{1,2,3});tResults("TestDebugInt16", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []int32{1,2,3});tResults("TestDebugInt32", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []int64{1,2,3});tResults("TestDebugInt64", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []uint{1,2,3});  tResults("TestDebugUint", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []uint8{1,2,3}); tResults("TestDebugUint8", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []uint16{1,2,3});tResults("TestDebugUint16", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []uint32{1,2,3});tResults("TestDebugUint32", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
		out.Reset(); log.Debug("m", []uint64{1,2,3});tResults("TestDebugUint64", `{"l":"debug","m":[1,2,3]}`+"\n", out, t)
	})
	t.Run("Object", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Debug("m", LogObj{"a":"aa","b":19, "d":true, "g":[]int{15,16}})
		tResults("TestDebugObj", `{"l":"debug","m":{"a":"aa","b":19,"d":true,"g":[15,16]}}`+"\n", out, t)
	})
}

func TestInfo(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Info("m", "s1");    tResults("TestInfoStr", `{"l":"info","m":"s1"}`+"\n", out, t)
	})
	t.Run("Array", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Info("m", []bool{true,false});  tResults("TestInfoBool", `{"l":"info","m":[true,false]}`+"\n", out, t)
	})
	t.Run("Object", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Info("m", LogObj{"a":"aa","b":19, "d":true, "g":[]int{15,16}})
		tResults("TestInfoObj", `{"l":"info","m":{"a":"aa","b":19,"d":true,"g":[15,16]}}`+"\n", out, t)
	})
}

func TestWarn(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Warn("m", "s1");    tResults("TestInfoStr", `{"l":"warn","m":"s1"}`+"\n", out, t)
	})
}

func TestError(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.Error("m", "s1");    tResults("TestInfoStr", `{"l":"error","m":"s1"}`+"\n", out, t)
	})
}


func TestLevel(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out, false)
		log.SetLogLevel(LogFatal)
		log.Error("m", "s1");    tResults("TestLevel", "", out, t)
	})
	t.Run("level", func(t *testing.T) {
		level := []LogLevel{LogDebug, LogInfo, LogWarn, LogError, LogPanic, LogDisabled}
		out := &bytes.Buffer{}
		log := New(out, false)
		for lvl := range level {
			tResultsBool(log, lvl, level, t)
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


func tResultsBool(log *Logger, lvl int, levels []LogLevel, t *testing.T) {
	log.SetLogLevel(LogLevel(lvl)); 
	for l := range levels {
		got := log.Enabled(LogLevel(l))
		if l >= lvl {
			if got != true {
				t.Errorf("LogLevel test failed on %v lvl(%v):\ngot:  %v\nwant: %v", l, lvl, got, true)
			}
		} else {
			if got != false {
				t.Errorf("LogLevel test failed on %v lvl(%v):\ngot:  %v\nwant: %v", l, lvl, got, false)
			}
		}
	}
}
