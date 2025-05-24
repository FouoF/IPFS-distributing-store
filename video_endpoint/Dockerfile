FROM debian:stable-slim

COPY /bin/video_endpoint /
COPY /video.mp4 /data/

WORKDIR /

EXPOSE 50052

CMD ["./video_endpoint"]