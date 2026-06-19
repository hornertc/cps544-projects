// Package eth is used for the ethernet switch program
package eth

// imports
import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"log"
)

type internalPort struct {
	Port
	sendQueue chan Frame
}

type forwardFrame struct {
	Frame
	from int
}

// EthernetSwitch implements the ethernet switch functionality
type EthernetSwitch struct {
	ports                 []internalPort
	sizeCh                chan int
	readError, writeError chan error
	incoming              chan forwardFrame
}

// NewEthernetSwitch creates a new switch with sendQueueSize for the size of the sending buffers and the provided ports.
func NewEthernetSwitch(sendQueueSize int, ports ...Port) *EthernetSwitch {
	n := len(ports)
	ps := make([]internalPort, n)
	for i, p := range ports {
		ps[i] = internalPort{
			Port:      p,
			sendQueue: make(chan Frame, sendQueueSize),
		}
	}

	sw := &EthernetSwitch{
		ports:      ps,
		sizeCh:     make(chan int),
		readError:  make(chan error, n),
		writeError: make(chan error, n),
		incoming:   make(chan forwardFrame),
	}

	return sw
}

func (sw *EthernetSwitch) forward() {
	macTable := make(map[MACAddress]chan Frame)

loop:
	for {
		select {
		case sw.sizeCh <- len(macTable):
		case frame, ok := <-sw.incoming:
			if !ok {
				break loop
			}

			_, ok = macTable[frame.Source]
			if !ok {
				macTable[frame.Source] = sw.ports[frame.from].sendQueue
			}

			if frame.Destination == BroadcastAddress {
				sw.broadcastFrame(frame)
				continue
			}

			destQueue, ok := macTable[frame.Destination]
			if !ok {
				sw.broadcastFrame(frame)
				continue
			}

			send(destQueue, frame.Frame)
		}
	}

	for _, p := range sw.ports {
		close(p.sendQueue)
	}
}

// Run the ethernet switch.
// Blocks until all the io.Readers of the ports are closed (return io.EOF).
// Returns any unrecoverable error from reading (other than io.EOF) or writing to the ports.
// Before returning, this closes the writer for each port.
func (sw *EthernetSwitch) Run() error {
	go sw.forward()

	for p := range sw.ports {
		go sw.readPort(p)
		go sw.writePort(p)
	}

	errs := make([]error, 0, 2*len(sw.ports))
	for range sw.ports {
		err := <-sw.readError
		errs = append(errs, err)
	}

	close(sw.incoming)

	for range sw.ports {
		err := <-sw.writeError
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

// readFrames reads frames from the port
func (sw *EthernetSwitch) readPort(p int) {
	for {
		frame, err := ReadFrame(sw.ports[p])
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			sw.readError <- err
			return
		}
		if frame == nil {
			continue
		}
		sw.incoming <- forwardFrame{*frame, p}
	}
}

// writeFrames writes frames to the port
func (sw *EthernetSwitch) writePort(p int) {
	port := sw.ports[p]
	for frame := range port.sendQueue {
		_, err := WriteFrame(port, frame)
		if err != nil {
			sw.writeError <- err
			return
		}
	}
	sw.writeError <- port.Close()
}

// RunSize returns the number of elements of the MAC table.
// This may only be called while Run() is called.
func (sw *EthernetSwitch) RunSize() int {
	return <-sw.sizeCh
}

// ReadFrame reads a single frame from r.
// If the frame is not valid, return a nil Frame and a nil error.
func ReadFrame(r io.Reader) (*Frame, error) {
	buf := make([]byte, 6+6+2+1500+4)

	_, err := io.ReadFull(r, buf[0:14])
	if err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}

	var frame Frame
	copy(frame.Destination[:], buf[0:6])
	copy(frame.Source[:], buf[6:12])
	size := binary.BigEndian.Uint16(buf[12:14])

	_, err = io.ReadFull(r, buf[14:14+size+4])
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	frame.Data = make([]byte, size)
	copy(frame.Data, buf[14:14+size])

	checksum := binary.BigEndian.Uint32(buf[14+size:])
	if checksum != crc32.ChecksumIEEE(buf[:14+size]) {
		log.Print("Skipping frame due to invalid checksum")
		return nil, nil
	}

	return &frame, nil
}

// WriteFrame writes the ethernet frame to w.
func WriteFrame(w io.Writer, frame Frame) (int, error) {
	frameData := frame.serializeWithoutChecksum()

	calculatedChecksum := crc32.ChecksumIEEE(frameData)

	frameData = append(frameData, byte(calculatedChecksum>>24), byte(calculatedChecksum>>16), byte(calculatedChecksum>>8), byte(calculatedChecksum))

	n, err := w.Write(frameData)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return n, nil
}

// serialize data without checksum
func (frame Frame) serializeWithoutChecksum() []byte {
	destSize := len(frame.Destination)
	srcSize := len(frame.Source)
	dataLenSize := 2

	totalSize := destSize + srcSize + dataLenSize + len(frame.Data)

	buf := make([]byte, 0, totalSize)

	buf = append(buf, frame.Destination[:]...)
	buf = append(buf, frame.Source[:]...)
	buf = append(buf, byte(len(frame.Data)>>8), byte(len(frame.Data)))
	buf = append(buf, frame.Data...)

	return buf
}

// broadcast frame
func (sw *EthernetSwitch) broadcastFrame(frame forwardFrame) {
	for i, sendQueue := range sw.ports {
		if i == frame.from {
			continue
		}
		send(sendQueue.sendQueue, frame.Frame)
	}
}

func send(queue chan<- Frame, frame Frame) {
	select {
	case queue <- frame:
	default:
		log.Println("Broadcasting: Dropped packet due to send queue being full:", frame.String())
	}
}
