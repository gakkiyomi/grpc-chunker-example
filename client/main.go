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
	"context"
	"io"
	"log"

	"grpc-chunker-example/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewChunkerClient(conn)
	stream, err := client.Chunker(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	var blob []byte
	for {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				return
			}
			log.Fatal(err)
		}
		blob = append(blob, c.Chunk...)
	}
}
