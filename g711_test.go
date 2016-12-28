/*
	Copyright (C) 2016 - 2017, Lefteris Zafiris <zaf@fastmail.com>

	This program is free software, distributed under the terms of
	the BSD 3-Clause License. See the LICENSE file
	at the top of the source tree.

	Package g711 implements encoding and decoding of G711.0 compressed sound data.
	G.711 is an ITU-T standard for audio companding.
*/

package g711

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

// Benchmark Encoding data to Alaw
func BenchmarkAEncode(b *testing.B) {
	rawData, err := ioutil.ReadFile("testing/speech.raw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder, err := NewAlawEncoder(ioutil.Discard, Lpcm)
		if err != nil {
			log.Printf("Failed to create Writer: %s\n", err)
			b.FailNow()
		}
		_, err = encoder.Write(rawData)
		if err != nil {
			log.Printf("Encoding failed: %s\n", err)
			b.FailNow()
		}
	}
}

// Benchmark Encoding data to Ulaw
func BenchmarkUEncode(b *testing.B) {
	rawData, err := ioutil.ReadFile("testing/speech.raw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder, err := NewUlawEncoder(ioutil.Discard, Lpcm)
		if err != nil {
			log.Printf("Failed to create Writer: %s\n", err)
			b.FailNow()
		}
		_, err = encoder.Write(rawData)
		if err != nil {
			log.Printf("Encoding failed: %s\n", err)
			b.FailNow()
		}
	}
}

// Benchmark transcoding g711 data via Reader
func BenchmarkTranscodeR(b *testing.B) {
	alawData, err := ioutil.ReadFile("testing/speech.alaw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transcoder, err := NewAlawDecoder(bytes.NewReader(alawData), Ulaw)
		if err != nil {
			log.Printf("Failed to create Reader: %s\n", err)
			b.FailNow()
		}
		_, err = ioutil.ReadAll(transcoder)
		if err != nil {
			log.Printf("Transcoding failed: %s\n", err)
			b.FailNow()
		}
	}
}

// Benchmark transcoding g711 data via Writer
func BenchmarkTranscodeW(b *testing.B) {
	alawData, err := ioutil.ReadFile("testing/speech.alaw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transcoder, err := NewAlawEncoder(ioutil.Discard, Ulaw)
		if err != nil {
			log.Printf("Failed to create Writer: %s\n", err)
			b.FailNow()
		}
		_, err = transcoder.Write(alawData)
		if err != nil {
			log.Printf("Transcoding failed: %s\n", err)
			b.FailNow()
		}
	}
}

// Benchmark Decoding Ulaw data
func BenchmarkUDecode(b *testing.B) {
	ulawData, err := ioutil.ReadFile("testing/speech.ulaw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder, err := NewUlawDecoder(bytes.NewReader(ulawData), Lpcm)
		if err != nil {
			log.Printf("Failed to create Reader: %s\n", err)
			b.FailNow()
		}
		_, err = ioutil.ReadAll(decoder)
		if err != nil {
			log.Printf("Decoding failed: %s\n", err)
			b.FailNow()
		}
	}
}

// Benchmark Decoding Alaw data
func BenchmarkADecode(b *testing.B) {
	alawData, err := ioutil.ReadFile("testing/speech.alaw")
	if err != nil {
		log.Printf("Failed to read test data: %s\n", err)
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder, err := NewAlawDecoder(bytes.NewReader(alawData), Lpcm)
		if err != nil {
			log.Printf("Failed to create Reader: %s\n", err)
			b.FailNow()
		}
		_, err = ioutil.ReadAll(decoder)
		if err != nil {
			log.Printf("Decoding failed: %s\n", err)
			b.FailNow()
		}
	}
}
