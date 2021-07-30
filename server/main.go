// grpc-chunker-example
// Copyright (c) 2021-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.
package main

import (
	"crypto/rand"
	"grpc-chunker-example/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const chunkSize = 1024 * 1024

type chunkerSrv []byte

func (c chunkerSrv) Chunker(_ *pb.Empty, srv pb.Chunker_ChunkerServer) error {
	chunk := &pb.Chunk{}
	n := len(c)

	for cur := 0; cur < n; cur += chunkSize {
		if cur+chunkSize > n {
			chunk.Chunk = c[cur:n]
		} else {
			chunk.Chunk = c[cur : cur+chunkSize]
		}
		if err := srv.Send(chunk); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	blob := make([]byte, 1024*1024*1024)
	rand.Read(blob)
	pb.RegisterChunkerServer(s, chunkerSrv(blob))
	log.Println("serving on localhost:8888")
	log.Fatal(s.Serve(listen))
}
