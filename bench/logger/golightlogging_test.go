//
// Copied from : https://github.com/imkira/go-loggers-bench
// 
//
package bench

import (
	"testing"
	"sync/atomic"
	"github.com/dlmc/golight/logger"
)


type blackholeStream struct {
	writeCount uint64
}

func (s *blackholeStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *blackholeStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}


func BenchmarkGolightLoggingTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	log := logger.New(stream, true)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGolightLoggingTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	log := logger.New(stream, true).
		Level(logger.LogError)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
//			log.InfoMsg("The quick brown fox jumps over the lazy dog")
			log.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGolightLoggingJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	log := logger.New(stream, true).
		Level(logger.LogError)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")

			//log.Info().Msg("m", logger.LogObj{"rate":"15", "low":16,"high":123.2,"m":"The quick brown fox jumps over the lazy dog"})
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGolightLoggingJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	log := logger.New(stream, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//log.Info().Interface("m", logger.LogObj{"rate":"15", "low":16,"high":123.2,"m":"The quick brown fox jumps over the lazy dog"})

			log.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count got %v, want %v", stream.WriteCount(), uint64(b.N) )
	}
}