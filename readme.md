# SockStream

This project is my continuing experimentation with streaming video data over web-sockets in binary format,
and reconstructing the video frame on the client side.

## Goal

The goal of this project is to get a more real-time experience that one can usually get from more
traditional methods like HLS, while also having a data channel that allows you to have synchronized
arbitrary meta data with your video stream.

## TODO

- [ ] Break up the frames in smaller pieces.
- [ ] Reconstruct the smaller pieces on the client side when a complete frame is received.